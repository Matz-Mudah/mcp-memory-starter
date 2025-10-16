"""Database operations for memory storage"""
import sqlite3
import json
from typing import List, Dict, Any, Optional
from pathlib import Path

db = None

def init_database(db_path: str) -> None:
    """
    Initialize the SQLite database with memories table.

    TODO: Implement this function!
    - Create the database directory if it doesn't exist
    - Connect to SQLite database
    - Create memories table with: id, text, embedding, timestamp
    - Store the connection in the global 'db' variable

    Hint: Use sqlite3.connect() and execute CREATE TABLE IF NOT EXISTS
    """
    pass


def store_memory(text: str, embedding: List[float], metadata: Optional[Dict] = None) -> int:
    """
    Store a memory with its embedding in the database.

    Args:
        text: The memory text
        embedding: The embedding vector (list of numbers)
        metadata: Optional metadata dictionary

    Returns:
        The ID of the inserted memory

    TODO: Implement this function!
    - Store text, embedding (as JSON), metadata (as JSON), and timestamp
    - Return the lastrowid from the insert

    Hint: Use json.dumps() to convert lists/dicts to JSON strings
    """
    pass


def search_memories(
    query_embedding: List[float],
    limit: int = 5,
    min_similarity: float = 0.0
) -> List[Dict[str, Any]]:
    """
    Search for similar memories using cosine similarity.

    Args:
        query_embedding: The query embedding vector
        limit: Maximum number of results
        min_similarity: Minimum similarity score (0-1)

    Returns:
        List of dicts with 'memory' and 'similarity' keys

    TODO: Implement this function!
    - Fetch all memories from database
    - Calculate similarity between query_embedding and each memory's embedding
    - Filter by min_similarity
    - Sort by similarity (highest first)
    - Return top 'limit' results

    Hint: Use json.loads() to parse JSON strings back to lists
    """
    pass
