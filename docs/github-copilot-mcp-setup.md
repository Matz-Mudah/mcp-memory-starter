# Connecting Your MCP Server to GitHub Copilot Chat

GitHub Copilot Chat (in VSCode) supports MCP servers! This is likely the most accessible option for students since many schools provide GitHub Copilot access.

## Prerequisites

- ✅ Your MCP server is built (`npm run build`)
- ✅ Node.js is installed and in your PATH
- ✅ VSCode installed
- ✅ GitHub Copilot Chat extension enabled
- ✅ GitHub account with Copilot access (student/teacher license)

## Configuration

### Step 1: Locate Your MCP Config File

**Windows:**
```
C:\Users\YourName\AppData\Roaming\Code\User\mcp.json
```

**Mac:**
```
~/Library/Application Support/Code/User/mcp.json
```

**Linux:**
```
~/.config/Code/User/mcp.json
```

> **Note:** If using VSCode Insiders, replace `Code` with `Code - Insiders`

### Step 2: Edit mcp.json

Open the file and add your MCP server:

```json
{
  "mcpServers": {
    "mcp-memory-starter": {
      "command": "node",
      "args": [
        "/path/to/mcp-memory-starter/examples/basic-typescript-example/build/index.js"
      ],
      "type": "stdio"
    }
  }
}
```

**Important:**
- ✅ Use **absolute paths** to your `build/index.js`
- ✅ Use forward slashes `/` or double backslashes `\`
- ✅ Add `"type": "stdio"` for Copilot compatibility
- ✅ Replace the path with your actual project location

### Step 3: Enable MCP in Copilot Chat

1. Open VSCode Settings (`Ctrl+,` or `Cmd+,`)
2. Search for: `github.copilot.chat.mcp`
3. Enable: **GitHub Copilot > Chat > MCP: Enabled**
4. Reload VSCode

### Step 4: Access MCP Servers in Chat

1. Open GitHub Copilot Chat (sidebar icon or `Ctrl+Alt+I`)
2. Click the **settings icon** in the chat input
3. Click **"MCP Servers"**
4. Your server should appear in the list!

## Configuration Examples

### Basic Configuration
```json
{
  "mcpServers": {
    "memory-system": {
      "command": "node",
      "args": [
        "C:/Users/Student/projects/mcp-memory-starter/examples/basic-typescript-example/build/index.js"
      ],
      "type": "stdio"
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
      },
      "type": "stdio"
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
      "args": ["C:/path/to/memory/build/index.js"],
      "type": "stdio"
    },
    "another-tool": {
      "command": "node",
      "args": ["C:/path/to/other/build/index.js"],
      "type": "stdio"
    }
  }
}
```

## Testing Your Setup

### 1. Open Copilot Chat

Click the Copilot Chat icon in the sidebar or use `Ctrl+Alt+I`

### 2. Access MCP Servers

1. Click the settings/gear icon in the chat input
2. Select "MCP Servers"
3. Your server should be listed

### 3. Test the Tools

In Copilot Chat, try:

```
Store this memory: I'm learning MCP with GitHub Copilot
```

Copilot should recognize and use your `store_memory` tool!

Then test search:
```
What am I learning?
```

## Troubleshooting

### MCP Server Not Appearing

1. **Check the config file location:**
   ```bash
   # Windows (PowerShell)
   cat "$env:APPDATA\Code\User\mcp.json"
   
   # Mac/Linux
   cat ~/Library/Application\ Support/Code/User/mcp.json
   ```

2. **Verify VSCode Insiders vs Stable:**
   - VSCode Stable: `Code\User\mcp.json`
   - VSCode Insiders: `Code - Insiders\User\mcp.json`

3. **Check JSON syntax:**
   - No trailing commas
   - Proper quotes
   - Valid structure

### "Failed to Start MCP Server"

**Check the build exists:**
```bash
# Verify the file exists
ls "C:\path\to\build\index.js"

