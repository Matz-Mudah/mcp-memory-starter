/**
 * Store Memory Tool
 *
 * This MCP tool allows AI assistants to store new memories
 */

import { Config } from '../types.js';
import { generateEmbedding } from '../storage/embeddings.js';
import { storeMemory } from '../storage/database.js';

/**
 * Tool definition for MCP
 *
 * This tells the AI what the tool does and what parameters it needs
 */
export const storeMemoryTool = {
  name: 'store_memory',
  description: 'Store a new memory with semantic embedding for future retrieval',
  inputSchema: {
    type: 'object',
    properties: {
      text: {
        type: 'string',
        description: 'The memory text to store',
      },
      metadata: {
        type: 'object',
        description: 'Optional metadata (tags, category, importance, etc.)',
      },
    },
    required: ['text'],
  },
};

/**
 * Handle store_memory tool calls
 *
 * @param args - Tool arguments from AI
 * @param config - Configuration
 * @returns Success message with memory ID
 *
 * TODO: Implement this function!
 *
 * Steps:
 * 1. Extract text and metadata from args
 * 2. Generate embedding for the text
 * 3. Store in database
 * 4. Return success message with the memory ID
 */
export async function handleStoreMemory(
  args: { text: string; metadata?: Record<string, unknown> },
  config: Config
): Promise<string> {
  // Validate input
  if (!args.text || args.text.trim().length === 0) {
    throw new Error('Memory text cannot be empty');
  }

  // Generate embedding
  const embedding = await generateEmbedding(args.text, config);

  // Store in database
  const memoryId = storeMemory(args.text, embedding, args.metadata);

  // Add debug logging if enabled
  if (config.debug) {
    console.error(
      `Stored memory ${memoryId}: "${args.text.substring(0, 50)}..."`
    );
  }

  // Return success message
  return `Memory stored successfully with ID: ${memoryId}`;
}
