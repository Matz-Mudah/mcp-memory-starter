"""Search memory tool handler"""
from typing import Dict, Any
from ..storage.database import search_memories
from ..storage.embeddings import generate_embedding

async def handle_search_memory(arguments: Dict[str, Any]) -> str:
    """
    Handle search_memory tool requests.

    Args:
        arguments: Dictionary with 'query', optional 'limit' and 'minSimilarity'

    Returns:
        Formatted string with search results

    TODO: Implement this function!
    - Extract query, limit (default 5), and minSimilarity (default 0.0)
    - Generate embedding for the query
    - Search memories using search_memories()
    - Format results as a readable string
    - Return "No relevant memories found." if empty

    Example format:
    "Found 2 relevant memories:

    1. [Similarity: 87.3%] I love TypeScript
    2. [Similarity: 65.1%] Python is great"
    """
    pass
