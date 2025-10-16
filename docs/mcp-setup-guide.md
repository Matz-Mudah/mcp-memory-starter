# Connecting Your MCP Server to Claude Desktop

Once you've built your memory system, you'll want to connect it to Claude Desktop so you can use it in conversations!

## Prerequisites

- ✅ Your MCP server is built (`npm run build`)
- ✅ Node.js is installed and in your PATH
- ✅ Claude Desktop is installed

## Step-by-Step Setup

### 1. Locate Your Config File

**Windows:**
```
%APPDATA%\Claude\claude_desktop_config.json
```

**Mac:**
```
~/Library/Application Support/Claude/claude_desktop_config.json
```

**Linux:**
```
~/.config/Claude/claude_desktop_config.json
```

### 2. Edit the Config File

Open the file in a text editor. If it doesn't exist, create it with this content:

```json
{
  "mcpServers": {}
}
```

### 3. Add Your MCP Server

Add your server to the `mcpServers` object:

```json
{
  "mcpServers": {
    "my-memory-system": {
      "command": "node",
      "args": [
        "/absolute/path/to/your/project/build/index.js"
      ],
      "env": {}
    }
  }
}
```

**Important Path Notes:**
- ✅ Use **absolute paths** (e.g., `C:\Users\...` not `.\build\...`)
- ✅ On Windows: Use `\\` (double backslash) or `/` (forward slash)
- ✅ Replace `/absolute/path/to/your/project/` with your actual path

**Example for Windows:**
```json
{
  "mcpServers": {
    "my-memory-system": {
      "command": "node",
      "args": [
        "C:\\Users\\YourName\\Documents\\mcp-memory-starter\\examples\\basic-typescript-example\\build\\index.js"
      ],
      "env": {}
    }
  }
}
```

**Example for Mac/Linux:**
```json
{
  "mcpServers": {
    "my-memory-system": {
      "command": "node",
      "args": [
        "/Users/YourName/Documents/mcp-memory-starter/examples/basic-typescript-example/build/index.js"
      ],
      "env": {}
    }
  }
}
```

### 4. Adding Environment Variables (Optional)

If your server needs specific environment variables:

```json
{
  "mcpServers": {
    "my-memory-system": {
      "command": "node",
      "args": [
        "C:\\path\\to\\build\\index.js"
      ],
      "env": {
        "EMBEDDING_BASE_URL": "http://localhost:1234/v1",
        "DEBUG": "true"
      }
    }
  }
}
```

> **Note:** Most settings should be in your `.env` file, not here!

### 5. Restart Claude Desktop

**Important:** You must **completely close and reopen** Claude Desktop for changes to take effect.

- Windows: Right-click tray icon → Quit
- Mac: Cmd+Q to quit fully
- Then reopen Claude Desktop

### 6. Verify It's Working

In a new conversation with Claude, try:

```
Store this memory: I love TypeScript programming
```

You should see Claude use the `store_memory` tool!

Then test search:
```
What programming language do I like?
```

Claude should use `search_memory` and find your memory!

## Troubleshooting

### Tools Don't Appear

**Check the logs:**
- Windows: `%APPDATA%\Claude\logs\mcp*.log`
- Mac: `~/Library/Logs/Claude/mcp*.log`

**Common issues:**
1. Path is not absolute
2. Path has typos
3. Build folder doesn't exist (run `npm run build`)
4. Node.js not in PATH

### "Failed to Start MCP Server"

**Test manually:**
```bash
node C:\\path\\to\\your\\build\\index.js
```

If this fails, your server has issues. Check:
- TypeScript compiled successfully
- All dependencies installed (`npm install`)
- No syntax errors in code

### Server Starts But Tools Don't Work

**Check MCP Inspector first:**
```bash
npm run inspect
```

If tools work in Inspector but not Claude:
- Verify `.env` file exists in project root
- Check LM Studio is running (if using embeddings)
- Look at Claude Desktop logs for errors

### Multiple MCP Servers

You can run multiple servers! Just add more entries:

```json
{
  "mcpServers": {
    "my-memory-system": {
      "command": "node",
      "args": ["C:\\path\\to\\memory\\build\\index.js"]
    },
    "another-server": {
      "command": "node",
      "args": ["C:\\path\to\\other\build\\index.js"]
    }
  }
}
```

## Security Note

MCP servers run on your local machine with the same permissions as Claude Desktop. Only add servers you trust!

## Testing Tips

1. **Start simple** - Test with MCP Inspector first
2. **Check logs** - They tell you exactly what's wrong
3. **Use absolute paths** - Relative paths cause 90% of issues
4. **Restart fully** - Claude Desktop caches configs

## Next Steps

- [Test with MCP Inspector](../examples/basic-typescript-example/README.md#testing-the-system)
- [Build your own MCP server](../starter-templates/typescript-template/)
- [Learn more about MCP](https://modelcontextprotocol.io/)

---

**Questions?** Check the [main README](../README.md) or ask your teacher!
