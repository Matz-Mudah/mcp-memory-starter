# Working Memory System Example - TypeScript

This is a **complete, working implementation** of an AI memory system using MCP (Model Context Protocol). Students can use this as a reference to understand what success looks like!

## What This Example Demonstrates

- ✅ Database initialization with SQLite
- ✅ Embedding generation using LM Studio
- ✅ Cosine similarity calculation for semantic search
- ✅ Memory storage with metadata support
- ✅ Semantic search that finds memories by meaning
- ✅ Complete MCP server setup
- ✅ Error handling and input validation

## Quick Start

### 1. Install Dependencies

```bash
npm install
```

### 2. Configure Environment

```bash
# Copy the example environment file
cp .env.example .env

# Make sure LM Studio is running on port 1234 with an embedding model!
```

### 3. Build the Project

```bash
npm run build
```

### 4. Test with MCP Inspector

```bash
npm run inspect
```

This opens a web interface where you can test the tools!

## How It Works

### Architecture

```
┌─────────────────────────────────┐
│   AI Client (Claude, LM Studio) │
└──────────────┬──────────────────┘
               │ MCP Protocol
               ↓
┌───────────────────────────────────┐
│     Your MCP Server (index.ts)    │
│  ┌─────────────────────────────┐  │
│  │  store_memory / search_memory│  │
│  └──────────┬──────────────────┘  │
│             ↓                      │
│  ┌──────────────────────────────┐ │
│  │  Embeddings (LM Studio API)  │ │
│  └──────────┬───────────────────┘ │
│             ↓                      │
│  ┌──────────────────────────────┐ │
│  │  SQLite Database             │ │
│  │  (stores text + embeddings)  │ │
│  └──────────────────────────────┘ │
└───────────────────────────────────┘
```

### Key Files Explained

#### [src/index.ts](src/index.ts)
Main MCP server entry point. Handles:
- Server initialization
- Tool registration
- Request routing

#### [src/storage/database.ts](src/storage/database.ts)
SQLite database operations:
- `initDatabase()` - Creates database and tables
- `storeMemory()` - Saves memory with embedding
- `searchMemories()` - Finds similar memories
- `getAllMemories()` - Retrieves all memories
- `deleteMemory()` - Removes a memory
- `clearAllMemories()` - Wipes the database

#### [src/storage/embeddings.ts](src/storage/embeddings.ts)
Embedding generation and similarity:
- `generateEmbedding()` - Calls LM Studio API
- `cosineSimilarity()` - Calculates vector similarity
- `generateEmbeddingsBatch()` - Batch processing

#### [src/tools/store-memory.ts](src/tools/store-memory.ts)
MCP tool for storing memories:
- Validates input
- Generates embedding
- Stores in database
- Returns success message

#### [src/tools/search-memory.ts](src/tools/search-memory.ts)
MCP tool for searching memories:
- Validates input
- Generates query embedding
- Searches for similar memories
- Formats and returns results

## Testing the System

### Option 1: MCP Inspector (Recommended for Development)

The easiest way to test your tools:

1. Run `npm run inspect`
2. Test `store_memory`:
   ```json
   {
     "text": "My favorite color is blue"
   }
   ```
3. Test `search_memory`:
   ```json
   {
     "query": "What color do I like?",
     "limit": 5
   }
   ```

### Option 2: GitHub Copilot Chat (Best for Students!)

Free for students via GitHub Student Pack!

> **📖 Full GitHub Copilot Guide:** See [docs/github-copilot-mcp-setup.md](../../docs/github-copilot-mcp-setup.md)

**Quick Setup:**
1. Edit `C:\Users\YourName\AppData\Roaming\Code\User\mcp.json`
2. Add your server with `"type": "stdio"`
3. Enable MCP in Copilot settings
4. Chat in VSCode!

**Get Student Access:** https://education.github.com/pack

### Option 3: Claude Code (VSCode Extension)

Use your MCP server directly in VSCode!

> **📖 Full Claude Code Guide:** See [docs/claude-code-mcp-setup.md](../../docs/claude-code-mcp-setup.md)

**Quick Setup:**
1. Run: `claude mcp add`
2. Or edit `~/.claude.json` with your server path
3. Reload MCP servers in VSCode
4. Chat with Claude in your editor!

**Config location:** `C:\Users\YourName\.claude.json` (Windows) or `~/.claude.json` (Mac/Linux)

