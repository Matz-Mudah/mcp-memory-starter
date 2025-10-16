/**
 * Embedding generation service
 *
 * This file handles communication with LM Studio to generate embeddings
 * from text. Embeddings are arrays of numbers that represent the meaning
 * of the text.
 */

import { Config, EmbeddingResponse } from '../types.js';

/**
 * Generate an embedding for the given text
 *
 * @param text - The text to convert to an embedding
 * @param config - Configuration with LM Studio URL and model
 * @returns Array of numbers representing the text's meaning
 *
 * TODO: Implement this function!
 *
 * Steps:
 * 1. Make a POST request to LM Studio's embeddings endpoint
 * 2. Send the text and model name in the request body
 * 3. Parse the response to extract the embedding array
 * 4. Return the embedding
 *
 * API Endpoint: {config.embeddingBaseUrl}/embeddings
 *
 * Request Body:
 * {
 *   "input": text,
 *   "model": config.embeddingModel
 * }
 *
 * Response will be EmbeddingResponse type (see types.ts)
 */
export async function generateEmbedding(
  text: string,
  config: Config
): Promise<number[]> {
  // TODO: Implement embedding generation
  //
  // Hint: Use the fetch API to make HTTP requests
  // Example:
  //   const response = await fetch(url, {
  //     method: 'POST',
  //     headers: { 'Content-Type': 'application/json' },
  //     body: JSON.stringify({ input: text, model: config.embeddingModel })
  //   });
  //
  // Hint: Check for errors!
  //   if (!response.ok) {
  //     throw new Error(`Failed to generate embedding: ${response.statusText}`);
  //   }
  //
  // Hint: Parse the JSON response
  //   const data: EmbeddingResponse = await response.json();
  //   return data.data[0].embedding;

  throw new Error('generateEmbedding not implemented yet!');
}

/**
 * Calculate cosine similarity between two embeddings
 *
 * Cosine similarity measures how similar two vectors are.
 * Returns a value between 0 (completely different) and 1 (identical)
 *
 * @param a - First embedding
 * @param b - Second embedding
 * @returns Similarity score (0-1)
 *
 * TODO: Implement this function!
 *
 * Formula: cosine_similarity = (a · b) / (|a| * |b|)
 * Where:
 *   a · b = dot product (sum of element-wise multiplication)
 *   |a| = magnitude (square root of sum of squares)
 *
 * Steps:
 * 1. Calculate dot product: sum(a[i] * b[i] for all i)
 * 2. Calculate magnitude of a: sqrt(sum(a[i]^2 for all i))
 * 3. Calculate magnitude of b: sqrt(sum(b[i]^2 for all i))
 * 4. Return: dotProduct / (magnitudeA * magnitudeB)
 */
export function cosineSimilarity(a: number[], b: number[]): number {
  // TODO: Implement cosine similarity calculation
  //
  // Hint: Make sure both arrays have the same length!
  //   if (a.length !== b.length) {
  //     throw new Error('Embeddings must have same dimensions');
  //   }
  //
  // Hint: Use array methods like .reduce() for calculations
  //   const dotProduct = a.reduce((sum, val, i) => sum + val * b[i], 0);
  //
  // Hint: Math.sqrt() for square root, Math.pow(x, 2) for squaring

  throw new Error('cosineSimilarity not implemented yet!');
}

/**
 * Batch generate embeddings for multiple texts
 *
 * This is more efficient than calling generateEmbedding multiple times
 * because it makes one API call instead of many
 *
 * @param texts - Array of texts to embed
 * @param config - Configuration
 * @returns Array of embeddings in the same order as input texts
 *
 * EXTRA CREDIT: Implement this for bonus points!
 */
export async function generateEmbeddingsBatch(
  texts: string[],
  config: Config
): Promise<number[][]> {
  // TODO (Optional): Implement batch embedding generation
  // This is similar to generateEmbedding, but sends an array of texts

  // For now, fall back to individual calls
  return Promise.all(texts.map((text) => generateEmbedding(text, config)));
}
