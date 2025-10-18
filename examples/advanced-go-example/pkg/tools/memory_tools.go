package tools

import (
	"context"
	"fmt"

	"advanced-go-example/pkg/embeddings"
	"advanced-go-example/pkg/llm"
	"advanced-go-example/pkg/storage"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterMemoryTools registers all memory-related MCP tools
func RegisterMemoryTools(server *mcp.Server, store *storage.PostgresStore, embClient *embeddings.Client, llmClient *llm.Client) {
	// Wrap dependencies in a handler struct
	h := &memoryHandler{
		store:      store,
		embeddings: embClient,
		llm:        llmClient,
	}

	// Register tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "store_memory",
		Description: "Store a memory with vector embedding and optional metadata",
	}, h.handleStoreMemory)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "search_memories",
		Description: "Search for relevant memories using semantic similarity (pgvector)",
	}, h.handleSearchMemories)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "add_relationship",
		Description: "Create a graph relationship between two memories (Apache AGE)",
	}, h.handleAddRelationship)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "explore_connections",
		Description: "Find memories connected through graph relationships (Apache AGE)",
	}, h.handleExploreConnections)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "auto_detect_relationships",
		Description: "Automatically detect and create relationships using LLM analysis of semantic similarity",
	}, h.handleAutoDetectRelationships)
}

// memoryHandler holds dependencies for tool handlers
type memoryHandler struct {
	store      *storage.PostgresStore
	embeddings *embeddings.Client
	llm        *llm.Client
}

// StoreMemoryInput defines input for store_memory tool
type StoreMemoryInput struct {
	Text                   string `json:"text" jsonschema:"The text to remember"`
	GroupID                string `json:"group_id,omitempty" jsonschema:"Optional group identifier"`
	AutoDetectRelationships *bool  `json:"auto_detect_relationships,omitempty" jsonschema:"Automatically detect relationships using LLM (default: true)"`
}

// StoreMemoryOutput defines output for store_memory tool
type StoreMemoryOutput struct {
	Success              bool   `json:"success"`
	Message              string `json:"message,omitzero"`
	ID                   int64  `json:"id,omitzero"`
	RelationshipsCreated int    `json:"relationships_created,omitzero"`
}

func (h *memoryHandler) handleStoreMemory(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input StoreMemoryInput,
) (*mcp.CallToolResult, StoreMemoryOutput, error) {
	if input.Text == "" {
		return nil, StoreMemoryOutput{}, fmt.Errorf("text cannot be empty")
	}

	// Default auto_detect_relationships to true if not specified
	autoDetect := true
	if input.AutoDetectRelationships != nil {
		autoDetect = *input.AutoDetectRelationships
	}

	// Generate embedding
	embedding, err := h.embeddings.Generate(input.Text)
	if err != nil {
		return nil, StoreMemoryOutput{}, fmt.Errorf("failed to generate embedding: %w", err)
	}

	// Store in database
	id, err := h.store.StoreMemory(input.Text, embedding, input.GroupID)
	if err != nil {
		return nil, StoreMemoryOutput{}, fmt.Errorf("failed to store memory: %w", err)
	}

	relationshipsCreated := 0

	// Auto-detect relationships if enabled
	if autoDetect {
		// Find similar memories as candidates
		searchResults, err := h.store.SearchMemories(embedding, 10, 0.5, "")
		if err == nil {
			// Filter out the newly stored memory itself and convert to candidates
			var candidates []llm.CandidateMemory
			for _, result := range searchResults {
				if result.Memory.ID != id {
					candidates = append(candidates, llm.CandidateMemory{
						ID:         result.Memory.ID,
						Text:       result.Memory.Text,
						Similarity: result.Similarity,
					})
				}
			}

			// Use LLM to analyze relationships if we have candidates
			if len(candidates) > 0 {
				llmSuggestions, err := h.llm.AnalyzeRelationships(input.Text, id, candidates)
				if err != nil {
					// Don't fail the store operation if LLM analysis fails
					// The memory is already stored successfully
					return nil, StoreMemoryOutput{
						Success: true,
						Message: fmt.Sprintf("Memory stored with ID %d (relationship auto-detection failed: %v)", id, err),
						ID:      id,
					}, nil
				}

				// Create high-confidence relationships
				for _, suggestion := range llmSuggestions {
					if suggestion.Confidence >= 0.7 {
						props := map[string]interface{}{
							"reason":        suggestion.Reason,
							"confidence":    suggestion.Confidence,
							"auto_detected": true,
						}

						err := h.store.AddRelationship(id, suggestion.TargetID, suggestion.Type, props)
						if err != nil {
							// Return error if relationship creation fails
							return nil, StoreMemoryOutput{}, fmt.Errorf("failed to create auto-detected relationship %s from %d to %d: %w", suggestion.Type, id, suggestion.TargetID, err)
						}
						relationshipsCreated++
					}
				}
			}
		}
	}

	message := fmt.Sprintf("Memory stored successfully with ID %d", id)
	if relationshipsCreated > 0 {
		message = fmt.Sprintf("Memory stored with ID %d and %d relationships auto-created", id, relationshipsCreated)
	}

	return nil, StoreMemoryOutput{
		Success:              true,
		Message:              message,
		ID:                   id,
		RelationshipsCreated: relationshipsCreated,
	}, nil
}

