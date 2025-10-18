package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/pgvector/pgvector-go"
)

// PostgresConfig holds Postgres connection configuration
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// PostgresStore implements storage using PostgreSQL with pgvector and Apache AGE
type PostgresStore struct {
	db *sql.DB
}

// Memory represents a stored memory with metadata
type Memory struct {
	ID        int64     `json:"id,omitzero"`
	Text      string    `json:"text"`
	Embedding []float64 `json:"embedding,omitzero"`
	GroupID   string    `json:"group_id,omitzero"`
	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
}

// SearchResult pairs a memory with its similarity score
type SearchResult struct {
	Memory           Memory  `json:"memory"`
	Similarity       float64 `json:"similarity,omitzero"`
	ViaRelationship  bool    `json:"via_relationship,omitzero"`  // True if found via graph traversal
	RelationshipHops int     `json:"relationship_hops,omitzero"` // Number of hops from vector result
}

// NewPostgresStore creates a new PostgreSQL store with pgvector and Apache AGE
func NewPostgresStore(config PostgresConfig) (*PostgresStore, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Note: Apache AGE initialization (LOAD 'age' and SET search_path) is done
	// per-query in methods that need it, since connection pool makes it unreliable
	// to set globally

	return &PostgresStore{db: db}, nil
}

// Close closes the database connection
func (s *PostgresStore) Close() error {
	return s.db.Close()
}

