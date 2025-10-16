/**
 * Configuration loader
 *
 * Loads settings from .env file and provides defaults
 */

import dotenv from 'dotenv';
import { Config } from './types.js';

// Load environment variables from .env file
dotenv.config();

/**
 * Load and validate configuration
 */
export function loadConfig(): Config {
  const config: Config = {
    embeddingBaseUrl: process.env.EMBEDDING_BASE_URL || 'http://localhost:1234/v1',
    embeddingModel: process.env.EMBEDDING_MODEL || 'nomic-embed-text',
    embeddingDimensions: parseInt(process.env.EMBEDDING_DIMENSIONS || '768'),
    dbPath: process.env.DB_PATH || './data/memories.db',
    debug: process.env.DEBUG === 'true',
  };

  // Validate configuration
  if (!config.embeddingBaseUrl) {
    throw new Error('EMBEDDING_BASE_URL is required in .env file');
  }

  if (!config.embeddingModel) {
    throw new Error('EMBEDDING_MODEL is required in .env file');
  }

  if (isNaN(config.embeddingDimensions) || config.embeddingDimensions <= 0) {
    throw new Error('EMBEDDING_DIMENSIONS must be a positive number');
  }

  if (config.debug) {
    console.error('Configuration loaded:', JSON.stringify(config, null, 2));
  }

  return config;
}
