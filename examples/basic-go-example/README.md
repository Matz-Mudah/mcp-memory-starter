# Basic Go Memory System 🧠

A simple, educational implementation of an AI memory system using Go 1.25 and the official MCP SDK v1.0.0.

## What You'll Learn

- **Go 1.25 Modern Patterns** - Latest language features (`omitzero`, clean APIs)
- **MCP Integration** - Building MCP servers with the official Go SDK
- **Semantic Search** - Vector embeddings and cosine similarity
- **SQLite** - Lightweight database for persistence
- **AI Integration** - Working with LM Studio for embeddings

## Features

✨ **Simple & Educational** - ~330 lines of well-commented code
💾 **SQLite Storage** - Single-file database, no Docker required
🔍 **Semantic Search** - Find memories by meaning, not keywords
🚀 **Modern Go** - Uses Go 1.25 features and best practices
🎯 **MCP Tools** - Two clean tools: `store_memory` and `search_memory`

## Requirements

- **Go 1.25+** - [Download here](https://go.dev/dl/)
- **LM Studio** - [Download here](https://lmstudio.ai/)
  - Load an embedding model (recommended: `text-embedding-embeddinggemma-300m-qat`)
  - Start the local server

## Quick Start

### 1. Setup Environment

```bash
# Copy example environment file
cp .env.example .env

# Edit .env if needed (defaults work with LM Studio)
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Build & Run

```bash
# Build the executable
go build -o memory-server

# On Windows, add .exe extension
go build -o memory-server.exe

# Run the server
./memory-server        # Linux/Mac
./memory-server.exe    # Windows
```

Or run directly:
```bash
go run main.go
```

## How It Works

### Architecture

```
┌─────────────┐
│ Claude Code │  (or other MCP client)
└──────┬──────┘
       │ stdio (MCP protocol)
       ▼
┌─────────────────┐
│  Go MCP Server  │
│  (this code)    │
└────┬──────┬─────┘
     │      │
     ▼      ▼
 ┌────┐  ┌────────┐
 │SQLite│ │LM Studio│
 │ DB   │ │Embeddings│
 └────┘  └────────┘
```

### Storage

- **Database**: `memories.db` (SQLite file)
- **Schema**: Simple table with text, embedding (JSON array), and timestamp
- **Embeddings**: Stored as JSON arrays of float64 values

### Search Algorithm

1. Generate embedding for search query
2. Load all memories from database
3. Calculate cosine similarity between query and each memory
4. Filter by minimum similarity threshold
5. Sort by similarity (descending)
6. Return top N results

## MCP Tools

### `store_memory`

Store text with semantic embedding for later retrieval.

**Input:**
```json
{
  "text": "The user loves Go programming"
}
```

**Output:**
```json
{
  "success": true,
  "message": "Memory stored successfully with ID 1",
  "id": 1
}
```

### `search_memory`

Search for relevant memories using semantic similarity. Returns results sorted by similarity score (highest first).

**Input:**
```json
{
  "query": "What does the user like?",
  "limit": 5
}
```

**Optional Parameters:**
- `min_similarity` (default: 0.0) - Filter results below this threshold. Usually not needed since results are sorted by relevance.

**Output:**
```json
{
  "results": [
    {
      "memory": {
        "id": 1,
        "text": "The user loves Go programming",
        "created_at": "2025-10-18T14:30:00Z"
      },
      "similarity": 0.87
    }
  ],
  "count": 1
}
```

## Configuration

Edit `.env` to customize:

```bash
# Database file path
DATABASE_PATH=memories.db

# LM Studio embedding API
EMBEDDING_BASE_URL=http://localhost:1234/v1
EMBEDDING_MODEL=text-embedding-embeddinggemma-300m-qat
EMBEDDING_API_KEY=not-needed

# Enable debug logging
DEBUG=false
```

## Connecting to Claude Code

Add to your Claude Code MCP settings (`%APPDATA%\Claude\claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "basic-go-memory": {
      "command": "C:\\path\\to\\basic-go-example\\memory-server.exe"
    }
  }
}
```

Restart Claude Code, and you'll have access to the memory tools!

## Code Structure

```go
main.go                 # Everything in one file for simplicity!
├── Config             # Environment configuration
├── Database           # SQLite initialization and schema
├── MCP Server         # Server setup with stdio transport
├── Tools              # store_memory and search_memory handlers
├── Embeddings         # LM Studio API client
└── Similarity         # Cosine similarity calculation
```

## Modern Go Features Used

### Go 1.25 JSON `omitzero`

Automatically omit zero values from JSON output:

```go
type Memory struct {
    ID        int64     `json:"id,omitzero"`         // Omit if 0
    Text      string    `json:"text"`                // Always include
    CreatedAt time.Time `json:"created_at,omitzero"` // Omit if zero time
}
```

### MCP SDK v1.0.0

Clean, idiomatic Go API:

```go
// Create server
server := mcp.NewServer(&mcp.Implementation{
    Name:    "basic-go-memory",
    Version: "1.0.0",
}, nil)

// Add tools with type-safe handlers
mcp.AddTool(server, &mcp.Tool{
    Name: "store_memory",
    Description: "Store a memory...",
}, handleStoreMemory)

// Run with stdio transport
server.Run(ctx, &mcp.StdioTransport{})
```

## Next Steps

Once you understand this basic example, check out:

- **[Advanced Go Example](../advanced-go-example/)** - Production patterns with Postgres + pgvector + Apache AGE
- **[TypeScript Example](../basic-typescript-example/)** - Same concepts in TypeScript
- **[Project Docs](../../docs/)** - Deep dives into concepts

## Troubleshooting

### "Failed to generate embedding"

- Make sure LM Studio is running
- Verify you have an embedding model loaded
- Check `EMBEDDING_BASE_URL` matches LM Studio's server address

### "Failed to open database"

- Ensure you have write permissions in the directory
- Check `DATABASE_PATH` is valid

### "Unknown tool"

- Restart Claude Code after adding MCP configuration
- Verify the executable path is correct
- Check Claude Code logs for errors

## Learning Resources

- **[Go Official Docs](https://go.dev/doc/)** - Learn Go basics
- **[MCP Documentation](https://modelcontextprotocol.io/)** - Understanding MCP
- **[SQLite Documentation](https://www.sqlite.org/docs.html)** - Database details
- **[Vector Embeddings Explained](https://www.youtube.com/watch?v=wjZofJX0v4M)** - Visual guide

---

**Built with Go 1.25 and the official MCP SDK v1.0.0** ✨

*Created for [mcp-memory-starter](../../README.md) educational project*