# Test manually
node "C:\path\to\build\index.js"
```

**Common issues:**
- Path is relative (must be absolute)
- Build folder doesn't exist (run `npm run build`)
- Node.js not in PATH
- Missing `"type": "stdio"`

### Tools Not Working

1. **Verify MCP is enabled:**
   - Settings → `github.copilot.chat.mcp.enabled` → ✅

2. **Reload VSCode:**
   - Command Palette (`Ctrl+Shift+P`)
   - Type: "Reload Window"

3. **Check VSCode Output:**
   - View → Output
   - Select "GitHub Copilot" from dropdown
   - Look for MCP-related errors

### "type": "stdio" Missing

GitHub Copilot requires `"type": "stdio"` in the configuration:

```json
{
  "mcpServers": {
    "your-server": {
      "command": "node",
      "args": ["path/to/index.js"],
      "type": "stdio"  // ← Required for Copilot!
    }
  }
}
```

## Advantages of GitHub Copilot

✅ **Widely Available** - Many students/schools have access
✅ **Integrated in VSCode** - No separate app needed
✅ **Code-Aware** - Understands your codebase
✅ **Free for Students** - Via GitHub Student Pack
✅ **MCP Support** - Full tool calling capabilities

## Copilot vs Other Platforms

| Feature | GitHub Copilot | Claude Code | LM Studio |
|---------|---------------|-------------|-----------|
| **Config File** | `Code/User/mcp.json` | `~/.claude.json` | LM Studio's `mcp.json` |
| **Type Field** | Required (`"stdio"`) | Not needed | Not needed |
| **Student Access** | ✅ Free via Student Pack | ⚠️ Requires subscription | ✅ Free |
| **Setup Location** | VSCode settings | Separate config | In-app setup |
| **Best For** | Students with Copilot | Pro users | Privacy/offline |

## GitHub Student Pack

Students can get **free GitHub Copilot** access:

1. Go to: https://education.github.com/pack
2. Verify student status (school email or ID)
3. Get access to Copilot + other tools
4. Enable in VSCode

## Path Format Examples

**Windows (PowerShell):**
```json
{
  "mcpServers": {
    "memory": {
      "command": "node",
      "args": [
        "C:/Users/Student/projects/mcp-memory-starter/examples/basic-typescript-example/build/index.js"
      ],
      "type": "stdio"
    }
  }
}
```

**Mac:**
```json
{
  "mcpServers": {
    "memory": {
      "command": "node",
      "args": [
        "/Users/student/projects/mcp-memory-starter/examples/basic-typescript-example/build/index.js"
      ],
      "type": "stdio"
    }
  }
}
```

**Linux:**
```json
{
  "mcpServers": {
    "memory": {
      "command": "node",
      "args": [
        "/home/student/projects/mcp-memory-starter/examples/basic-typescript-example/build/index.js"
      ],
      "type": "stdio"
    }
  }
}
```

## Finding the Config File

### Windows (PowerShell):
```powershell
# Open config file directly
code "$env:APPDATA\Code\User\mcp.json"

# Or create if it doesn't exist
New-Item -Path "$env:APPDATA\Code\User\mcp.json" -ItemType File -Force
```

### Mac/Linux (Terminal):
```bash
# Open config file
code ~/Library/Application\ Support/Code/User/mcp.json

# Or create if it doesn't exist
mkdir -p ~/Library/Application\ Support/Code/User/
touch ~/Library/Application\ Support/Code/User/mcp.json
```

## Next Steps

- [Test with MCP Inspector first](../examples/basic-typescript-example/README.md#testing-the-system)
- [Compare all AI platforms](choosing-ai-platform.md)
- [Get GitHub Student Pack](https://education.github.com/pack)
- [Build your own MCP server](../starter-templates/typescript-template/)

---

**Questions?** Check the [main README](../README.md) or ask your teacher!
