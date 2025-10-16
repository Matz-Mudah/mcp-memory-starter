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
 */
export function initDatabase(config: Config): void {
  // Create directory if it doesn't exist
  const dir = dirname(config.dbPath);
  if (!existsSync(dir)) {
    mkdirSync(dir, { recursive: true });
  }

  // Open database
  db = new Database(config.dbPath);

  // Create table
  db.exec(`
    CREATE TABLE IF NOT EXISTS memories (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      text TEXT NOT NULL,
      embedding TEXT NOT NULL,
      created_at TEXT NOT NULL,
      metadata TEXT
    )
  `);

  // Create index for faster searches
  db.exec(`
    CREATE INDEX IF NOT EXISTS idx_created_at ON memories(created_at)
  `);

  console.error('Database initialized successfully');
}

/**
 * Store a memory in the database
 */
export function storeMemory(
  text: string,
  embedding: number[],
  metadata?: Record<string, unknown>
): number {
  // Check database is initialized
  if (!db) {
    throw new Error('Database not initialized. Call initDatabase() first');
  }

  // Convert arrays/objects to JSON for storage
  const embeddingJson = JSON.stringify(embedding);
  const metadataJson = metadata ? JSON.stringify(metadata) : null;

  // Get current timestamp
  const createdAt = new Date().toISOString();

  // Use prepared statement for safe SQL
  const stmt = db.prepare(`
    INSERT INTO memories (text, embedding, created_at, metadata)
    VALUES (?, ?, ?, ?)
  `);
  const result = stmt.run(text, embeddingJson, createdAt, metadataJson);
  return result.lastInsertRowid as number;
}

/**
 * Search for similar memories
 *
 * This is where the magic happens! We:
 * 1. Get all memories from the database
 * 2. Calculate similarity between query and each memory
 * 3. Sort by similarity
 * 4. Return top results
 */
export function searchMemories(
  queryEmbedding: number[],
  limit: number = 5,
  minSimilarity: number = 0.0
): SearchResult[] {
  // Check database is initialized
  if (!db) {
    throw new Error('Database not initialized');
  }

  // Get all memories (we'll filter in code for now)
  const stmt = db.prepare('SELECT * FROM memories');
  const rows = stmt.all();

  // Calculate similarity for each memory
  const results: SearchResult[] = rows.map((row: any) => {
    const embedding = JSON.parse(row.embedding);
    const similarity = cosineSimilarity(queryEmbedding, embedding);
    return {
      memory: {
        id: row.id,
        text: row.text,
        embedding,
        createdAt: row.created_at,
        metadata: row.metadata ? JSON.parse(row.metadata) : undefined,
      },
      similarity,
    };
  });

  // Filter by minimum similarity
  const filtered = results.filter((r) => r.similarity >= minSimilarity);

  // Sort by similarity (highest first)
  filtered.sort((a, b) => b.similarity - a.similarity);

  // Return top N results
  return filtered.slice(0, limit);
}

/**
 * Get a memory by ID
 */
export function getMemoryById(id: number): Memory | null {
  if (!db) {
    throw new Error('Database not initialized');
  }

  // Use prepared statement
  const stmt = db.prepare('SELECT * FROM memories WHERE id = ?');
  const row = stmt.get(id) as any;

  // Return null if not found
  if (!row) return null;

  // Parse JSON fields and return Memory object
  return {
    id: row.id,
    text: row.text,
    embedding: JSON.parse(row.embedding),
    createdAt: row.created_at,
    metadata: row.metadata ? JSON.parse(row.metadata) : undefined,
  };
}

/**
 * Get all memories (for debugging)
 */
export function getAllMemories(limit: number = 100): Memory[] {
  if (!db) {
    throw new Error('Database not initialized');
  }

  // Get all memories ordered by most recent first
  const stmt = db.prepare(
    'SELECT * FROM memories ORDER BY created_at DESC LIMIT ?'
  );
  const rows = stmt.all(limit);

  // Parse and return memories
  return rows.map((row: any) => ({
    id: row.id,
    text: row.text,
    embedding: JSON.parse(row.embedding),
    createdAt: row.created_at,
    metadata: row.metadata ? JSON.parse(row.metadata) : undefined,
  }));
}

/**
 * Delete a memory by ID
 */
export function deleteMemory(id: number): boolean {
  if (!db) {
    throw new Error('Database not initialized');
  }

  const stmt = db.prepare('DELETE FROM memories WHERE id = ?');
  const result = stmt.run(id);
  return result.changes > 0;
}

/**
 * Clear all memories (use with caution!)
 */
export function clearAllMemories(): number {
  if (!db) {
    throw new Error('Database not initialized');
  }

  const stmt = db.prepare('DELETE FROM memories');
  const result = stmt.run();
  return result.changes;
}

/**
 * Close the database connection
 */
export function closeDatabase(): void {
  if (db) {
    db.close();
    db = null;
  }
}