### Option 3: LM Studio (Local AI Models)

Use your MCP server with local models - no API costs!

> **📖 Full LM Studio Guide:** See [docs/lm-studio-mcp-setup.md](../../docs/lm-studio-mcp-setup.md)

**Quick Steps:**
1. Open LM Studio → Chat tab → **Program** (right menu) → **Install**
2. Edit `mcp.json` and add your server path
3. Reload and chat with a local model!

### Option 4: Claude Desktop

> **📖 Full Claude Desktop Guide:** See [docs/mcp-setup-guide.md](../../docs/mcp-setup-guide.md)

**Step 1: Build the project**
```bash
npm run build
```

**Step 2: Find your Claude config file**

- **Windows:** `%APPDATA%\Claude\claude_desktop_config.json`
- **Mac:** `~/Library/Application Support/Claude/claude_desktop_config.json`

**Step 3: Add the MCP server to your config**

```json
{
  "mcpServers": {
    "mcp-memory-starter": {
      "command": "node",
      "args": [
        "C:\\personal\\mcp-memory-starter\\examples\\basic-typescript-example\\build\\index.js"
      ],
      "env": {}
    }
  }
}
```

> **Important:**
> - Use **absolute paths** (not relative)
> - On Windows, use double backslashes `\\` or forward slashes `/`
> - Replace the path with your actual project location

**Step 4: Restart Claude Desktop**

Close and reopen Claude Desktop completely.

**Step 5: Test it!**

Chat with Claude:
- "Store this memory: I love pizza"
- "What food do I like?"
- "Remember that my favorite language is TypeScript"
- "What programming language do I prefer?"

You should see Claude using the `store_memory` and `search_memory` tools!

## Common Issues

### "Database not initialized"
Make sure `initDatabase()` is called in [src/index.ts](src/index.ts:31)

### "Failed to generate embedding"
- Check that LM Studio is running (http://localhost:1234)
- Make sure an embedding model is loaded
- Verify the model name in `.env` matches

### "Embeddings must have same dimensions"
All embeddings must use the same model. Clear the database if you change models:
```bash
rm -rf data/
```

### Claude Desktop Issues

**Tools not appearing:**
1. Check Claude Desktop logs:
   - Windows: `%APPDATA%\Claude\logs\`
   - Mac: `~/Library/Logs/Claude/`
2. Verify the config path is absolute (not relative)
3. Make sure the build was successful (`npm run build`)
4. Restart Claude Desktop completely

**"Failed to start MCP server":**
- Check that Node.js is in your PATH
- Verify the path to `index.js` exists
- Try running manually: `node C:\path\to\build\index.js`

**Environment variables not working:**
- MCP servers started by Claude Desktop may not have access to LM Studio
- Make sure LM Studio is running on `http://localhost:1234`
- Check `.env` file is in the project root

## Understanding the Code

### How Embeddings Work

```typescript
// Text is converted to numbers that represent meaning
const text = "I love programming";
const embedding = await generateEmbedding(text, config);
// Result: [0.23, -0.15, 0.87, ...] (768 numbers)
```

### How Semantic Search Works

```typescript
// 1. Convert query to embedding
const queryEmbedding = await generateEmbedding("coding", config);

// 2. Compare to all stored embeddings
const results = searchMemories(queryEmbedding, 5);

// 3. Results are sorted by similarity
// "I love programming" will have high similarity!
```

### Cosine Similarity Math

```typescript
// Measures angle between two vectors
// Result: 0 (completely different) to 1 (identical)
const similarity = cosineSimilarity(embedding1, embedding2);
```

## Next Steps for Students

Now that you've seen a working example, try:

1. **Add metadata filtering** - Search by category or tag
2. **Implement deletion** - Add a `delete_memory` tool
3. **Add timestamps** - Filter by date range
4. **Improve search** - Add keyword + semantic hybrid search
5. **Build a UI** - Create a web interface to view memories

## Learn More

- [Understanding Embeddings](../../docs/01-concepts.md)
- [MCP Documentation](https://modelcontextprotocol.io/)
- [SQLite Tutorial](https://www.sqlitetutorial.net/)
- [TypeScript Handbook](https://www.typescriptlang.org/docs/)

---

**Questions?** Check the [main README](../../README.md) or ask your teacher!

**Good luck with your own implementation!** 🚀
