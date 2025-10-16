# Connecting Your MCP Server to LM Studio

LM Studio 0.3.30+ supports MCP servers directly! This is a great option for students who want to use local models.

## Prerequisites

- ✅ Your MCP server is built (`npm run build`)
- ✅ Node.js is installed and in your PATH
- ✅ LM Studio 0.3.30 or higher installed
- ✅ A language model loaded in LM Studio

## Step-by-Step Setup

### 1. Open LM Studio

Launch LM Studio and load a language model that supports tool calling (e.g., qwen3-4b, llama-3.1, etc.)

### 2. Navigate to Program Tab

1. Click on the **Chat** tab (left sidebar)
2. On the right side menu, click **"Program"**
3. Click **"Install"** under "Integrations"

### 3. Edit mcp.json

LM Studio will open the MCP configuration file. Add your server:

```json
{
  "mcpServers": {
    "memory-system": {
      "command": "node",
      "args": [
        "C:\\personal\\mcp-memory-starter\\examples\\basic-typescript-example\\build\\index.js"
      ],
      "env": {}
    }
  }
}
```

**Important:**
- Use **absolute paths** to your `build/index.js` file
- On Windows: Use `\` (double backslash) or `/` (forward slash)
- Replace the path with your actual project location

### 4. Save and Restart

1. Save the `mcp.json` file
2. Restart LM Studio (or reload the MCP servers)
3. The server should appear in the "Integrations" section

### 5. Test It!

In the chat, try:
```
Store this memory: I'm learning about MCP servers in LM Studio
```

LM Studio should show it's calling the `store_memory` tool!

Then test search:
```
What am I learning about?
```

## Configuration Examples

### Basic Configuration
```json
{
  "mcpServers": {
    "my-memory-system": {
      "command": "node",
      "args": [
        "C:\\Users\\YourName\\projects\\mcp-memory-starter\\examples\\basic-typescript-example\\build\\index.js"
      ]
    }
  }
}
```

### With Environment Variables
```json
{
  "mcpServers": {
    "my-memory-system": {
      "command": "node",
      "args": [
        "C:\\path\\to\\build\\index.js"
      ],
      "env": {
        "DEBUG": "true"
      }
    }
  }
}
```

### Multiple Servers
```json
{
  "mcpServers": {
    "memory-system": {
      "command": "node",
      "args": ["C:\\path\to\\memory\\build\\index.js"]
    },
    "another-tool": {
      "command": "node",
      "args": ["C:\\path\to\\other\\build\\index.js"]
    }
  }
}
```

## Advantages of LM Studio

✅ **Runs locally** - No internet required
✅ **Free to use** - No API costs
✅ **Privacy** - Data stays on your machine
✅ **Visual interface** - Easy to see MCP tools in action
✅ **Model choice** - Use any compatible model

## Troubleshooting

### MCP Server Not Appearing

1. **Check the path is absolute**
   ```bash
   # Test manually first
   node C:\path\to\build\index.js
   ```

2. **Verify LM Studio version**
   - MCP support requires v0.3.30+
   - Update if needed: Help → Check for Updates

3. **Check Node.js is in PATH**
   ```bash
   node --version
   ```

### "Failed to Start MCP Server"

**Check LM Studio logs:**
- Windows: Look in LM Studio console output
- Check for error messages about the MCP server

**Common issues:**
- Path has typos or is relative (not absolute)
- Node.js not installed or not in PATH
- Build folder doesn't exist (run `npm run build`)
- TypeScript errors (check build output)

### Tools Not Working

**Verify in MCP Inspector first:**
```bash
npm run inspect
```

If tools work there but not in LM Studio:
- Make sure `.env` file exists in project root
- Check that embedding model is loaded in LM Studio
- Try restarting LM Studio completely

### Environment Variables Not Loading

Unlike Claude Desktop, LM Studio may have different behavior with `.env` files.

**Solution:**
Add them directly to `mcp.json`:
```json
{
  "mcpServers": {
    "memory-system": {
      "command": "node",
      "args": ["C:\\path\\to\\build\\index.js"],
      "env": {
        "EMBEDDING_BASE_URL": "http://localhost:1234/v1",
        "EMBEDDING_MODEL": "nomic-embed-text"
      }
    }
  }
}
```

## Recommended Models for MCP

Models that work well with tool calling:
- **qwen3-4b** - Fast, good for learning
- **llama-3.1-8b** - Better understanding
- **mistral-small** - Good balance

Load one of these in LM Studio for best results!

## Testing Workflow

1. **Build** your MCP server: `npm run build`
2. **Test** with MCP Inspector: `npm run inspect`
3. **Add** to LM Studio `mcp.json`
4. **Reload** MCP servers in LM Studio
5. **Chat** and watch the tools in action!

## Comparison: LM Studio vs Claude Desktop

| Feature | LM Studio | Claude Desktop |
|---------|-----------|----------------|
| **Cost** | Free | Requires Claude subscription |
| **Privacy** | Fully local | Sends to Anthropic |
| **Model Choice** | Any local model | Claude only |
| **Setup** | Edit `mcp.json` | Edit `claude_desktop_config.json` |
| **Best For** | Learning, experimenting | Production use |

## Next Steps

- [Test with MCP Inspector](../examples/basic-typescript-example/README.md#testing-the-system)
- [Build your own MCP server](../starter-templates/typescript-template/)
- [Learn more about LM Studio](https://lmstudio.ai/docs)

---

**Questions?** Check the [main README](../README.md) or ask your teacher!