// StoreMemory stores a memory with its vector embedding
func (s *PostgresStore) StoreMemory(text string, embedding []float64, groupID string) (int64, error) {
	// Convert []float64 to []float32 for pgvector
	embedding32 := make([]float32, len(embedding))
	for i, v := range embedding {
		embedding32[i] = float32(v)
	}

	var id int64
	query := `
		INSERT INTO memories (text, embedding, group_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := s.db.QueryRow(query, text, pgvector.NewVector(embedding32), groupID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to store memory: %w", err)
	}

	// Create corresponding node in Apache AGE graph
	// Initialize AGE for this connection
	if _, err := s.db.Exec("LOAD 'age'; SET search_path = ag_catalog, '$user', public;"); err != nil {
		return 0, fmt.Errorf("failed to initialize Apache AGE for memory %d: %w", id, err)
	}

	cypherQuery := fmt.Sprintf(`
		SELECT * FROM cypher('memory_graph', $$
			CREATE (m:Memory {id: %d, text: '%s'})
			RETURN m
		$$) as (memory agtype);
	`, id, escapeString(text))

	_, err = s.db.Exec(cypherQuery)
	if err != nil {
		return 0, fmt.Errorf("failed to create AGE node for memory %d: %w", id, err)
	}

	return id, nil
}

// escapeString escapes single quotes for Cypher queries
func escapeString(s string) string {
	// Escape single quotes by replacing ' with \'
	return strings.Replace(s, "'", "\\'", -1)
}

// SearchMemories performs vector similarity search using pgvector
func (s *PostgresStore) SearchMemories(queryEmbedding []float64, limit int, minSimilarity float64, groupID string) ([]SearchResult, error) {
	// Convert []float64 to []float32 for pgvector
	embedding32 := make([]float32, len(queryEmbedding))
	for i, v := range queryEmbedding {
		embedding32[i] = float32(v)
	}

	// Build query with optional group filter
	query := `
		SELECT
			id, text, embedding::text, group_id, created_at, updated_at,
			1 - (embedding <=> $1) as similarity
		FROM memories
	`
	args := []interface{}{pgvector.NewVector(embedding32)}

	if groupID != "" {
		query += " WHERE group_id = $2"
		args = append(args, groupID)
		if minSimilarity > 0 {
			query += " AND 1 - (embedding <=> $1) >= $3"
			args = append(args, minSimilarity)
			query += " ORDER BY embedding <=> $1 LIMIT $4"
			args = append(args, limit)
		} else {
			query += " ORDER BY embedding <=> $1 LIMIT $3"
			args = append(args, limit)
		}
	} else {
		if minSimilarity > 0 {
			query += " WHERE 1 - (embedding <=> $1) >= $2"
			args = append(args, minSimilarity)
			query += " ORDER BY embedding <=> $1 LIMIT $3"
			args = append(args, limit)
		} else {
			query += " ORDER BY embedding <=> $1 LIMIT $2"
			args = append(args, limit)
		}
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search memories: %w", err)
	}
	defer rows.Close()

	var results []SearchResult
	for rows.Next() {
		var memory Memory
		var embeddingStr string
		var similarity float64
		var groupIDPtr *string

		err := rows.Scan(
			&memory.ID,
			&memory.Text,
			&embeddingStr,
			&groupIDPtr,
			&memory.CreatedAt,
			&memory.UpdatedAt,
			&similarity,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan search result: %w", err)
		}

		// Set optional fields
		if groupIDPtr != nil {
			memory.GroupID = *groupIDPtr
		}

		// Don't parse embedding for output - saves tokens and not needed by user
		// The embedding is only used by pgvector for similarity calculation

		results = append(results, SearchResult{
			Memory:     memory,
			Similarity: similarity,
		})
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating search results: %w", err)
	}

	// Graph-enhanced search: traverse relationships from top results
	if len(results) > 0 {
		// Extract IDs from vector search results
		resultIDs := make([]int64, len(results))
		resultMap := make(map[int64]bool) // Track which IDs we already have
		for i, result := range results {
			resultIDs[i] = result.Memory.ID
			resultMap[result.Memory.ID] = true
		}

		// Find connected memories (1 hop)
		connectedIDs, err := s.getConnectedMemoryIDs(resultIDs, 1)
		if err == nil && len(connectedIDs) > 0 {
			// Fetch connected memories that aren't already in results
			var newIDs []int64
			for id := range connectedIDs {
				if !resultMap[id] {
					newIDs = append(newIDs, id)
				}
			}

			if len(newIDs) > 0 {
				// Build query to fetch connected memories
				placeholders := ""
				args := make([]interface{}, len(newIDs))
				for i, id := range newIDs {
					if i > 0 {
						placeholders += ","
					}
					placeholders += fmt.Sprintf("$%d", i+1)
					args[i] = id
				}

				fetchQuery := fmt.Sprintf(`
					SELECT id, text, group_id, created_at, updated_at
					FROM memories
					WHERE id IN (%s)
				`, placeholders)

				connectedRows, err := s.db.Query(fetchQuery, args...)
				if err == nil {
					defer connectedRows.Close()

					for connectedRows.Next() {
						var memory Memory
						var groupIDPtr *string

						if err := connectedRows.Scan(
							&memory.ID,
							&memory.Text,
							&groupIDPtr,
							&memory.CreatedAt,
							&memory.UpdatedAt,
						); err == nil {
							if groupIDPtr != nil {
								memory.GroupID = *groupIDPtr
							}

							// Add as graph-discovered result with lower similarity
							results = append(results, SearchResult{
								Memory:           memory,
								Similarity:       0, // No vector similarity, found via graph
								ViaRelationship:  true,
								RelationshipHops: connectedIDs[memory.ID],
							})
						}
					}
				}
			}
		}
	}

	return results, nil
}

// AddRelationship creates a graph edge between two memories using Apache AGE
func (s *PostgresStore) AddRelationship(fromID, toID int64, relType string, properties map[string]interface{}) error {
	// Initialize AGE for this connection
	if _, err := s.db.Exec("LOAD 'age'; SET search_path = ag_catalog, '$user', public;"); err != nil {
		return fmt.Errorf("failed to initialize AGE: %w", err)
	}

	// Build Cypher properties string
	propsStr := ""
	if len(properties) > 0 {
		propsStr = "{"
		first := true
		for key, value := range properties {
			if !first {
				propsStr += ", "
			}
			first = false
			// Format value as string for Cypher
			switch v := value.(type) {
			case string:
				// Escape single quotes for Cypher by replacing ' with \'
				escaped := strings.Replace(v, "'", "\\'", -1)
				propsStr += fmt.Sprintf("%s: '%s'", key, escaped)
			case float64, int, int64:
				propsStr += fmt.Sprintf("%s: %v", key, v)
			case bool:
				propsStr += fmt.Sprintf("%s: %v", key, v)
			default:
				escaped := strings.Replace(fmt.Sprintf("%v", v), "'", "\\'", -1)
				propsStr += fmt.Sprintf("%s: '%s'", key, escaped)
			}
		}
		propsStr += "}"
	}

	// Create relationship using Apache AGE Cypher
	query := fmt.Sprintf(`
		SELECT * FROM cypher('memory_graph', $$
			MATCH (a:Memory {id: %d})
			MATCH (b:Memory {id: %d})
			MERGE (a)-[r:%s %s]->(b)
			RETURN r
		$$) as (relationship agtype);
	`, fromID, toID, relType, propsStr)

	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create relationship: %w", err)
	}

	return nil
}

// GetMemoryByID retrieves a single memory by its ID
func (s *PostgresStore) GetMemoryByID(id int64) (*Memory, error) {
	query := `SELECT id, text, group_id, created_at, updated_at FROM memories WHERE id = $1`

	var memory Memory
	var groupIDPtr *string

	err := s.db.QueryRow(query, id).Scan(
		&memory.ID,
		&memory.Text,
		&groupIDPtr,
		&memory.CreatedAt,
		&memory.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("memory not found: %d", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get memory: %w", err)
	}

	if groupIDPtr != nil {
		memory.GroupID = *groupIDPtr
	}

	return &memory, nil
}

// ExploreConnections finds related memories using Apache AGE graph traversal
func (s *PostgresStore) ExploreConnections(memoryID int64, maxDepth int) ([]Memory, error) {
	// Initialize AGE for this connection
	if _, err := s.db.Exec("LOAD 'age'; SET search_path = ag_catalog, '$user', public;"); err != nil {
		return nil, fmt.Errorf("failed to initialize AGE: %w", err)
	}

	// Use Apache AGE to find connected memories
	query := fmt.Sprintf(`
		SELECT * FROM cypher('memory_graph', $$
			MATCH path = (start:Memory {id: %d})-[*1..%d]-(connected:Memory)
			RETURN DISTINCT connected.id
		$$) as (connected_id agtype);
	`, memoryID, maxDepth)

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to explore connections: %w", err)
	}
	defer rows.Close()

	var connectedIDs []int64
	for rows.Next() {
		var idJSON string
		if err := rows.Scan(&idJSON); err != nil {
			continue
		}

		// Parse agtype JSON
		var id int64
		if err := json.Unmarshal([]byte(idJSON), &id); err != nil {
			continue
		}

		connectedIDs = append(connectedIDs, id)
	}

	// Fetch full memory objects
	if len(connectedIDs) == 0 {
		return []Memory{}, nil
	}

	// Build query to fetch memories
	placeholders := ""
	args := make([]interface{}, len(connectedIDs))
	for i, id := range connectedIDs {
		if i > 0 {
			placeholders += ","
		}
		placeholders += fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	fetchQuery := fmt.Sprintf(`
		SELECT id, text, group_id, created_at, updated_at
		FROM memories
		WHERE id IN (%s)
	`, placeholders)

	memoryRows, err := s.db.Query(fetchQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch connected memories: %w", err)
	}
	defer memoryRows.Close()

	var memories []Memory
	for memoryRows.Next() {
		var memory Memory
		var groupIDPtr *string

		err := memoryRows.Scan(
			&memory.ID,
			&memory.Text,
			&groupIDPtr,
			&memory.CreatedAt,
			&memory.UpdatedAt,
		)
		if err != nil {
			continue
		}

		if groupIDPtr != nil {
			memory.GroupID = *groupIDPtr
		}

		memories = append(memories, memory)
	}

	return memories, nil
}

// getConnectedMemoryIDs finds memories connected to the given IDs via graph relationships
// Returns a map of connected memory ID -> number of hops from original results
func (s *PostgresStore) getConnectedMemoryIDs(startIDs []int64, maxDepth int) (map[int64]int, error) {
	if len(startIDs) == 0 {
		return make(map[int64]int), nil
	}

	// Initialize AGE
	if _, err := s.db.Exec("LOAD 'age'; SET search_path = ag_catalog, '$user', public;"); err != nil {
		return nil, fmt.Errorf("failed to initialize AGE: %w", err)
	}

	connectedIDs := make(map[int64]int) // memory ID -> hop distance

	// For each starting ID, traverse the graph
	for _, startID := range startIDs {
		query := fmt.Sprintf(`
			SELECT * FROM cypher('memory_graph', $$
				MATCH path = (start:Memory {id: %d})-[*1..%d]-(connected:Memory)
				RETURN DISTINCT connected.id, length(path) as hops
			$$) as (connected_id agtype, hop_count agtype);
		`, startID, maxDepth)

		rows, err := s.db.Query(query)
		if err != nil {
			// Graph might not have this node, continue
			continue
		}

		for rows.Next() {
			var idJSON, hopsJSON string
			if err := rows.Scan(&idJSON, &hopsJSON); err != nil {
				continue
			}

			var id, hops int64
			if err := json.Unmarshal([]byte(idJSON), &id); err != nil {
				continue
			}
			if err := json.Unmarshal([]byte(hopsJSON), &hops); err != nil {
				continue
			}

			// Skip if we already found this ID from a shorter path
			if existingHops, exists := connectedIDs[id]; !exists || int(hops) < existingHops {
				connectedIDs[id] = int(hops)
			}
		}
		rows.Close()
	}

	return connectedIDs, nil
}
