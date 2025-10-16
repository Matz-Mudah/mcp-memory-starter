"""Embedding generation and similarity calculation"""
import requests
from typing import List
from ..config import config

def generate_embedding(text: str) -> List[float]:
    """
    Generate an embedding for the given text using LM Studio API.

    Args:
        text: The text to embed

    Returns:
        List of floats representing the embedding vector

    TODO: Implement this function!
    - Make a POST request to {config.EMBEDDING_BASE_URL}/embeddings
    - Send JSON: {"input": text, "model": config.EMBEDDING_MODEL}
    - Parse the response and extract the embedding array
    - Return the embedding

    Hint: response.json()['data'][0]['embedding']
    """
    pass


def cosine_similarity(a: List[float], b: List[float]) -> float:
    """
    Calculate cosine similarity between two vectors.

    Args:
        a: First vector
        b: Second vector

    Returns:
        Similarity score between -1 and 1 (higher = more similar)

    TODO: Implement this function!
    - Calculate dot product: sum(a[i] * b[i] for all i)
    - Calculate magnitude of a: sqrt(sum(a[i]^2 for all i))
    - Calculate magnitude of b: sqrt(sum(b[i]^2 for all i))
    - Return: dot_product / (magnitude_a * magnitude_b)

    Hint: You can use math.sqrt() or ** 0.5 for square root
    """
    pass
