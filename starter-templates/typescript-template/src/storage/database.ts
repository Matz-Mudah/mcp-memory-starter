/**
 * Database operations using SQLite
 *
 * SQLite is a simple file-based database - perfect for learning!
 * This file handles storing and retrieving memories.
 */

import Database from 'better-sqlite3';
import { existsSync, mkdirSync } from 'fs';
import { dirname } from 'path';
import { Memory, SearchResult, Config } from '../types.js';
import { cosineSimilarity } from './embeddings.js';

let db: Database.Database | null = null;

/**
 * Initialize the database
 *
 * Creates the database file and table if they don't exist
 *
 * @param config - Configuration with database path
 *
 * TODO: Implement this function!
 *
 * Steps:
 * 1. Create the data directory if it doesn't exist
 * 2. Open/create the database file
 * 3. Create the memories table with appropriate columns
 *
 * Table Schema:
 * - id: INTEGER PRIMARY KEY AUTOINCREMENT
 * - text: TEXT (the memory content)
 * - embedding: TEXT (JSON array of numbers)
 * - created_at: TEXT (ISO timestamp)
 * - metadata: TEXT (optional JSON object)
 */
export function initDatabase(config: Config): void {
  // TODO: Implement database initialization
  //
  // Hint: Create directory if it doesn't exist
  //   const dir = dirname(config.dbPath);
  //   if (!existsSync(dir)) {
  //     mkdirSync(dir, { recursive: true });
  //   }
  //
  // Hint: Open database
  //   db = new Database(config.dbPath);
  //
  // Hint: Create table (use db.exec() to run SQL)
  //   db.exec(`
  //     CREATE TABLE IF NOT EXISTS memories (
  //       id INTEGER PRIMARY KEY AUTOINCREMENT,
  //       text TEXT NOT NULL,
  //       embedding TEXT NOT NULL,
  //       created_at TEXT NOT NULL,
  //       metadata TEXT
  //     )
  //   `);
  //
  // Hint: Create index for faster searches (optional but recommended)
  //   db.exec(`
  //     CREATE INDEX IF NOT EXISTS idx_created_at ON memories(created_at)
  //   `);

  throw new Error('initDatabase not implemented yet!');
}

/**
 * Store a memory in the database
 *
 * @param text - The memory text
 * @param embedding - The embedding vector
 * @param metadata - Optional additional data
 * @returns The ID of the inserted memory
 *
 * TODO: Implement this function!
 *
 * Steps:
 * 1. Check that database is initialized
 * 2. Convert embedding array to JSON string
 * 3. Insert into memories table
 * 4. Return the inserted row ID
 */
export function storeMemory(
  text: string,
  embedding: number[],
  metadata?: Record<string, unknown>
): number {
  // TODO: Implement memory storage
  //
  // Hint: Check database is initialized
  //   if (!db) {
  //     throw new Error('Database not initialized. Call initDatabase() first');
  //   }
  //
  // Hint: Convert arrays/objects to JSON for storage
  //   const embeddingJson = JSON.stringify(embedding);
  //   const metadataJson = metadata ? JSON.stringify(metadata) : null;
  //
  // Hint: Get current timestamp
  //   const createdAt = new Date().toISOString();
  //
  // Hint: Use prepared statement for safe SQL
  //   const stmt = db.prepare(`
  //     INSERT INTO memories (text, embedding, created_at, metadata)
  //     VALUES (?, ?, ?, ?)
  //   `);
  //   const result = stmt.run(text, embeddingJson, createdAt, metadataJson);
  //   return result.lastInsertRowid as number;

  throw new Error('storeMemory not implemented yet!');
}

/**
 * Search for similar memories
 *
 * This is where the magic happens! We'll:
 * 1. Get all memories from the database
 * 2. Calculate similarity between query and each memory
 * 3. Sort by similarity
 * 4. Return top results
 *
 * @param queryEmbedding - The embedding to search for
 * @param limit - Maximum number of results to return
 * @param minSimilarity - Minimum similarity score (0-1)
 * @returns Array of search results sorted by similarity
 *
 * TODO: Implement this function!
 */
export function searchMemories(
  queryEmbedding: number[],
  limit: number = 5,
  minSimilarity: number = 0.0
): SearchResult[] {
  // TODO: Implement memory search
  //
  // Hint: Check database is initialized
  //   if (!db) {
  //     throw new Error('Database not initialized');
  //   }
  //
  // Hint: Get all memories (we'll filter in code for now)
  //   const stmt = db.prepare('SELECT * FROM memories');
  //   const rows = stmt.all();
  //
  // Hint: Calculate similarity for each memory
  //   const results: SearchResult[] = rows.map((row: any) => {
  //     const embedding = JSON.parse(row.embedding);
  //     const similarity = cosineSimilarity(queryEmbedding, embedding);
  //     return {
  //       memory: {
  //         id: row.id,
  //         text: row.text,
  //         embedding,
  //         createdAt: row.created_at,
  //         metadata: row.metadata ? JSON.parse(row.metadata) : undefined
  //       },
  //       similarity
  //     };
  //   });
  //
  // Hint: Filter by minimum similarity
  //   const filtered = results.filter(r => r.similarity >= minSimilarity);
  //
  // Hint: Sort by similarity (highest first)
  //   filtered.sort((a, b) => b.similarity - a.similarity);
  //
  // Hint: Return top N results
  //   return filtered.slice(0, limit);

  throw new Error('searchMemories not implemented yet!');
}

/**
 * Get a memory by ID
 *
 * @param id - The memory ID
 * @returns The memory, or null if not found
 *
 * TODO: Implement this function!
 */
export function getMemoryById(id: number): Memory | null {
  // TODO: Implement get by ID
  //
  // Hint: Use prepared statement
  //   const stmt = db.prepare('SELECT * FROM memories WHERE id = ?');
  //   const row = stmt.get(id) as any;
  //
  // Hint: Return null if not found
  //   if (!row) return null;
  //
  // Hint: Parse JSON fields and return Memory object

  throw new Error('getMemoryById not implemented yet!');
}

/**
 * Get all memories (for debugging)
 *
 * @param limit - Maximum number of memories to return
 * @returns Array of memories
 *
 * TODO: Implement this function!
 */
export function getAllMemories(limit: number = 100): Memory[] {
  // TODO: Implement get all memories
  //
  // Hint: Similar to getMemoryById but gets all rows
  //   const stmt = db.prepare('SELECT * FROM memories ORDER BY created_at DESC LIMIT ?');
  //   const rows = stmt.all(limit);

  throw new Error('getAllMemories not implemented yet!');
}

/**
 * Delete a memory by ID
 *
 * @param id - The memory ID to delete
 * @returns True if deleted, false if not found
 *
 * EXTRA CREDIT: Implement this for bonus points!
 */
export function deleteMemory(id: number): boolean {
  // TODO (Optional): Implement memory deletion

  throw new Error('deleteMemory not implemented yet!');
}

/**
 * Clear all memories (use with caution!)
 *
 * @returns Number of memories deleted
 *
 * EXTRA CREDIT: Implement this for bonus points!
 */
export function clearAllMemories(): number {
  // TODO (Optional): Implement clear all memories

  throw new Error('clearAllMemories not implemented yet!');
}

/**
 * Close the database connection
 *
 * Call this when shutting down
 */
export function closeDatabase(): void {
  if (db) {
    db.close();
    db = null;
  }
}
