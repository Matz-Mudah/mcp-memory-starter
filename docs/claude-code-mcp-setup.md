# Connecting Your MCP Server to Claude Code (VSCode)

Claude Code is the VSCode extension that brings Claude AI directly into your editor. It has built-in MCP support!

## Prerequisites

- ✅ Your MCP server is built (`npm run build`)
- ✅ Node.js is installed and in your PATH
- ✅ VSCode with Claude Code extension installed
- ✅ Claude subscription (Pro or Team)

## Setup Methods

### Method 1: Using CLI (Easiest)

Claude Code provides a CLI command to add MCP servers:

```bash
claude mcp add
```

This will:
1. Prompt you for the server name (e.g., "memory-system")
2. Ask for the command (`node`)
3. Ask for arguments (path to your `build/index.js`)
4. Automatically update your `.claude.json`

### Method 2: Manual Configuration

#### Step 1: Locate Your Config File

**Windows:**
```
C:\Users\YourName\.claude.json
```

**Mac/Linux:**
```
~/.claude.json
```

#### Step 2: Edit the Config

Open `.claude.json` and add your MCP server:

```json
{
  "mcpServers": {
    "mcp-memory-starter": {
      "command": "node",
      "args": [
        "/path/to/mcp-memory-starter/examples/basic-typescript-example/build/index.js"
      ],
      "env": {}
    }
  }
}
```

**Important Path Notes:**
- ✅ Use **absolute paths** (e.g., `C:/Users/...` not `./build/...`)
- ✅ On Windows: Use `/` (forward slash) - Claude Code handles this better
- ✅ Replace the path with your actual project location

**Example Configuration:**
```json
{
  "mcpServers": {
    "mcp-memory-starter": {
      "command": "node",
      "args": [
        "/path/to/projects/mcp-memory-starter/examples/basic-typescript-example/build/index.js"
      ],
      "env": {}
    }
  }
}
```

#### Step 3: Reload Claude Code

1. Open VSCode Command Palette (`Ctrl+Shift+P` or `Cmd+Shift+P`)
2. Type: `Claude: Reload MCP Servers`
3. Or restart VSCode

### Method 3: Using VSCode Settings

You can also configure MCP servers in VSCode settings:

1. `Ctrl+,` (or `Cmd+,`) to open Settings
2. Search for "Claude MCP"
3. Edit settings.json directly

## Testing Your Setup

### 1. Check MCP Status

In VSCode:
- Look for the Claude icon in the sidebar
- Click to open Claude Code
- Check the MCP servers list (should show your server)

### 2. Test in Chat

Start a conversation with Claude Code and try:

```
Store this memory: I'm learning MCP in VSCode
```

You should see Claude Code use the `store_memory` tool!

Then test search:
```
What am I learning?
```

### 3. View Tool Calls

Claude Code shows:
- 🔧 When tools are called
- 📊 Tool execution results
- ⚠️ Any errors

## Configuration Examples

### Basic Setup
```json
{
  "mcpServers": {
    "memory-system": {
      "command": "node",
      "args": [
        "C:/Users/YourName/projects/mcp-memory-starter/examples/basic-typescript-example/build/index.js"
      ]
    }
  }
}
```

### With Environment Variables
```json
{
  "mcpServers": {
    "memory-system": {
      "command": "node",
      "args": [
        "C:/path/to/build/index.js"
      ],
      "env": {
        "DEBUG": "true",
        "EMBEDDING_BASE_URL": "http://localhost:1234/v1"
      }
    }
  }
}
```

### Multiple MCP Servers
```json
{
  "mcpServers": {
    "memory-system": {
      "command": "node",
      "args": ["C:/path/to/memory/build/index.js"]
    },
    "another-tool": {
      "command": "node",
      "args": ["C:/path/to/other/build/index.js"]
    }
  }
}
```

## Troubleshooting

### "MCP Server Not Found"

1. **Verify the path exists:**
   ```bash
   # Windows
   dir "C:\path\to\build\index.js"
   
   # Mac/Linux
   ls -la /path/to/build/index.js
   ```

2. **Check it runs manually:**
   ```bash
   node C:/path/to/build/index.js
   ```

3. **Use forward slashes on Windows:**
   ```json
   "args": ["/path/to/project/build/index.js"]
   ```
   NOT:
   ```json
   "args": ["C:\Users\YourName\project\build\index.js"]
   ```

### "Failed to Start MCP Server"

**Check Claude Code Output:**
1. View → Output
2. Select "Claude" from the dropdown
3. Look for MCP startup errors

**Common issues:**
- Node.js not in PATH
- Build folder doesn't exist (run `npm run build`)
- Syntax errors in code
- Permission issues

### Tools Not Appearing

1. **Reload MCP servers:**
   - Command Palette → `Claude: Reload MCP Servers`

2. **Check `.claude.json` syntax:**
   - Valid JSON (no trailing commas)
   - Proper quotes
   - Correct brackets

3. **Verify server is running:**
   - Check Claude Code output for errors
   - Test with MCP Inspector first: `npm run inspect`

### Environment Variables Not Working

Claude Code may handle `.env` files differently:

**Solution:** Add them to `.claude.json`:
```json
{
  "mcpServers": {
    "memory-system": {
      "command": "node",
      "args": ["C:/path/to/build/index.js"],
      "env": {
        "EMBEDDING_BASE_URL": "http://localhost:1234/v1",
        "EMBEDDING_MODEL": "nomic-embed-text",
        "DB_PATH": "./data/memories.db"
      }
    }
  }
}
```

## Advantages of Claude Code

✅ **Integrated** - Claude AI directly in VSCode
✅ **Context-aware** - Can see your open files
✅ **Powerful AI** - Uses Claude Sonnet 4
✅ **MCP support** - Use custom tools
✅ **Code-friendly** - Perfect for development

## CLI Commands Reference

```bash
# Add a new MCP server
claude mcp add

# List configured MCP servers
claude mcp list

# Remove an MCP server
claude mcp remove <server-name>

# Test MCP server connection
claude mcp test <server-name>
```

## Best Practices

1. **Use absolute paths** - Always use full paths to avoid issues
2. **Test first** - Use MCP Inspector before adding to Claude Code
3. **Check logs** - View Output panel for debugging
4. **Reload often** - Reload MCP servers after config changes
5. **Environment vars** - Add them to config, not just `.env`

## Comparison: Claude Code vs Claude Desktop vs LM Studio

| Feature | Claude Code | Claude Desktop | LM Studio |
|---------|-------------|----------------|-----------|
| **Location** | VSCode extension | Standalone app | Standalone app |
| **Config File** | `~/.claude.json` | `~/Library/.../claude_desktop_config.json` | `mcp.json` in LM Studio |
| **Path Format** | `/` (forward slash) | `\` or `/` | `\` or `/` |
| **CLI Support** | ✅ `claude mcp add` | ❌ Manual only | ❌ Manual only |
| **Best For** | Development | General use | Local/offline |

## Next Steps

- [Test with MCP Inspector](../examples/basic-typescript-example/README.md#testing-the-system)
- [Read Claude Code docs](https://docs.claude.com/en/docs/claude-code/mcp)
- [Build your own MCP server](../starter-templates/typescript-template/)

---

**Questions?** Check the [main README](../README.md) or [Claude Code MCP docs](https://docs.claude.com/en/docs/claude-code/mcp)!
