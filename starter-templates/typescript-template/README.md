# MCP Memory System - TypeScript Template

A starter template for building an AI memory system using the Model Context Protocol (MCP).

> **⚠️ This is a TEMPLATE with TODOs for students to implement!**
> **👀 Want to see a working example first?** Check out [examples/basic-typescript-example](../../examples/basic-typescript-example/)

## 🚀 Quick Start

### 1. Install Dependencies

```bash
npm install
```

### 2. Configure Environment

```bash
# Copy the example environment file
cp .env.example .env

# Edit .env with your settings (LM Studio should be running on port 1234)
```

### 3. Build the Project

```bash
npm run build
```

### 4. Test Your MCP Server

```bash
# Option A: Run directly
npm start

# Option B: Use MCP Inspector (recommended for development)
npm run inspect
```

## 📁 Project Structure

```
typescript-template/
├── src/
│   ├── index.ts              # Main MCP server entry point
│   ├── config.ts             # Configuration loader
│   ├── tools/
│   │   ├── store-memory.ts   # TODO: Implement store_memory tool
│   │   └── search-memory.ts  # TODO: Implement search_memory tool
│   ├── storage/
│   │   ├── database.ts       # SQLite database setup
│   │   └── embeddings.ts     # Embedding generation service
│   └── types.ts              # TypeScript type definitions
├── data/                     # Database files (created automatically)
├── package.json              # Dependencies and scripts
├── tsconfig.json             # TypeScript configuration
└── .env                      # Your configuration (don't commit!)
```

## 🎯 Your Mission

This template provides the basic structure. **You need to implement:**

### Phase 1: Basic Setup ✅
- [x] Project structure created
- [x] Dependencies installed
- [ ] Environment configured
- [ ] Server runs without errors

### Phase 2: Storage (Week 5-6)
- [ ] Database initialization
- [ ] `store_memory` tool implementation
- [ ] Embedding generation
- [ ] Data persistence

### Phase 3: Search (Week 7-8)
- [ ] `search_memory` tool implementation
- [ ] Vector similarity calculation
- [ ] Result ranking
- [ ] Test with various queries

### Phase 4: Polish (Week 9-10)
- [ ] Error handling
- [ ] Input validation
- [ ] Documentation
- [ ] Demo video

## 🛠️ Development Commands

```bash
# Build once
npm run build

# Build and watch for changes (auto-rebuild)
npm run dev

# Run the server
npm start

# Open MCP Inspector (visual testing tool)
npm run inspect
```

## 🔌 Testing Your MCP Server

### Option A: LM Studio (Recommended for Beginners)

LM Studio 0.3.29+ can directly use MCP tools!

1. Build your project: `npm run build`
2. In LM Studio, go to **Developer** tab
3. Add your MCP server (it needs an HTTP endpoint - see advanced setup)
4. Chat with a local model that can call your memory tools!

### Option B: Claude Desktop

1. Build your project: `npm run build`

2. Add to Claude Desktop config:

**Windows:** `%APPDATA%\Claude\claude_desktop_config.json`
**Mac:** `~/Library/Application Support/Claude/claude_desktop_config.json`

```json
{
  "mcpServers": {
    "memory-system": {
      "command": "node",
      "args": ["/path/to/your-project/build/index.js"]
    }
  }
}
```

3. Restart Claude Desktop

4. Test by asking: "Store this memory: My favorite color is blue"

### Option C: MCP Inspector (Best for Development)

```bash
npm run inspect
```

This opens a web interface where you can test your tools directly!

## 📚 Implementation Guides

Follow these docs in order:

1. [Understanding Concepts](../../docs/01-concepts.md) - Learn about embeddings and MCP
2. [Setup Guide](../../docs/02-setup-guide.md) - Install required tools
3. [First MCP Tool](../../docs/03-first-mcp-tool.md) - Build a simple tool
4. [Memory Storage](../../docs/04-memory-storage.md) - Implement persistence
5. [Semantic Search](../../docs/05-semantic-search.md) - Add vector search
6. [Final Project](../../docs/06-final-project.md) - Polish and present

## 🔍 Key Files to Implement

### `src/storage/embeddings.ts`
Generate embeddings from text using LM Studio:

```typescript
export async function generateEmbedding(text: string): Promise<number[]> {
  // TODO: Call LM Studio API
  // TODO: Return embedding array
}
```

### `src/storage/database.ts`
SQLite database operations:

```typescript
export function storeMemory(text: string, embedding: number[]): void {
  // TODO: Insert into database
}

export function searchMemories(queryEmbedding: number[], limit: number): Memory[] {
  // TODO: Find similar embeddings
  // TODO: Return top results
}
```

### `src/tools/store-memory.ts`
MCP tool for storing memories:

```typescript
{
  name: "store_memory",
  description: "Store a new memory",
  inputSchema: {
    type: "object",
    properties: {
      text: { type: "string" }
    },
    required: ["text"]
  }
}
```

### `src/tools/search-memory.ts`
MCP tool for searching memories:

```typescript
{
  name: "search_memory",
  description: "Search memories by meaning",
  inputSchema: {
    type: "object",
    properties: {
      query: { type: "string" },
      limit: { type: "number", default: 5 }
    },
    required: ["query"]
  }
}
```

## 🐛 Troubleshooting

### "Cannot find module"
Run `npm install` again

### "Port already in use"
Change the port or stop the conflicting process

### "Embeddings are null"
Make sure LM Studio server is running on port 1234

### "Database locked"
Close any other connections to the database file

## 💡 Tips for Success

1. **Start Simple** - Get basic storage working before adding search
2. **Test Incrementally** - Test each function as you write it
3. **Use Console Logs** - Debug with `console.error()` (goes to AI client logs)
4. **Check Types** - TypeScript will catch many errors early
5. **Read Error Messages** - They usually tell you exactly what's wrong!

## 📖 Additional Resources

- [MCP TypeScript SDK Docs](https://github.com/modelcontextprotocol/typescript-sdk)
- [Better-SQLite3 Docs](https://github.com/WiseLibs/better-sqlite3)
- [TypeScript Handbook](https://www.typescriptlang.org/docs/)
- [LM Studio API Reference](https://lmstudio.ai/docs/api)

## 🆘 Getting Help

1. **[Look at the working example](../../examples/basic-typescript-example/)** - See how it's done!
2. Check the [FAQ](../../docs/faq.md)
3. Review [common mistakes](../../docs/common-mistakes.md) (if available)
4. Ask your teacher
5. Compare your code to the complete example

---

**Good luck! You've got this!** 🚀

*Remember: Every expert was once a beginner who refused to give up!*
