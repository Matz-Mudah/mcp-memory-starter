#!/usr/bin/env node

/**
 * MCP Memory System - Main Entry Point
 *
 * This is the core MCP server that connects to Claude Desktop
 * and handles tool calls for memory storage and retrieval.
 */

import { Server } from '@modelcontextprotocol/sdk/server/index.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';
import {
  CallToolRequestSchema,
  ListToolsRequestSchema,
} from '@modelcontextprotocol/sdk/types.js';

import { loadConfig } from './config.js';
import { initDatabase, closeDatabase } from './storage/database.js';
import { storeMemoryTool, handleStoreMemory } from './tools/store-memory.js';
import { searchMemoryTool, handleSearchMemory } from './tools/search-memory.js';

/**
 * Main server initialization
 */
async function main() {
  try {
    // Load configuration from .env
    const config = loadConfig();

    // Initialize database
    initDatabase(config);

    // Create MCP server
    const server = new Server(
      {
        name: 'mcp-memory-system',
        version: '1.0.0',
      },
      {
        capabilities: {
          tools: {}, // We provide tools for AI to call
        },
      }
    );

    /**
     * Handle list_tools request
     *
     * This tells Claude Desktop what tools are available
     */
    server.setRequestHandler(ListToolsRequestSchema, async () => {
      return {
        tools: [storeMemoryTool, searchMemoryTool],
      };
    });

    /**
     * Handle call_tool request
     *
     * This is called when Claude wants to use one of our tools
     */
    server.setRequestHandler(CallToolRequestSchema, async (request) => {
      const { name, arguments: args } = request.params;

      try {
        let result: string;

        switch (name) {
          case 'store_memory':
            result = await handleStoreMemory(args as any, config);
            break;

          case 'search_memory':
            result = await handleSearchMemory(args as any, config);
            break;

          default:
            throw new Error(`Unknown tool: ${name}`);
        }

        return {
          content: [
            {
              type: 'text',
              text: result,
            },
          ],
        };
      } catch (error) {
        // Return errors in a format Claude can understand
        const errorMessage = error instanceof Error ? error.message : String(error);
        return {
          content: [
            {
              type: 'text',
              text: `Error: ${errorMessage}`,
            },
          ],
          isError: true,
        };
      }
    });

    // Connect to AI client via stdio
    const transport = new StdioServerTransport();
    await server.connect(transport);

    // Log startup (goes to AI client logs)
    console.error('MCP Memory System started successfully!');
    console.error(`Database: ${config.dbPath}`);
    console.error(`Embedding API: ${config.embeddingBaseUrl}`);
    console.error(`Model: ${config.embeddingModel}`);

    // Handle graceful shutdown
    process.on('SIGINT', () => {
      console.error('Shutting down...');
      closeDatabase();
      process.exit(0);
    });

    process.on('SIGTERM', () => {
      console.error('Shutting down...');
      closeDatabase();
      process.exit(0);
    });
  } catch (error) {
    console.error('Fatal error during startup:', error);
    process.exit(1);
  }
}

// Start the server
main();
