# Adding Memory Storage

Learn how to store and retrieve memories using SQLite and embeddings.

## What You'll Build

By the end of this guide, your MCP server will be able to:
- Store text memories in a database
- Generate embeddings for semantic search
- Persist data across restarts

## Quick Reference: Working Example

👉 **[See Complete Implementation](../examples/basic-typescript-example/src/storage/)**

The example shows:
- ✅ Database initialization
- ✅ Embedding generation
- ✅ Memory storage with metadata

## Core Concepts

### 1. Embeddings

Embeddings convert text to numbers that represent meaning:

```typescript
"I love pizza" → [0.23, 0.87, -0.42, ...] // 768 numbers
"Pizza is great" → [0.21, 0.89, -0.38, ...] // Similar numbers!
```

**Why?** Similar meanings = Similar numbers = Semantic search!

### 2. Vector Storage

We store both the text AND its embedding:

```sql
CREATE TABLE memories (
  id INTEGER PRIMARY KEY,
  text TEXT,              -- "I love pizza"
  embedding TEXT,         -- "[0.23, 0.87, ...]"
  created_at TEXT,
  metadata TEXT
)
```

## Implementation Steps

### Step 1: Database Setup

File: `src/storage/database.ts`

**What to implement:**
1. Create database directory
2. Open SQLite connection
3. Create memories table
4. Add indexes for performance

**See the example:**
- [database.ts:21-48](../examples/basic-typescript-example/src/storage/database.ts)

### Step 2: Embedding Generation

File: `src/storage/embeddings.ts`

**What to implement:**
1. Call LM Studio API
2. Extract embedding array
3. Handle errors

**API endpoint:** `http://localhost:1234/v1/embeddings`

**See the example:**
- [embeddings.ts:36-64](../examples/basic-typescript-example/src/storage/embeddings.ts)

### Step 3: Storage Function

**What to implement:**
```typescript
export function storeMemory(
  text: string,
  embedding: number[],
  metadata?: Record<string, unknown>
): number {
  // 1. Validate database is initialized
  // 2. Convert embedding to JSON
  // 3. Insert into database
  // 4. Return the memory ID
}
```

**See the example:**
- [database.ts:53-77](../examples/basic-typescript-example/src/storage/database.ts)

## Testing Your Implementation

### 1. Test Database Initialization

```bash
npm run build
npm run inspect
```

Check that `data/memories.db` is created.

### 2. Test Embedding Generation

Make sure LM Studio is running with an embedding model:
1. Open LM Studio
2. Load an embedding model (e.g., "nomic-embed-text")
3. Start the server

### 3. Test Memory Storage

In MCP Inspector, call `store_memory`:
```json
{
  "text": "My favorite color is blue"
}
```

**Expected:** "Memory stored successfully with ID: 1"

### 4. Verify Database

```bash
sqlite3 data/memories.db "SELECT * FROM memories;"
```

You should see your stored memory!

## Common Issues

### "Database not initialized"
- Check `initDatabase()` is called in `src/index.ts`
- Verify the data directory was created

### "Failed to generate embedding"
- Is LM Studio running on port 1234?
- Is an embedding model loaded?
- Check the model name in `.env` matches

### "Cannot find module 'better-sqlite3'"
- Run `npm install`
- Rebuild: `npm run build`

## Code Comparison

### Template (TODOs)
```typescript
export function storeMemory(...): number {
  // TODO: Implement memory storage
  throw new Error('Not implemented');
}
```

### Working Example
```typescript
export function storeMemory(...): number {
  if (!db) throw new Error('Database not initialized');
  
  const embeddingJson = JSON.stringify(embedding);
  const stmt = db.prepare(`INSERT INTO...`);
  const result = stmt.run(text, embeddingJson, ...);
  
  return result.lastInsertRowid as number;
}
```

**💡 Tip:** Open both files side-by-side and copy the structure!

## Understanding the Flow

```
User → "Store: I love pizza"
  ↓
MCP Tool Handler
  ↓
Generate Embedding
  ↓  "I love pizza" → [0.23, 0.87, ...]
  ↓
Store in Database
  ↓  INSERT INTO memories (text, embedding, ...)
  ↓
Return Success
  ↓
"Memory stored with ID: 1"
```

## Next Steps

1. ✅ Database stores memories
2. ✅ Embeddings are generated
3. 🔍 **[Add Semantic Search](05-semantic-search.md)** - Find similar memories!

## Additional Resources

- [SQLite Tutorial](https://www.sqlitetutorial.net/)
- [Better-SQLite3 Docs](https://github.com/WiseLibs/better-sqlite3)
- [Understanding Embeddings](01-concepts.md#embeddings)

---

**Pro tip:** If you get stuck, compare your code with the [working example](../examples/basic-typescript-example/src/storage/) line by line! 🔍
