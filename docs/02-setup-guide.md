# Setup Your Environment 🛠️

This guide will walk you through installing everything you need to build your AI memory system.

**Time Required:** 30-60 minutes (mostly downloading models)

---

## 📋 What We're Installing

1. **NVM (Node Version Manager)** - Manage Node.js versions professionally
2. **Node.js** - JavaScript runtime (required for TypeScript, optional for Python)
3. **LM Studio** - Local AI model runner (no API costs!)
4. **AI Models** - Embedding model + LLM for testing
5. **Git** - Version control

---

## 1️⃣ Install Node.js via NVM (Recommended)

**NVM (Node Version Manager)** lets you install and switch between different Node.js versions. This is the professional way to manage Node.js and will save you headaches later!

### Why NVM?

- ✅ **Switch Node versions** easily between projects
- ✅ **No permission issues** - installs in your user folder
- ✅ **Industry standard** - what professionals use
- ✅ **Future-proof** - update Node anytime without reinstalling

### Install NVM

**Windows:**

1. Go to [github.com/coreybutler/nvm-windows/releases](https://github.com/coreybutler/nvm-windows/releases)
2. Download `nvm-setup.exe` from the latest release
3. Run the installer
4. Accept all defaults

**Mac / Linux:**

1. Open Terminal
2. Run the install script:
   ```bash
   curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash
   ```
3. Close and reopen your terminal

### Install Node.js with NVM

Once NVM is installed, install Node.js:

```bash
# Install the LTS version
nvm install --lts

# Use it as default
nvm use --lts
```

### Verify Installation

```bash
node --version
# Should show: v20.x.x or similar

npm --version
# Should show: 10.x.x or similar

nvm --version
# Should show: (Windows) 1.x.x or (Mac/Linux) 0.40.x
```

✅ If you see version numbers, you're good!

> 💡 **Tip:** In the future, you can easily update Node with: `nvm install --lts && nvm use --lts`

### Alternative: Direct Node.js Install

If you have trouble with NVM, you can install Node.js directly:

1. Go to [nodejs.org](https://nodejs.org/)
2. Download the **LTS version**
3. Run the installer and accept defaults

This works fine, but you'll miss out on easy version management.

---

## 2️⃣ Install LM Studio

LM Studio runs AI models locally on your computer - no internet required, no API costs, complete privacy!

### Download

1. Go to [lmstudio.ai](https://lmstudio.ai/)
2. Download for your operating system (Windows/Mac/Linux)
3. Install like any normal application

### First Launch

1. Open LM Studio
2. You'll see a clean interface with a search bar

---

## 3️⃣ Download AI Models

**You only need ONE model for the basic project:**

1. **Embedding Model** (Required) - Converts text to numbers (for semantic search)
2. **LLM** (Optional) - Only needed for advanced features or testing with LM Studio

> 💡 **For this assignment, just download the embedding model!** The LLM is only needed if you're doing advanced extensions with Neo4j or want to test with LM Studio instead of Claude/Copilot.

### Download Embedding Model (Required)

**Recommended: embeddinggemma-300M (Smaller, Newer)**

1. Search for: `embeddinggemma`
2. Look for: **`embeddinggemma-300m-qat-GGUF`**

**Size:** ~229 MB
**Dimensions:** 768 (default), can be truncated to 256/512 for speed
**Speed:** Very fast, great for low-end hardware
**Why:** Lightweight from Google, perfect for students with modest computers

> - **embeddinggemma-300M** = Smaller, faster, still great quality!

### Download LLM (Optional - Skip for Basic Project)

**Only download if you're doing advanced features or want to test with LM Studio!**

**Recommended: Qwen3-4B-Instruct-2507**

1. In LM Studio, search for: `qwen3-4b-2507`
2. Look for: **`lmstudio-community/Qwen3-4B-Instruct-2507-GGUF`**
3. Download: `Qwen3-4B-Instruct-2507-Q4_K_M.gguf` (good balance)

**Size:** ~2.5 GB
**Requirements:** Works great on laptops with 6GB+ GPU (tested on RTX 4050) or 8GB+ RAM
**Why:** Latest Qwen model, excellent tool-calling ability, fast response times

### Check Downloads

Click **💾 My Models** (left sidebar) to see your downloaded models.

**For the basic project, you should have:**
- ✅ One embedding model (~229 MB)

**Optional (only if doing advanced features):**
- One LLM (~2.5 GB)

---

## 4️⃣ Start LM Studio Server

Now let's make these models accessible via API!

### Load Your Models

1. In LM Studio, click **🔬 Developer** (left sidebar)
2. You'll see a section where you can load multiple models simultaneously
3. **Load the embedding model:**
   - Click "Select a model to load"
   - Choose your **embedding model** (`text-embedding-embeddinggemma-300m-qat`)
   - The model will show as "READY" when loaded

> 💡 **That's it!** You don't need to load the LLM for the basic project. You'll use Claude/Copilot as your AI assistant.

### Configure Server Settings

Before starting the server, configure settings (click the ⚙️ **Settings** button):
- **Enable CORS**: Toggle ON (allows your code to connect to the API)
- **Allow MCP**: Set to **"Remote"** if you want to use MCP features
- **Server Port**: Leave as `1234` (default)

> 💡 **Important:** You MUST enable CORS for your code to connect to LM Studio's API!

### Start the Server

Once your models are loaded and settings configured:
1. Look for the **"Status: Running"** toggle at the top
2. Make sure it's switched ON (green)
3. You should see: `Reachable at: http://127.0.0.1:1234`

### Test the Server

Open a new terminal and run:

```bash
curl http://localhost:1234/v1/models
```

You should see JSON output listing your loaded models. ✅

> 💡 **Tip:** You can load both models in one LM Studio instance! No need for multiple windows.

---

## 5️⃣ Code Editor

You'll need a code editor. Use whatever you're comfortable with! Popular choices:
- **VS Code** ([code.visualstudio.com](https://code.visualstudio.com/)) - Most popular
- **Cursor** - VS Code fork with AI features
- **Any other editor** - Whatever you prefer!

---

## 6️⃣ Install Git

Git helps you track changes and submit your project.

### Download & Install

1. Go to [git-scm.com](https://git-scm.com/)
2. Download for your OS
3. Install with defaults

### Verify Installation

```bash
git --version
# Should show: git version 2.x.x
```

### Configure Git (First Time Only)

```bash
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"
```

---

## 7️⃣ Create Your Project

Now let's set up your project folder!

### Clone the Starter Template

```bash
# Navigate to where you want your project
cd C:\personal\projects  # Windows
# or
cd ~/projects  # Mac/Linux

# Clone the starter repository
git clone https://github.com/Matz-Mudah/mcp-memory-starter.git
cd mcp-memory-starter

# Choose your language and navigate to template:
cd starter-templates/typescript-template  # For TypeScript
# OR
cd starter-templates/python-template      # For Python

# Install dependencies (TypeScript)
npm install
# OR
# Install dependencies (Python)
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
pip install -r requirements.txt
```

### Project Structure

**TypeScript template:**
```
typescript-template/
├── src/
│   ├── index.ts          # Main entry point
│   ├── tools/            # MCP tools (store, search)
│   └── storage/          # Database logic
├── package.json          # Dependencies
└── tsconfig.json         # TypeScript config
```

**Python template:**
```
python-template/
├── src/
│   ├── server.py         # Main entry point
│   ├── tools/            # MCP tools (store, search)
│   └── storage/          # Database logic
├── requirements.txt      # Dependencies
└── .env.example          # Configuration template
```

---

## 8️⃣ Test Your Setup

Let's make sure everything works!

### Test 1: Node.js

```bash
node --version
npm --version
```

✅ Both should show version numbers

### Test 2: LM Studio API

```bash
curl http://localhost:1234/v1/models
```

✅ Should return JSON with your model info

### Test 3: Generate Embeddings

```bash
curl http://localhost:1234/v1/embeddings \
  -H "Content-Type: application/json" \
  -d '{
    "input": "Hello world",
    "model": "text-embedding-embeddinggemma-300m-qat"
  }'
```

✅ Should return JSON with an array of numbers (the embedding!)

### Test 4: Build Your Project

**TypeScript:**
```bash
cd typescript-template
npm run build
```
✅ Should compile without errors

**Python:**
```bash
cd python-template
python -c "import mcp; print('✅ MCP SDK installed')"
```
✅ Should print success message

---

## 🎯 Configuration File

Create a `.env` file in your project root:

```bash
# LM Studio Embedding API
EMBEDDING_BASE_URL=http://localhost:1234/v1
EMBEDDING_MODEL=text-embedding-embeddinggemma-300m-qat

# Database
DB_PATH=./data/memories.db
```

> 💡 **Note:** `embeddinggemma-300M` produces 768-dimensional embeddings by default. The "300M" refers to the model size (300 million parameters), not the embedding dimension.

---

## 🚨 Troubleshooting

### Issue: "nvm: command not found" (Mac/Linux)

**Solution:** NVM install script didn't update your shell profile.
- Close and reopen your terminal
- Or manually add to `~/.bashrc` or `~/.zshrc`:
  ```bash
  export NVM_DIR="$HOME/.nvm"
  [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
  ```
- Run `source ~/.bashrc` (or `source ~/.zshrc`)

### Issue: "node: command not found"

**Solution:** Node.js didn't install correctly or NVM isn't active.
- Make sure you ran `nvm use --lts` after installing
- Restart your terminal
- Try `nvm list` to see installed versions
- On Windows, try restarting your computer
- If all else fails, use the direct Node.js installer

### Issue: "Cannot connect to localhost:1234"

**Solution:** LM Studio server isn't running or CORS is disabled.
- Open LM Studio
- Go to **Developer** tab
- Make sure the **"Status: Running"** toggle is ON (green)
- **Check that CORS is enabled** in Settings (⚙️ button)
- Check that nothing else is using port 1234

### Issue: "CORS error" or "Access-Control-Allow-Origin"

**Solution:** CORS is not enabled in LM Studio.
- In LM Studio **Developer** tab, click the **⚙️ Settings** button
- Toggle **"Enable CORS"** to ON (it should be green)
- The server will automatically apply the change

### Issue: "Model not found"

**Solution:** Model isn't loaded in Developer tab.
- In LM Studio, go to **Developer** tab
- Make sure you've loaded your embedding model (should show "READY")
- Check the model name matches what you're requesting
- Try reloading the model if needed

### Issue: npm install fails

**Solution:** Network or permissions issue.
- Try running as administrator (Windows)
- Try: `npm cache clean --force`
- Check your internet connection
- Delete `node_modules` folder and try again

### Issue: Embeddings return wrong dimensions

**Solution:** Check the model documentation for dimensions.
- Check which model you're actually running in LM Studio
- `embeddinggemma-300M` produces 768 dimensions by default
- Most embedding models use 768, 1024, or 1536 dimensions

---

## ✅ Setup Checklist

Before moving to the next step, verify:

- [ ] NVM installed (`nvm --version` works)
- [ ] Node.js installed via NVM (`node --version` works)
- [ ] LM Studio installed and opened
- [ ] Embedding model downloaded
- [ ] LM Studio server running on port 1234
- [ ] Can generate embeddings via curl
- [ ] Code editor ready (VS Code, Cursor, etc.)
- [ ] Git installed and configured
- [ ] Project template cloned
- [ ] `npm install` completed successfully
- [ ] `.env` file created with correct settings

**All checked?** You're ready to build! 🚀

---

## 🎮 Bonus: LM Studio as MCP Client

**NEW in LM Studio 0.3.29+**: You can use LM Studio itself to test your MCP tools!

### Why This is Awesome

Instead of needing Claude Desktop or another AI application, you can:
- Test your memory system with **local models** (like qwen3-4b)
- Stay **completely offline**
- Get **instant feedback** during development
- Use any model you want!

### Requirements

To use LM Studio as an MCP client:

1. **In LM Studio Developer tab settings:**
   - **Enable CORS**: Must be ON
   - **Allow MCP**: Must be set to **"Remote"** (not "Off")
   - This lets LM Studio connect to external MCP servers

2. **Your MCP server must:**
   - Have an **HTTP/SSE endpoint** (not just stdio)
   - Be accessible at a URL like `http://localhost:8080/sse`

### How It Works

LM Studio can connect to MCP servers and let local models use your tools through the `/v1/responses` endpoint:

```bash
# Your AI model can call your memory tools!
POST http://localhost:1235/v1/responses
{
  "model": "qwen3-4b-2507",
  "messages": [...],
  "tools": [
    {
      "type": "mcp",
      "server_label": "memory-system",
      "server_url": "http://localhost:8080/sse"
    }
  ]
}
```

### Setup Steps (Advanced)

**Note:** This is more advanced than the basic assignment. Here's the overview:

1. **Modify your MCP server** to expose an HTTP/SSE endpoint (in addition to stdio)
2. **Enable "Allow MCP" → Remote** in LM Studio Developer tab settings
3. **Load your LLM** in the Developer tab
4. **Use the `/v1/responses` endpoint** with `type: "mcp"` in tools array
5. **Watch your local model autonomously call your memory tools!** 🎯

**This is EXTRA CREDIT territory** - but imagine showing off a completely local AI system with memory!

We'll cover this in an advanced tutorial, but know that it's possible!

---

## 📊 Hardware Requirements

**Minimum (Will Work):**
- **CPU:** Any modern processor (Intel i3, AMD Ryzen 3, or equivalent)
- **RAM:** 8 GB
- **Storage:** 5 GB free space
- **GPU:** Not required (CPU-only is fine for small models)

**Recommended (Better Performance):**
- **CPU:** Intel i5/AMD Ryzen 5 or better
- **RAM:** 16 GB
- **Storage:** 10 GB free space (SSD preferred)
- **GPU:** Any NVIDIA GPU with 4GB+ VRAM (enables faster inference)

> 💡 **Tip:** The models we selected are specifically chosen to run well on modest hardware. Even a 5-year-old laptop should work!

---

## 🎓 Understanding Your Setup

### Why LM Studio?

- **Free & Open Source** - No API costs ever
- **Privacy** - Everything runs locally
- **Learning** - See exactly how models work
- **Portable** - Works offline, no internet needed

### Why These Specific Models?

**embeddinggemma-300M:**
- Dedicated embedding model from Google
- Very lightweight (229 MB)
- 768 dimensions (can be truncated to 256/512 for speed via Matryoshka)
- Perfect for student computers - tested on modest hardware
- Fast on CPU, even faster with GPU

**Qwen3-4B-Instruct-2507:**
- Small enough for student computers (2.5 GB)
- Excellent tool-calling ability (can use your MCP tools!)
- Fast response times - tested on RTX 4050 laptop
- Great at code and technical text
- Latest version with improved performance

---

**Next:** [Build Your First MCP Tool](03-first-mcp-tool.md) →

---

## 📚 Additional Resources

- [LM Studio Documentation](https://lmstudio.ai/docs)
- [Node.js Guides](https://nodejs.org/en/docs/guides/)
- [VS Code Tips](https://code.visualstudio.com/docs)
- [Understanding Embedding Models](https://huggingface.co/blog/getting-started-with-embeddings)

*Having issues? Ask your teacher or check the [Troubleshooting Guide](troubleshooting.md)!*
