# Python MCP Memory Server - Starter Template

A beginner-friendly template for building an AI memory system with Python.

## What You'll Build

A memory system that:
- Stores information with embeddings (semantic meaning)
- Searches memories by meaning (not just keywords)
- Connects to AI assistants via MCP

## Prerequisites

- Python 3.10 or higher
- LM Studio running with embedding model loaded
- Basic Python knowledge

## Setup

1. **Create virtual environment:**
```bash
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
```

2. **Install dependencies:**
```bash
pip install -r requirements.txt
```

3. **Create .env file:**
```bash
cp .env.example .env
# Edit .env with your settings
```

4. **Run the server:**
```bash
python src/server.py
```

## Your Tasks

The template has TODO comments marking what you need to implement:

### 1. Storage (`src/storage/database.py`)
- [ ] `init_database()` - Set up SQLite database
- [ ] `store_memory()` - Save text and embedding
- [ ] `search_memories()` - Find similar memories

### 2. Embeddings (`src/storage/embeddings.py`)
- [ ] `generate_embedding()` - Call LM Studio API
- [ ] `cosine_similarity()` - Calculate similarity score

### 3. Tools (`src/tools/`)
- [ ] `store_memory_handler()` - Handle store requests
- [ ] `search_memory_handler()` - Handle search requests

## Testing

Use MCP Inspector to test your tools:
```bash
npx @modelcontextprotocol/inspector python src/server.py
```

## Project Structure

```
python-template/
├── src/
│   ├── server.py           # Main MCP server
│   ├── config.py           # Configuration
│   ├── storage/
│   │   ├── database.py     # SQLite operations
│   │   └── embeddings.py   # Embedding generation
│   └── tools/
│       ├── store_memory.py # Store tool handler
│       └── search_memory.py # Search tool handler
├── .env.example            # Example configuration
├── requirements.txt        # Python dependencies
└── README.md              # This file
```

## Need Help?

- Compare with TypeScript example: `../../examples/basic-typescript-example/`
- Check the docs: `../../docs/`
- See [Python MCP SDK docs](https://github.com/modelcontextprotocol/python-sdk)

Good luck! 🚀
