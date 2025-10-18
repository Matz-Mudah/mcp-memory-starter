package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"advanced-go-example/pkg/embeddings"
	"advanced-go-example/pkg/llm"
	"advanced-go-example/pkg/storage"
	"advanced-go-example/pkg/tools"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	ServerName    = "advanced-go-memory"
	ServerVersion = "1.0.0"
)

func main() {
	// Load configuration from environment
	config := loadConfig()

	// Initialize storage layer (Postgres + pgvector + Apache AGE)
	store, err := storage.NewPostgresStore(config.PostgresConfig)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	defer store.Close()

	// Initialize embedding client (LM Studio)
	embeddingClient := embeddings.NewClient(config.EmbeddingConfig)

	// Initialize LLM client for relationship detection
	llmClient := llm.NewClient(config.LLMConfig)

	// Create MCP server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    ServerName,
		Version: ServerVersion,
	}, nil)

	// Register memory tools
	tools.RegisterMemoryTools(server, store, embeddingClient, llmClient)

	// Handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Shutting down gracefully...")
		cancel()
	}()

	// Run server with stdio transport
	log.Printf("Starting %s v%s", ServerName, ServerVersion)
	log.Printf("Storage: PostgreSQL with pgvector + Apache AGE")
	log.Printf("Embeddings: %s", config.EmbeddingConfig.BaseURL)

	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// Config holds all application configuration
type Config struct {
	PostgresConfig  storage.PostgresConfig
	EmbeddingConfig embeddings.Config
	LLMConfig       llm.Config
	Debug           bool
}

// loadConfig loads configuration from environment variables
func loadConfig() Config {
	return Config{
		PostgresConfig: storage.PostgresConfig{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			User:     getEnv("POSTGRES_USER", "memoryuser"),
			Password: getEnv("POSTGRES_PASSWORD", "memorypass"),
			Database: getEnv("POSTGRES_DB", "memorydb"),
		},
		EmbeddingConfig: embeddings.Config{
			BaseURL: getEnv("EMBEDDING_BASE_URL", "http://localhost:1234/v1"),
			Model:   getEnv("EMBEDDING_MODEL", "text-embedding-embeddinggemma-300m-qat"),
			APIKey:  getEnv("EMBEDDING_API_KEY", "not-needed"),
		},
		LLMConfig: llm.Config{
			BaseURL: getEnv("LLM_BASE_URL", "http://localhost:1234/v1"),
			Model:   getEnv("LLM_MODEL", "qwen/qwen3-4b-2507"),
			APIKey:  getEnv("LLM_API_KEY", "not-needed"),
		},
		Debug: getEnv("DEBUG", "false") == "true",
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
