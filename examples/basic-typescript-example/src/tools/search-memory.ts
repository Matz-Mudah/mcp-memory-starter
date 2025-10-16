/**
 * Search Memory Tool
 *
 * This MCP tool allows AI assistants to search for relevant memories
 */

import { Config } from '../types.js';
import { generateEmbedding } from '../storage/embeddings.js';
import { searchMemories } from '../storage/database.js';

/**
 * Tool definition for MCP
 */
export const searchMemoryTool = {
  name: 'search_memory',
  description:
    'Search for relevant memories using semantic similarity. Finds memories by meaning, not just keywords.',
  inputSchema: {
    type: 'object',
    properties: {
      query: {
        type: 'string',
        description: 'The search query (can be a question or statement)',
      },
      limit: {
        type: 'number',
        description: 'Maximum number of results to return',
        default: 5,
      },
      minSimilarity: {
        type: 'number',
        description: 'Minimum similarity score (0-1) for results',
        default: 0.0,
      },
    },
    required: ['query'],
  },
};

/**
 * Handle search_memory tool calls
 *
 * @param args - Tool arguments from AI
 * @param config - Configuration
 * @returns Formatted search results
 *
 * TODO: Implement this function!
 *
 * Steps:
 * 1. Extract query and options from args
 * 2. Generate embedding for the query
 * 3. Search database for similar memories
 * 4. Format and return results
 */
export async function handleSearchMemory(
  args: { query: string; limit?: number; minSimilarity?: number },
  config: Config
): Promise<string> {
  // Validate input
  if (!args.query || args.query.trim().length === 0) {
    throw new Error('Search query cannot be empty');
  }

  // Set defaults
  const limit = args.limit || 5;
  const minSimilarity = args.minSimilarity || 0.0;

  // Generate query embedding
  const queryEmbedding = await generateEmbedding(args.query, config);

  // Search database
  const results = searchMemories(queryEmbedding, limit, minSimilarity);

  // Handle no results
  if (results.length === 0) {
    return 'No relevant memories found.';
  }

  // Format results nicely
  const formatted = results
    .map((result, index) => {
      return `${index + 1}. [Similarity: ${(result.similarity * 100).toFixed(1)}%] ${result.memory.text}`;
    })
    .join('\n\n');

  // Add debug logging
  if (config.debug) {
    console.error(
      `Search for "${args.query}" returned ${results.length} results`
    );
  }

  return `Found ${results.length} relevant memories:\n\n${formatted}`;
}