// SearchMemoriesInput defines input for search_memories tool
type SearchMemoriesInput struct {
	Query         string  `json:"query" jsonschema:"The search query"`
	Limit         int     `json:"limit,omitempty" jsonschema:"Maximum results (default: 5)"`
	MinSimilarity float64 `json:"min_similarity,omitempty" jsonschema:"Minimum similarity 0-1 (default: 0.0)"`
	GroupID       string  `json:"group_id,omitempty" jsonschema:"Optional group filter"`
}

// SearchMemoriesOutput defines output for search_memories tool
type SearchMemoriesOutput struct {
	Results []storage.SearchResult `json:"results"`
	Count   int                    `json:"count"`
}

func (h *memoryHandler) handleSearchMemories(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input SearchMemoriesInput,
) (*mcp.CallToolResult, SearchMemoriesOutput, error) {
	if input.Query == "" {
		return nil, SearchMemoriesOutput{}, fmt.Errorf("query cannot be empty")
	}

	// Set defaults
	if input.Limit == 0 {
		input.Limit = 5
	}

	// Generate query embedding
	queryEmbedding, err := h.embeddings.Generate(input.Query)
	if err != nil {
		return nil, SearchMemoriesOutput{}, fmt.Errorf("failed to generate embedding: %w", err)
	}

	// Search using pgvector
	results, err := h.store.SearchMemories(queryEmbedding, input.Limit, input.MinSimilarity, input.GroupID)
	if err != nil {
		return nil, SearchMemoriesOutput{}, fmt.Errorf("failed to search memories: %w", err)
	}

	// Ensure results is never nil (empty array instead)
	if results == nil {
		results = []storage.SearchResult{}
	}

	return nil, SearchMemoriesOutput{
		Results: results,
		Count:   len(results),
	}, nil
}

// AddRelationshipInput defines input for add_relationship tool
type AddRelationshipInput struct {
	FromID     int64                  `json:"from_id" jsonschema:"Source memory ID"`
	ToID       int64                  `json:"to_id" jsonschema:"Target memory ID"`
	Type       string                 `json:"type" jsonschema:"Relationship type (e.g., RELATES_TO, SIMILAR_TO)"`
	Properties map[string]interface{} `json:"properties,omitempty" jsonschema:"Optional relationship properties"`
}

// AddRelationshipOutput defines output for add_relationship tool
type AddRelationshipOutput struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitzero"`
}

func (h *memoryHandler) handleAddRelationship(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input AddRelationshipInput,
) (*mcp.CallToolResult, AddRelationshipOutput, error) {
	if input.FromID == 0 || input.ToID == 0 {
		return nil, AddRelationshipOutput{}, fmt.Errorf("from_id and to_id are required")
	}
	if input.Type == "" {
		return nil, AddRelationshipOutput{}, fmt.Errorf("relationship type is required")
	}

	// Create relationship using Apache AGE
	err := h.store.AddRelationship(input.FromID, input.ToID, input.Type, input.Properties)
	if err != nil {
		return nil, AddRelationshipOutput{}, fmt.Errorf("failed to create relationship: %w", err)
	}

	return nil, AddRelationshipOutput{
		Success: true,
		Message: fmt.Sprintf("Relationship %s created between memories %d and %d", input.Type, input.FromID, input.ToID),
	}, nil
}

// ExploreConnectionsInput defines input for explore_connections tool
type ExploreConnectionsInput struct {
	MemoryID int64 `json:"memory_id" jsonschema:"Starting memory ID"`
	MaxDepth int   `json:"max_depth,omitempty" jsonschema:"Maximum traversal depth (default: 2)"`
}

// ExploreConnectionsOutput defines output for explore_connections tool
type ExploreConnectionsOutput struct {
	Memories []storage.Memory `json:"memories,omitzero"`
	Count    int              `json:"count,omitzero"`
}

func (h *memoryHandler) handleExploreConnections(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input ExploreConnectionsInput,
) (*mcp.CallToolResult, ExploreConnectionsOutput, error) {
	if input.MemoryID == 0 {
		return nil, ExploreConnectionsOutput{}, fmt.Errorf("memory_id is required")
	}

	// Set default depth
	if input.MaxDepth == 0 {
		input.MaxDepth = 2
	}

	// Explore using Apache AGE graph traversal
	memories, err := h.store.ExploreConnections(input.MemoryID, input.MaxDepth)
	if err != nil {
		return nil, ExploreConnectionsOutput{}, fmt.Errorf("failed to explore connections: %w", err)
	}

	return nil, ExploreConnectionsOutput{
		Memories: memories,
		Count:    len(memories),
	}, nil
}

