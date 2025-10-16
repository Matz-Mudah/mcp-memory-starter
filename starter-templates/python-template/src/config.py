"""Configuration management for MCP Memory Server"""
import os
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

class Config:
    """Application configuration"""

    # Embedding API
    EMBEDDING_BASE_URL = os.getenv('EMBEDDING_BASE_URL', 'http://localhost:1234/v1')
    EMBEDDING_MODEL = os.getenv('EMBEDDING_MODEL', 'text-embedding-embeddinggemma-300m-qat')

    # Database
    DB_PATH = os.getenv('DB_PATH', './data/memories.db')

config = Config()
