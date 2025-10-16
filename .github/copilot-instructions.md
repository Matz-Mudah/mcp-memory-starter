# MCP Memory Starter AI Agent Instructions

This is an educational TypeScript project teaching students to build AI memory systems using the Model Context Protocol (MCP). The project demonstrates semantic search, embedding storage, and AI-assistant integration.

## Project Architecture

### Core MCP Server Pattern
- **Entry Point**: `src/index.ts` - MCP server with stdio transport, handles `list_tools` and `call_tool` requests
- **Tool Registration**: Tools defined as objects with `name`, `description`, `inputSchema`, then handlers implement the logic
- **Error Handling**: Return `{ content: [...], isError: true }` format for MCP clients
- **Graceful Shutdown**: Handle SIGINT/SIGTERM, close database connections

### Memory Storage Architecture 
```
AI Client → MCP Protocol → Server → Embedding API → SQLite Database
         ↗ (LM Studio, Claude Desktop, Copilot Chat)
```

### Key Components
- **Database**: SQLite with `better-sqlite3`, stores `text + embedding + metadata`
- **Embeddings**: Generated via LM Studio API (`http://localhost:1234/v1`), typically 768-dimension vectors
- **Search**: Cosine similarity between query embedding and stored embeddings
- **Tools**: `store_memory(text, metadata?)` and `search_memory(query, limit?)`

## Development Workflows

### Building and Testing
```bash
npm run build          # TypeScript compilation to build/
npm run inspect        # Opens MCP Inspector for tool testing
npm run dev           # Watch mode development
```

### MCP Inspector Testing
- Use `npm run inspect` to test tools interactively
- Test `store_memory`: `{"text": "My favorite color is blue"}`
- Test `search_memory`: `{"query": "What color do I like?", "limit": 5}`

### AI Client Setup
- **Claude Desktop**: Edit `claude_desktop_config.json` with absolute paths to `build/index.js`
- **Copilot Chat**: Configure in `mcp.json` with `"type": "stdio"`
- **LM Studio**: Add server to MCP settings, requires local embedding model

## Project-Specific Patterns

### Configuration System
- Uses `dotenv` with `.env` file for API endpoints and model names
- Config validation in `loadConfig()` with meaningful error messages
- Default values: `http://localhost:1234/v1`, `nomic-embed-text`, 768 dimensions

### Error Handling Strategy
- MCP tools return formatted error messages, not thrown exceptions
- Database initialization checked before operations
- Embedding API failures handled gracefully with retry logic
- All console.error() output goes to AI client logs for debugging

### Database Schema
```sql
CREATE TABLE memories (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  text TEXT NOT NULL,
  embedding TEXT NOT NULL,        -- JSON stringified float array
  created_at TEXT NOT NULL,       -- ISO timestamp
  metadata TEXT                   -- JSON stringified object
)
```

### Embedding Workflow
1. Text → LM Studio API → Float array embedding
2. Store embedding as JSON string in SQLite
3. Search: Query → embedding → cosine similarity → ranked results
4. Results include similarity scores and original text

## File Structure Conventions

### Two-Track System
- **`examples/`**: Complete working implementations for reference
- **`starter-templates/`**: Scaffold with TODO comments for students
- **`docs/`**: Step-by-step learning materials with concepts and setup

### TypeScript Setup
- ES modules (`"type": "module"` in package.json)
- Node.js 20+ required for MCP SDK compatibility
- Build target: ES2022, output to `build/` directory
- Import paths use `.js` extensions (ES module requirement)

## Integration Points

### LM Studio Dependency
- Requires running LM Studio with embedding model loaded
- Default endpoint: `http://localhost:1234/v1/embeddings`
- Common models: `nomic-embed-text`, `all-MiniLM-L6-v2`
- API format matches OpenAI embeddings specification

### MCP Client Communication
- Uses stdio transport for process communication
- Clients spawn server as child process with node command
- Configuration paths must be absolute (not relative)
- Environment variables may not be available in client-spawned processes

### Cross-Platform Considerations
- Windows paths use double backslashes in JSON config files
- Database file paths created recursively if directories don't exist
- Node.js executable must be in PATH for MCP clients

## Common Development Issues

### Database Connection
- Call `initDatabase(config)` before any database operations
- Handle missing directories with `mkdirSync({ recursive: true })`
- Close connections on process exit to prevent locks

### Embedding Dimensions
- All embeddings must have same dimensions (768 for nomic-embed-text)
- Clear database when switching embedding models
- Validate embedding response before storage

### MCP Protocol Compliance
- Tool responses must use `content` array format: `[{ type: 'text', text: '...' }]`
- Use `console.error()` for logging (goes to client logs, not response)
- Handle both stdio and other transport types in production

## Student Learning Path
1. Examine working example in `examples/basic-typescript-example/`
2. Copy starter template and follow TODO comments
3. Test with MCP Inspector before connecting AI clients
4. Configure one AI client (Copilot Chat recommended for students)
5. Extend with custom features (metadata filtering, delete tool, etc.)

When helping students: emphasize understanding the semantic search concept, proper error handling, and MCP protocol compliance over complex features.