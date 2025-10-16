# Implementing Semantic Search

Learn how to find memories by meaning, not just keywords!

## What You'll Build

A search system that understands concepts:
- Search "favorite food" → finds "I love pizza"
- Search "programming language" → finds "TypeScript is great"
- Search by meaning, not exact words!

## Quick Reference

👉 **[See Complete Implementation](../examples/basic-typescript-example/src/)**

Check out:
- [embeddings.ts:90-109](../examples/basic-typescript-example/src/storage/embeddings.ts) - Cosine similarity
- [database.ts:88-126](../examples/basic-typescript-example/src/storage/database.ts) - Search function
- [search-memory.ts:55-94](../examples/basic-typescript-example/src/tools/search-memory.ts) - Tool handler

## How Semantic Search Works

### Traditional Search (Keywords)
```
Stored: "I love pizza"
Search: "pizza" → ✅ Found (exact match)
Search: "favorite food" → ❌ Not found (different words)
```

### Semantic Search (Meaning)
```
Stored: "I love pizza" → embedding: [0.23, 0.87, ...]
Search: "favorite food" → embedding: [0.21, 0.89, ...]
                          Similar numbers! → ✅ Found!
```

## Core Algorithm: Cosine Similarity

Measures how similar two embeddings are (0 = different, 1 = identical):

```typescript
function cosineSimilarity(a: number[], b: number[]): number {
  // 1. Dot product: sum(a[i] * b[i])
  const dotProduct = a.reduce((sum, val, i) => sum + val * b[i], 0);
  
  // 2. Magnitudes
  const magA = Math.sqrt(a.reduce((sum, val) => sum + val * val, 0));
  const magB = Math.sqrt(b.reduce((sum, val) => sum + val * val, 0));
  
  // 3. Similarity score
  return dotProduct / (magA * magB);
}
```

## Implementation Steps

### Step 1: Implement Cosine Similarity

File: `src/storage/embeddings.ts`

**The template has hints!** Follow the TODOs:
```typescript
export function cosineSimilarity(a: number[], b: number[]): number {
  // TODO: Check dimensions match
  // TODO: Calculate dot product
  // TODO: Calculate magnitudes
  // TODO: Return similarity
}
```

**See the working version:**
- [embeddings.ts:90-109](../examples/basic-typescript-example/src/storage/embeddings.ts)

### Step 2: Implement Search Function

File: `src/storage/database.ts`

**What it does:**
1. Get all memories from database
2. Calculate similarity for each memory
3. Filter by minimum similarity
4. Sort by similarity (highest first)
5. Return top N results

**The template guides you:**
```typescript
export function searchMemories(
  queryEmbedding: number[],
  limit: number = 5,
  minSimilarity: number = 0.0
): SearchResult[] {
  // TODO: Get all memories
  // TODO: Calculate similarity for each
  // TODO: Filter and sort
  // TODO: Return top results
}
```

**See the working version:**
- [database.ts:88-126](../examples/basic-typescript-example/src/storage/database.ts)

### Step 3: Create Search Tool Handler

File: `src/tools/search-memory.ts`

**What it does:**
1. Generate embedding for search query
2. Call searchMemories()
3. Format results nicely
4. Return to user

**See the working version:**
- [search-memory.ts:55-94](../examples/basic-typescript-example/src/tools/search-memory.ts)

## Testing Semantic Search

### 1. Store Some Memories

```bash
npm run inspect
```

Call `store_memory` multiple times:
```json
{"text": "My favorite programming language is TypeScript"}
{"text": "I love building AI applications"}  
{"text": "Pizza is the best food"}
{"text": "I enjoy working with Node.js"}
```

### 2. Test Semantic Search

Call `search_memory` with different queries:

**Query 1: Programming**
```json
{
  "query": "What language do I like to code in?",
  "limit": 3
}
```
**Expected:** Finds TypeScript and Node.js memories (high similarity!)

**Query 2: Food**
```json
{
  "query": "What do I like to eat?",
  "limit": 2
}
```
**Expected:** Finds pizza memory

**Query 3: AI/Tech**
```json
{
  "query": "What kind of software does the user build?",
  "limit": 2
}
```
**Expected:** Finds "AI applications" memory

## Understanding Similarity Scores

The tool shows similarity as percentages:

```
1. [Similarity: 87.3%] My favorite programming language is TypeScript
2. [Similarity: 65.1%] I enjoy working with Node.js
3. [Similarity: 23.4%] Pizza is the best food
```

**What it means:**
- **>80%:** Very relevant
- **60-80%:** Related
- **40-60%:** Somewhat related
- **<40%:** Probably not what you're looking for

## Performance Optimization

### Current Implementation (Simple)
```typescript
// Get ALL memories, calculate similarity for each
const rows = db.prepare('SELECT * FROM memories').all();
```

**Works great for:**
- Learning and development
- Small datasets (<1000 memories)

### For Production (Advanced)
Use a proper vector database:
- **Qdrant** - Fast vector search
- **ChromaDB** - Python-friendly
- **Weaviate** - Kubernetes-ready

## Code Comparison

### Template
```typescript
export function searchMemories(...): SearchResult[] {
  // TODO: Implement memory search
  throw new Error('Not implemented');
}
```

### Working Example
```typescript
export function searchMemories(...): SearchResult[] {
  const stmt = db.prepare('SELECT * FROM memories');
  const rows = stmt.all();
  
  const results = rows.map(row => ({
    memory: {...},
    similarity: cosineSimilarity(queryEmbedding, rowEmbedding)
  }));
  
  return results
    .filter(r => r.similarity >= minSimilarity)
    .sort((a, b) => b.similarity - a.similarity)
    .slice(0, limit);
}
```

## Common Issues

### All Results Have Same Similarity
- Check cosine similarity implementation
- Make sure embeddings are different (not all zeros)

### No Results Found
- Lower `minSimilarity` to 0
- Check embeddings are being stored correctly
- Verify embedding model is the same for storage and search

### Slow Performance
- Normal for SQLite with basic search
- Consider vector database for >1000 memories
- Add limit to queries

## Testing Workflow

1. **Store test data** - Add diverse memories
2. **Test obvious queries** - Should work perfectly
3. **Test semantic queries** - Different words, same meaning
4. **Test edge cases** - Unrelated queries
5. **Check similarity scores** - Make sense?

## Next Steps

1. ✅ Cosine similarity works
2. ✅ Search finds relevant memories
3. ✅ Similarity scores look good
4. 🎨 **[Polish & Present](06-final-project.md)** - Finish your project!

## Additional Resources

- [Understanding Cosine Similarity](https://en.wikipedia.org/wiki/Cosine_similarity)
- [Vector Search Explained](https://www.pinecone.io/learn/vector-similarity/)
- [Working Example Code](../examples/basic-typescript-example/src/)

---

**Pro tip:** The best way to understand semantic search is to experiment! Try different queries and see what it finds. 🔍
