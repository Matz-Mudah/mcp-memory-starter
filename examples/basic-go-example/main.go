package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	_ "modernc.org/sqlite"
)

// Config holds application configuration
type Config struct {
	DatabasePath     string
	EmbeddingBaseURL string
	EmbeddingModel   string
	EmbeddingAPIKey  string
}

// Memory represents a stored memory with its embedding
type Memory struct {
	ID        int64     `json:"id,omitzero"`
	Text      string    `json:"text"`
	Embedding []float64 `json:"embedding,omitzero"`
	CreatedAt time.Time `json:"created_at,omitzero"`
}

// SearchResult pairs a memory with its similarity score
type SearchResult struct {
	Memory     Memory  `json:"memory"`
	Similarity float64 `json:"similarity,omitzero"`
}

// Global database connection
var db *sql.DB
var config Config

func main() {
	// Load configuration
	config = loadConfig()

	// Initialize database
	if err := initDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Create MCP server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "basic-go-memory",
		Version: "1.0.0",
	}, nil)

	// Register tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "store_memory",
		Description: "Store a memory with semantic embedding for later retrieval",
	}, handleStoreMemory)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "search_memory",
		Description: "Search for relevant memories using semantic similarity",
	}, handleSearchMemory)

	// Run server with stdio transport
	ctx := context.Background()
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// loadConfig loads configuration from environment variables
func loadConfig() Config {
	return Config{
		DatabasePath:     getEnv("DATABASE_PATH", "memories.db"),
		EmbeddingBaseURL: getEnv("EMBEDDING_BASE_URL", "http://localhost:1234/v1"),
		EmbeddingModel:   getEnv("EMBEDDING_MODEL", "text-embedding-embeddinggemma-300m-qat"),
		EmbeddingAPIKey:  getEnv("EMBEDDING_API_KEY", "not-needed"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// initDatabase creates the SQLite database and schema
func initDatabase() error {
	var err error
	db, err = sql.Open("sqlite", config.DatabasePath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Create memories table
	schema := `
	CREATE TABLE IF NOT EXISTS memories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL,
		embedding TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_created_at ON memories(created_at DESC);
	`

	if _, err := db.Exec(schema); err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	return nil
}

// StoreMemoryInput defines the input for store_memory tool
type StoreMemoryInput struct {
	Text string `json:"text" jsonschema:"The text to remember"`
}

// StoreMemoryOutput defines the output for store_memory tool
type StoreMemoryOutput struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitzero"`
	ID      int64  `json:"id,omitzero"`
}

// handleStoreMemory implements the store_memory tool
func handleStoreMemory(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input StoreMemoryInput,
) (*mcp.CallToolResult, StoreMemoryOutput, error) {
	if input.Text == "" {
		return nil, StoreMemoryOutput{}, fmt.Errorf("text cannot be empty")
	}

	// Generate embedding
	embedding, err := generateEmbedding(input.Text)
	if err != nil {
		return nil, StoreMemoryOutput{}, fmt.Errorf("failed to generate embedding: %w", err)
	}

	// Store in database
	embeddingJSON, err := json.Marshal(embedding)
	if err != nil {
		return nil, StoreMemoryOutput{}, fmt.Errorf("failed to marshal embedding: %w", err)
	}

	result, err := db.Exec(
		"INSERT INTO memories (text, embedding) VALUES (?, ?)",
		input.Text,
		string(embeddingJSON),
	)
	if err != nil {
		return nil, StoreMemoryOutput{}, fmt.Errorf("failed to store memory: %w", err)
	}

	id, _ := result.LastInsertId()

	return nil, StoreMemoryOutput{
		Success: true,
		Message: fmt.Sprintf("Memory stored successfully with ID %d", id),
		ID:      id,
	}, nil
}

// SearchMemoryInput defines the input for search_memory tool
type SearchMemoryInput struct {
	Query         string  `json:"query" jsonschema:"The search query"`
	Limit         int     `json:"limit,omitempty" jsonschema:"Maximum number of results (default: 5)"`
	MinSimilarity float64 `json:"min_similarity,omitempty" jsonschema:"Minimum similarity score 0-1 (default: 0.0)"`
}

// SearchMemoryOutput defines the output for search_memory tool
type SearchMemoryOutput struct {
	Results []SearchResult `json:"results,omitzero"`
	Count   int            `json:"count,omitzero"`
}

// handleSearchMemory implements the search_memory tool
func handleSearchMemory(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input SearchMemoryInput,
) (*mcp.CallToolResult, SearchMemoryOutput, error) {
	if input.Query == "" {
		return nil, SearchMemoryOutput{}, fmt.Errorf("query cannot be empty")
	}

	// Set defaults
	if input.Limit == 0 {
		input.Limit = 5
	}
	// MinSimilarity defaults to 0.0 (no filtering), users can optionally set a threshold

	// Generate query embedding
	queryEmbedding, err := generateEmbedding(input.Query)
	if err != nil {
		return nil, SearchMemoryOutput{}, fmt.Errorf("failed to generate embedding: %w", err)
	}

	// Search database
	rows, err := db.Query("SELECT id, text, embedding, created_at FROM memories")
	if err != nil {
		return nil, SearchMemoryOutput{}, fmt.Errorf("failed to query memories: %w", err)
	}
	defer rows.Close()

	var results []SearchResult
	for rows.Next() {
		var memory Memory
		var embeddingJSON string

		if err := rows.Scan(&memory.ID, &memory.Text, &embeddingJSON, &memory.CreatedAt); err != nil {
			return nil, SearchMemoryOutput{}, fmt.Errorf("failed to scan memory row: %w", err)
		}

		if err := json.Unmarshal([]byte(embeddingJSON), &memory.Embedding); err != nil {
			return nil, SearchMemoryOutput{}, fmt.Errorf("failed to unmarshal embedding for memory %d: %w", memory.ID, err)
		}

		// Calculate cosine similarity
		similarity := cosineSimilarity(queryEmbedding, memory.Embedding)

		if similarity >= input.MinSimilarity {
			// Clear embedding from output to save tokens
			memory.Embedding = nil

			results = append(results, SearchResult{
				Memory:     memory,
				Similarity: similarity,
			})
		}
	}

	// Sort by similarity (descending)
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].Similarity > results[i].Similarity {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	// Limit results
	if len(results) > input.Limit {
		results = results[:input.Limit]
	}

	return nil, SearchMemoryOutput{
		Results: results,
		Count:   len(results),
	}, nil
}

// generateEmbedding generates an embedding vector using LM Studio
func generateEmbedding(text string) ([]float64, error) {
	reqBody := map[string]interface{}{
		"model": config.EmbeddingModel,
		"input": text,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", config.EmbeddingBaseURL+"/embeddings",
		bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if config.EmbeddingAPIKey != "" {
		req.Header.Set("Authorization", "Bearer "+config.EmbeddingAPIKey)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("embedding API returned status %d", resp.StatusCode)
	}

	var result struct {
		Data []struct {
			Embedding []float64 `json:"embedding"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("no embedding returned")
	}

	return result.Data[0].Embedding, nil
}

// cosineSimilarity calculates the cosine similarity between two vectors
func cosineSimilarity(a, b []float64) float64 {
	if len(a) != len(b) {
		return 0.0
	}

	var dotProduct, normA, normB float64
	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	if normA == 0 || normB == 0 {
		return 0.0
	}

	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}
