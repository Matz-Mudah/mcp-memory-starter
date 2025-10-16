/**
 * Type definitions for the memory system
 *
 * These types help TypeScript understand our data structures
 * and catch errors before runtime!
 */

/**
 * Configuration loaded from .env file
 */
export interface Config {
  embeddingBaseUrl: string;
  embeddingModel: string;
  embeddingDimensions: number;
  dbPath: string;
  debug: boolean;
}

/**
 * A stored memory with its embedding
 */
export interface Memory {
  id: number;
  text: string;
  embedding: number[];
  createdAt: string;
  metadata?: Record<string, unknown>;
}

/**
 * Search result with similarity score
 */
export interface SearchResult {
  memory: Memory;
  similarity: number; // 0-1, higher is more similar
}

/**
 * Response from LM Studio embedding API
 */
export interface EmbeddingResponse {
  object: string;
  data: Array<{
    object: string;
    embedding: number[];
    index: number;
  }>;
  model: string;
  usage: {
    prompt_tokens: number;
    total_tokens: number;
  };
}
