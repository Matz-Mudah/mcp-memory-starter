"""MCP Memory Server - Main entry point"""
from mcp.server.fastmcp import FastMCP
from typing import Optional

from .config import config
from .storage.database import init_database
from .tools.store_memory import handle_store_memory
from .tools.search_memory import handle_search_memory

# Initialize database on startup
init_database(config.DB_PATH)
print(f"✅ Database initialized at: {config.DB_PATH}")

# Create MCP server
mcp = FastMCP("memory-server")

@mcp.tool()
async def store_memory(text: str, metadata: Optional[dict] = None) -> str:
    """
    Store a new memory with semantic embedding for future retrieval.

    Args:
        text: The memory text to store
        metadata: Optional metadata (tags, category, importance, etc.)

    Returns:
        Success message with memory ID
    """
    return await handle_store_memory({"text": text, "metadata": metadata})

@mcp.tool()
async def search_memory(
    query: str,
    limit: int = 5,
    min_similarity: float = 0.0
) -> str:
    """
    Search for relevant memories using semantic similarity.

    Args:
        query: The search query
        limit: Maximum number of results to return (default: 5)
        min_similarity: Minimum similarity score 0-1 (default: 0.0)

    Returns:
        Formatted search results
    """
    return await handle_search_memory({
        "query": query,
        "limit": limit,
        "minSimilarity": min_similarity
    })
