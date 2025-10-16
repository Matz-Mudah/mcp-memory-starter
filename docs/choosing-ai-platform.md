# Choosing Your AI Platform for MCP

Your MCP memory server can work with multiple AI platforms! Here's a guide to help you choose.

## Quick Comparison

| Feature | GitHub Copilot | Claude Code | Claude Desktop | LM Studio |
|---------|---------------|-------------|----------------|-----------|
| **Location** | VSCode built-in | VSCode extension | Standalone app | Standalone app |
| **AI Model** | GPT-5 | Claude Sonnet 4 | Claude Sonnet 4 | Local models |
| **Cost** | **Free for students!** | Subscription | Subscription | Free |
| **Privacy** | Cloud-based | Cloud-based | Cloud-based | Fully local |
| **Speed** | Fast (cloud) | Fast (cloud) | Fast (cloud) | Depends on hardware |
| **Setup Difficulty** | Easy | Easy | Easy | Moderate |
| **Config File** | `Code/User/mcp.json` | `~/.claude.json` | `claude_desktop_config.json` | `mcp.json` |
| **Student Access** | ✅ GitHub Student Pack | ⚠️ Requires paid plan | ⚠️ Requires paid plan | ✅ Free |
| **Best For** | **Students!** | Pro users in VSCode | General use | Privacy/offline |

## Platform Details

### 🟢 GitHub Copilot Chat (RECOMMENDED FOR STUDENTS!)

**Pros:**
- ✅ **FREE for students** via GitHub Student Pack!
- ✅ Built into VSCode (if you have Copilot)
- ✅ Code-aware AI assistant
- ✅ Fast cloud-based responses
- ✅ Easy MCP setup (just edit `mcp.json`)
- ✅ GPT-5 powered

**Cons:**
- ❌ Requires GitHub Copilot access
- ❌ Requires internet connection
- ❌ Data sent to OpenAI/GitHub

**Setup:**
Edit `Code/User/mcp.json` and add `"type": "stdio"`

**Config:** `C:\Users\YourName\AppData\Roaming\Code\User\mcp.json` (Windows)

**Get Student Access:** https://education.github.com/pack

**Guide:** [GitHub Copilot Setup](github-copilot-mcp-setup.md)

---

### 🔵 Claude Code (VSCode)

**Pros:**
- ✅ Integrated into your editor
- ✅ See your code while chatting
- ✅ Easy CLI setup (`claude mcp add`)
- ✅ Powerful Claude 4 Sonnet
- ✅ Context-aware (knows your files)

**Cons:**
- ❌ Requires Claude subscription
- ❌ Requires internet connection
- ❌ Data sent to Anthropic

**Setup:**
```bash
claude mcp add
# Follow the prompts!
```

**Config:** `C:\Users\YourName\.claude.json` (Windows) or `~/.claude.json` (Mac/Linux)

**Guide:** [Claude Code Setup](claude-code-mcp-setup.md)

---

### 🟣 Claude Desktop

**Pros:**
- ✅ Dedicated app for AI conversations
- ✅ Clean, focused interface
- ✅ Powerful Claude 4 Sonnet
- ✅ Great for non-coding tasks too

**Cons:**
- ❌ Requires Claude subscription
- ❌ Requires internet connection
- ❌ Manual config only

**Setup:**
Edit config file manually with server details.

**Config:** `%APPDATA%\Claude\claude_desktop_config.json` (Windows)

**Guide:** [Claude Desktop Setup](mcp-setup-guide.md)

---

### 🟢 LM Studio

**Pros:**
- ✅ **Completely free!**
- ✅ Fully offline/local
- ✅ Total privacy (no data sent anywhere)
- ✅ Choose any compatible model
- ✅ No API costs
- ✅ Visual MCP interface

**Cons:**
- ❌ Requires powerful hardware (GPU recommended)
- ❌ Models may be less capable than Claude
- ❌ Slower on older machines
- ❌ Initial setup more complex

**Setup:**
LM Studio → Chat → Program → Install → Edit mcp.json

**Config:** `mcp.json` inside LM Studio

**Guide:** [LM Studio Setup](lm-studio-mcp-setup.md)

---

## Decision Tree

### Choose **GitHub Copilot Chat** if:
- ✅ You're a student (free via Student Pack!)
- ✅ You already use GitHub Copilot
- ✅ You work in VSCode
- ✅ You want easy setup with good AI quality

### Choose **Claude Code** if:
- ✅ You have a Claude subscription
- ✅ You work primarily in VSCode
- ✅ You want the most powerful AI
- ✅ You need code context awareness

### Choose **Claude Desktop** if:
- ✅ You have a Claude subscription
- ✅ You want a standalone AI assistant
- ✅ You use it for more than just coding
- ✅ You prefer a dedicated app

### Choose **LM Studio** if:
- ✅ You want a free solution
- ✅ Privacy is important to you
- ✅ You want to learn with local models
- ✅ You have decent hardware (8GB+ RAM, GPU preferred)
- ✅ You don't have a Claude subscription

---

## For Students (VG1/VG2)

### Recommended Path:

**Week 1-2: Development**
- Use **MCP Inspector** for testing (`npm run inspect`)
- No AI needed yet, just verify tools work

**Week 3-6: Testing with AI**
- **Option A:** Use **LM Studio** (free, everyone can use)
- **Option B:** Use **Claude Code/Desktop** (if you have access)

**Week 7-10: Final Project**
- Use whichever platform you're comfortable with
- Show it working in your demo video

### Class Setup Ideas:

**If school has Claude licenses:**
- Use Claude Code in VSCode for development
- Best learning experience

**If no licenses available:**
- Use LM Studio (free!)
- Just needs decent computers
- Install models: qwen3-4b (4GB) or llama-3.1-8b (8GB)

**Mixed environment:**
- Teach both options
- Students choose based on their resources

---

## Can I Use Multiple?

**Yes!** You can configure the same MCP server for all three:

1. Add to `~/.claude.json` (Claude Code)
2. Add to `claude_desktop_config.json` (Claude Desktop)
3. Add to LM Studio's `mcp.json`

They all connect to the same database and work the same way!

---

## Performance Tips

### For Claude Code/Desktop:
- Fast responses (cloud AI)
- Best for complex queries
- Requires good internet

### For LM Studio:
- Speed depends on your hardware
- **4GB model** (qwen3-4b): Works on most laptops
- **8GB model** (llama-3.1-8b): Better quality, needs more RAM
- **GPU acceleration**: Much faster if you have NVIDIA GPU

---

## Next Steps

1. **Choose your platform** based on the criteria above
2. **Follow the setup guide** for your chosen platform:
   - [Claude Code Setup](claude-code-mcp-setup.md)
   - [Claude Desktop Setup](mcp-setup-guide.md)
   - [LM Studio Setup](lm-studio-mcp-setup.md)
3. **Test your MCP server** and start building!

---

**Questions?** Ask your teacher or check the [main README](../README.md)!