// AutoDetectRelationshipsInput defines input for auto_detect_relationships tool
type AutoDetectRelationshipsInput struct {
	MemoryID        int64   `json:"memory_id" jsonschema:"Memory ID to analyze for relationships"`
	MinSimilarity   float64 `json:"min_similarity,omitempty" jsonschema:"Minimum similarity for candidates (default: 0.5)"`
	MaxCandidates   int     `json:"max_candidates,omitempty" jsonschema:"Maximum candidates to analyze (default: 10)"`
	MinConfidence   float64 `json:"min_confidence,omitempty" jsonschema:"Minimum LLM confidence to create relationship (default: 0.7)"`
	DryRun          bool    `json:"dry_run,omitempty" jsonschema:"If true, return suggestions without creating relationships"`
}

// AutoDetectRelationshipsOutput defines output for auto_detect_relationships tool
type AutoDetectRelationshipsOutput struct {
	Suggestions       []RelationshipSuggestion `json:"suggestions"`
	RelationshipsCreated int                   `json:"relationships_created"`
	Message           string                   `json:"message"`
}

// RelationshipSuggestion represents a detected relationship
type RelationshipSuggestion struct {
	TargetID   int64   `json:"target_id"`
	Type       string  `json:"type"`
	Reason     string  `json:"reason"`
	Confidence float64 `json:"confidence"`
}

func (h *memoryHandler) handleAutoDetectRelationships(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input AutoDetectRelationshipsInput,
) (*mcp.CallToolResult, AutoDetectRelationshipsOutput, error) {
	if input.MemoryID == 0 {
		return nil, AutoDetectRelationshipsOutput{}, fmt.Errorf("memory_id is required")
	}

	// Set defaults
	if input.MinSimilarity == 0 {
		input.MinSimilarity = 0.5
	}
	if input.MaxCandidates == 0 {
		input.MaxCandidates = 10
	}
	if input.MinConfidence == 0 {
		input.MinConfidence = 0.7
	}

	// Get the source memory
	sourceMemory, err := h.store.GetMemoryByID(input.MemoryID)
	if err != nil {
		return nil, AutoDetectRelationshipsOutput{}, fmt.Errorf("failed to get source memory: %w", err)
	}

	// Generate embedding for the source memory
	sourceEmbedding, err := h.embeddings.Generate(sourceMemory.Text)
	if err != nil {
		return nil, AutoDetectRelationshipsOutput{}, fmt.Errorf("failed to generate embedding: %w", err)
	}

	// Find similar memories as candidates
	searchResults, err := h.store.SearchMemories(sourceEmbedding, input.MaxCandidates+1, input.MinSimilarity, "")
	if err != nil {
		return nil, AutoDetectRelationshipsOutput{}, fmt.Errorf("failed to search for candidates: %w", err)
	}

	// Filter out the source memory itself and convert to candidates
	var candidates []llm.CandidateMemory
	for _, result := range searchResults {
		if result.Memory.ID != input.MemoryID {
			candidates = append(candidates, llm.CandidateMemory{
				ID:         result.Memory.ID,
				Text:       result.Memory.Text,
				Similarity: result.Similarity,
			})
		}
	}

	if len(candidates) == 0 {
		return nil, AutoDetectRelationshipsOutput{
			Message: "No similar memories found for relationship analysis",
		}, nil
	}

	// Use LLM to analyze relationships
	llmSuggestions, err := h.llm.AnalyzeRelationships(sourceMemory.Text, sourceMemory.ID, candidates)
	if err != nil {
		return nil, AutoDetectRelationshipsOutput{}, fmt.Errorf("LLM analysis failed: %w", err)
	}

	// Convert to output format - ensure we always have an array (not nil)
	suggestions := make([]RelationshipSuggestion, 0)
	for _, s := range llmSuggestions {
		suggestions = append(suggestions, RelationshipSuggestion{
			TargetID:   s.TargetID,
			Type:       s.Type,
			Reason:     s.Reason,
			Confidence: s.Confidence,
		})
	}

	// Create relationships if not a dry run
	created := 0
	if !input.DryRun {
		for _, suggestion := range llmSuggestions {
			if suggestion.Confidence >= input.MinConfidence {
				props := map[string]interface{}{
					"reason":     suggestion.Reason,
					"confidence": suggestion.Confidence,
					"auto_detected": true,
				}

				err := h.store.AddRelationship(input.MemoryID, suggestion.TargetID, suggestion.Type, props)
				if err == nil {
					created++
				}
			}
		}
	}

	message := fmt.Sprintf("Found %d relationship suggestions", len(suggestions))
	if !input.DryRun {
		message = fmt.Sprintf("Created %d relationships from %d suggestions", created, len(suggestions))
	}

	return nil, AutoDetectRelationshipsOutput{
		Suggestions:          suggestions,
		RelationshipsCreated: created,
		Message:              message,
	}, nil
}
