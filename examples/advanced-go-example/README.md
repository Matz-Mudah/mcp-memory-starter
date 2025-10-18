## Advanced Go Memory System 🚀

**Production-ready AI memory with PostgreSQL + pgvector + Apache AGE**

A sophisticated implementation showcasing real-world patterns: vector similarity search AND graph relationships in a single PostgreSQL database!

## What Makes This Advanced?

### 🎯 Triple-Threat Database
- **PostgreSQL** - Industry-standard relational database
- **pgvector** - High-performance vector similarity search
- **Apache AGE** - Graph database with Cypher query language

### 🏗️ Production Architecture
- Clean package structure (`cmd/`, `pkg/`)
- Dependency injection for testability
- Connection pooling and resource management
- Graceful shutdown handling
- Modern Go 1.25 patterns

### 🔥 Advanced Features
1. **Vector Search** - Semantic similarity using pgvector's IVFFlat index
2. **Graph Relationships** - Link memories with typed relationships
3. **Graph Traversal** - Explore connected memories via Apache AGE
4. **AI-Powered Relationship Detection** - LLM automatically suggests and creates relationships
5. **Metadata Support** - Group IDs, importance scores, timestamps
6. **Docker Orchestration** - One-command database setup

## Requirements

- **Go 1.25+** - [Download here](https://go.dev/dl/)
- **Docker & Docker Compose** - [Get Docker](https://www.docker.com/)
- **LM Studio** - [Download here](https://lmstudio.ai/)
  - Load an embedding model (recommended: `text-embedding-embeddinggemma-300m-qat`)
  - Start the local server

## Quick Start

### 1. Start the Database

```bash
# Start PostgreSQL with pgvector + Apache AGE
docker-compose up -d

# Wait for database to be ready (watch for "database system is ready to accept connections")
docker-compose logs -f postgres
```

The migrations will run automatically, setting up:
- `pgvector` extension
- `Apache AGE` extension
- `memories` table with vector column
- `memory_graph` AGE graph
- Indexes for performance

### 2. Configure Environment

```bash
# Copy example environment
cp .env.example .env

# Defaults work if you used docker-compose!
```

### 3. Build & Run

```bash
# Build the server
go build -o memory-server ./cmd/server

# On Windows, add .exe extension
go build -o memory-server.exe ./cmd/server

# Run it
./memory-server        # Linux/Mac
./memory-server.exe    # Windows
```

Or run directly:
```bash
go run ./cmd/server
```

## Architecture

### Package Structure

```
advanced-go-example/
├── cmd/
│   └── server/
│       └── main.go           # Entry point, config loading
├── pkg/
│   ├── storage/
│   │   └── postgres.go       # PostgreSQL + pgvector + AGE
│   ├── embeddings/
│   │   └── client.go         # LM Studio embedding client
│   └── tools/
│       └── memory_tools.go   # MCP tool handlers
├── migrations/
│   └── 001_init.sql          # Database schema
├── docker-compose.yml        # PostgreSQL setup
└── .env.example              # Configuration template
```

### Data Flow

```
User Query
    ↓
MCP Client (Claude Code)
    ↓ (stdio)
Go MCP Server
    ↓
┌────────────────┐
│  Tool Handler  │
└────┬──────┬────┘
     │      │
     ▼      ▼
┌─────────┐ ┌──────────┐
│ Postgres│ │LM Studio │
│+pgvector│ │Embeddings│
│+AGE     │ └──────────┘
└─────────┘
```

## MCP Tools

### 1. `store_memory`

Store a memory with vector embedding and optional metadata. By default, automatically detects and creates relationships with similar memories using LLM analysis.

**Input:**
```json
{
  "text": "The user loves building with Go and PostgreSQL",
  "group_id": "technical_preferences",
  "auto_detect_relationships": true
}
```

**Output:**
```json
{
  "success": true,
  "message": "Memory stored with ID 1 and 2 relationships auto-created",
  "id": 1,
  "relationships_created": 2
}
```

**Note:** Set `auto_detect_relationships: false` to disable automatic relationship detection. Requires an LLM model loaded in LM Studio.

### 2. `search_memories` 🔍 GRAPH-ENHANCED!

**Hybrid search** combining vector similarity (pgvector) with graph traversal (Apache AGE).

The search automatically:
1. Finds semantically similar memories using vector search
2. Traverses graph relationships from top results (1 hop)
3. Includes connected memories in results

This means you get both semantically similar memories AND explicitly related memories in one search!

**Input:**
```json
{
  "query": "What technologies does the user like?",
  "limit": 5,
  "group_id": "technical_preferences"
}
```

**Optional Parameters:**
- `min_similarity` (default: 0.0) - Filter results below this threshold. Usually not needed since results are sorted by relevance.
- `group_id` - Filter results to a specific group

**Output:**
```json
{
  "results": [
    {
      "memory": {
        "id": 1,
        "text": "PostgreSQL is a powerful database",
        "group_id": "databases"
      },
      "similarity": 0.87
    },
    {
      "memory": {
        "id": 2,
        "text": "pgvector enables vector search in PostgreSQL",
        "group_id": "databases"
      },
      "similarity": 0,
      "via_relationship": true,
      "relationship_hops": 1
    }
  ],
  "count": 2
}
```

**Graph-Enhanced Fields:**
- `via_relationship` - True if memory was found via graph traversal (not vector search)
- `relationship_hops` - Number of hops from the vector search result (1 = directly connected)
- `similarity` - 0 for graph-discovered memories (no vector similarity score)

### 3. `add_relationship` ✨ NEW!

Create graph relationships between memories (Apache AGE).

**Input:**
```json
{
  "from_id": 1,
  "to_id": 2,
  "type": "RELATES_TO",
  "properties": {
    "strength": 0.9,
    "reason": "Both about programming"
  }
}
```

**Output:**
```json
{
  "success": true,
  "message": "Relationship RELATES_TO created between memories 1 and 2"
}
```

### 4. `explore_connections` ✨ NEW!

Find related memories via graph traversal (Apache AGE).

**Input:**
```json
{
  "memory_id": 1,
  "max_depth": 2
}
```

**Output:**
```json
{
  "memories": [
    {
      "id": 2,
      "text": "PostgreSQL is great for relational and vector data",
      "group_id": "technical_preferences"
    },
    {
      "id": 3,
      "text": "Go's concurrency model is elegant",
      "group_id": "technical_preferences"
    }
  ],
  "count": 2
}
```

### 5. `auto_detect_relationships` 🤖 AI-POWERED!

Automatically detect and create relationships using LLM analysis. This tool uses semantic similarity to find candidate memories, then analyzes them with an LLM to suggest meaningful relationships.

**Input:**
```json
{
  "memory_id": 1,
  "min_similarity": 0.5,
  "max_candidates": 10,
  "min_confidence": 0.7,
  "dry_run": false
}
```

**Output:**
```json
{
  "suggestions": [
    {
      "target_id": 2,
      "type": "RELATES_TO",
      "reason": "Both discuss PostgreSQL capabilities",
      "confidence": 0.85
    },
    {
      "target_id": 3,
      "type": "SIMILAR_TO",
      "reason": "Both describe Go programming features",
      "confidence": 0.78
    }
  ],
  "relationships_created": 2,
  "message": "Created 2 relationships from 2 suggestions"
}
```

**Features:**
- Finds semantically similar memories using vector search
- Uses LLM to classify relationship types (RELATES_TO, BUILDS_ON, CONTRADICTS, etc.)
- Provides confidence scores for each suggestion
- Dry-run mode to preview suggestions without creating relationships
- Auto-detection can also run automatically when storing memories

## How It Works

### Vector Search (pgvector)

1. **Store**: Text → Embedding (768-dim vector) → PostgreSQL `vector` column
2. **Index**: IVFFlat index for fast approximate nearest neighbor search
3. **Search**: Query → Embedding → Cosine similarity (`<=>` operator)
4. **Results**: Ordered by similarity, filtered by threshold

### Graph Relationships (Apache AGE)

1. **Store**: Create graph nodes and edges using Cypher
2. **Connect**: Link memories with typed relationships
3. **Traverse**: Follow edges to find related memories
4. **Query**: Use Cypher pattern matching for complex queries

### Combined Power 🔥

```
Example Workflow:

1. Store: "The user loves Go programming"
   → Vector embedding stored
   → Graph node created (memory ID 1)

2. Store: "Go is great for concurrent systems"
   → Vector embedding stored
   → Graph node created (memory ID 2)

3. Add Relationship:
   → Create edge: (1)-[RELATES_TO]->(2)

4. Search: "What does the user like?"
   → Vector search finds memory 1 (high similarity)

5. Explore Connections from memory 1:
   → Graph traversal finds memory 2
   → Return enriched results!
```

## Configuration

### Environment Variables

```bash
# PostgreSQL (matches docker-compose.yml)
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=memoryuser
POSTGRES_PASSWORD=memorypass
POSTGRES_DB=memorydb

# LM Studio Embeddings
EMBEDDING_BASE_URL=http://localhost:1234/v1
EMBEDDING_MODEL=text-embedding-embeddinggemma-300m-qat
EMBEDDING_API_KEY=not-needed

# Optional
DEBUG=false
```

### Docker Compose

The included `docker-compose.yml` uses the official Apache AGE image, which includes:
- PostgreSQL 16
- pgvector extension
- Apache AGE extension
- Automatic migration on startup

## Database Schema

### Memories Table

```sql
CREATE TABLE memories (
    id BIGSERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    embedding vector(768),        -- pgvector column!
    group_id TEXT,
    importance FLOAT DEFAULT 1.0,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Indexes for performance
CREATE INDEX idx_memories_embedding
    ON memories USING ivfflat (embedding vector_cosine_ops);
CREATE INDEX idx_memories_group_id ON memories(group_id);
```

### Graph (Apache AGE)

```cypher
// Graph created in 'memory_graph'
// Nodes: Memory {id: <memory_id>}
// Edges: Various relationship types (RELATES_TO, SIMILAR_TO, etc.)
```

## Production Patterns Demonstrated

### 1. Clean Architecture
- Separation of concerns (storage, embeddings, tools)
- Interface-based design for testability
- Dependency injection

### 2. Database Best Practices
- Connection pooling
- Prepared statements (via lib/pq)
- Index optimization
- Transaction support ready

### 3. Error Handling
- Wrapped errors with context
- Graceful degradation
- Resource cleanup

### 4. Modern Go Features
- Go 1.25 `omitzero` JSON tags
- Structured configuration
- Context-aware operations

## Connecting to Claude Code

Add to `%APPDATA%\Claude\claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "advanced-go-memory": {
      "command": "C:\\path\\to\\advanced-go-example\\memory-server.exe",
      "env": {
        "POSTGRES_HOST": "localhost",
        "POSTGRES_PORT": "5432",
        "POSTGRES_USER": "memoryuser",
        "POSTGRES_PASSWORD": "memorypass",
        "POSTGRES_DB": "memorydb"
      }
    }
  }
}
```

## Troubleshooting

### Search Returns No Results (Small Datasets)

**Problem**: Vector search returns 0 results even though memories exist.

**Cause**: The IVFFlat index requires ~1000+ rows to build proper clusters. With small datasets, it may return no results.

**Solution**: The index is commented out by default in migrations. pgvector will use sequential scans which work fine for small datasets (< 10,000 rows). Once you have sufficient data, you can add the index:

```sql
CREATE INDEX idx_memories_embedding
  ON memories USING ivfflat (embedding vector_cosine_ops)
  WITH (lists = 100);
```

### Database Connection Fails

```bash
# Check if PostgreSQL is running
docker-compose ps

# View logs
docker-compose logs postgres

# Restart
docker-compose restart postgres
```

### Migrations Didn't Run

```bash
# Stop and remove volumes
docker-compose down -v

# Start fresh
docker-compose up -d
```

### pgvector Extension Error

The Apache AGE Docker image includes pgvector. If you're using a different image:

```sql
CREATE EXTENSION IF NOT EXISTS vector;
```

### Apache AGE Query Errors

Ensure AGE is in your search path:

```sql
SET search_path = ag_catalog, "$user", public;
```

## Performance Tuning

### Vector Index Tuning

```sql
-- Adjust lists parameter (default: 100)
-- Higher = better recall, slower build
-- Lower = faster build, lower recall
CREATE INDEX idx_memories_embedding
    ON memories USING ivfflat (embedding vector_cosine_ops)
    WITH (lists = 100);
```

### Connection Pool

Edit `pkg/storage/postgres.go`:

```go
db.SetMaxOpenConns(25)   // Max concurrent connections
db.SetMaxIdleConns(5)     // Idle connections to keep
db.SetConnMaxLifetime(5 * time.Minute)
```

## Learning Path

### Compared to Basic Example

| Feature | Basic | Advanced |
|---------|-------|----------|
| Database | SQLite | PostgreSQL |
| Vector Search | In-memory cosine | pgvector IVFFlat |
| Graph Support | ❌ | ✅ Apache AGE |
| AI Relationship Detection | ❌ | ✅ LLM-powered |
| Structure | Single file | Multi-package |
| Deployment | Binary only | Docker Compose |
| Tools | 2 | 5 |

### Next Steps

1. **Add Tests** - Unit tests for each package
2. **Add Metrics** - Prometheus metrics for monitoring
3. **Add Caching** - Redis for frequently accessed data
4. **Add Migrations** - Use migrate/golang-migrate
5. **Add HTTP API** - REST endpoints alongside MCP

## Resources

- **[PostgreSQL](https://www.postgresql.org/)** - Main database
- **[pgvector](https://github.com/pgvector/pgvector)** - Vector extension
- **[Apache AGE](https://age.apache.org/)** - Graph extension
- **[MCP Go SDK](https://github.com/modelcontextprotocol/go-sdk)** - Official SDK
- **[lib/pq](https://github.com/lib/pq)** - PostgreSQL driver

---

**This example demonstrates production-ready patterns suitable for real-world AI applications!** 🚀

*Created for [mcp-memory-starter](../../README.md) educational project*
