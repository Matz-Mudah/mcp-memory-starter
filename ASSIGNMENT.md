# Assignment: Build Your Own AI Memory System 🧠

Build a memory system that allows AI assistants to remember and recall information across conversations - the same technology used by real AI companies today!

---

## 🎯 What You'll Build

A working AI memory system that:
- **Understands meaning** - Search by concept, not just keywords (semantic search)
- **Persists data** - Memories survive restarts
- **Integrates with AI** - Works with Claude, Copilot, or other MCP-compatible tools
- **Uses local models** - No expensive API calls (LM Studio)

---

## 📦 Core Requirements

Your project must include:

### ✅ Functionality
- [ ] **Store memory tool** - Saves information with embeddings
- [ ] **Search memory tool** - Finds relevant memories by meaning
- [ ] **Persistence** - Data survives server restart
- [ ] **MCP integration** - Connects to at least one AI platform

### 📝 Code Quality
- [ ] Clean, readable code
- [ ] Basic error handling
- [ ] Comments explaining key logic

### 📚 Documentation
- [ ] README with:
  - What your project does
  - Setup instructions (step-by-step)
  - How to use it
  - Example queries

### 🧪 Testing
- [ ] Demonstrates semantic search (finds by meaning, not keywords)
- [ ] Shows persistence (restart test)
- [ ] Works on fresh install

---

## 🚀 Getting Started

**Follow the learning path:**
1. [Understand Concepts](docs/01-concepts.md) - Embeddings and MCP
2. [Setup Environment](docs/02-setup-guide.md) - Install tools
3. [Build First Tool](docs/03-first-mcp-tool.md) - Create MCP server
4. [Add Storage](docs/04-memory-storage.md) - Implement persistence
5. [Semantic Search](docs/05-semantic-search.md) - Vector similarity
6. [Connect AI Platform](docs/choosing-ai-platform.md) - Choose your tool
7. [Finish Project](docs/06-final-project.md) - Polish and test

**Optional:** [Advanced Setup](docs/07-advanced-production.md) - Qdrant, Neo4j, Docker

---

## 💡 Make It Your Own

**Creative use case ideas:**
- **AI personality system** - Give your AI persistent memory and personality traits
- Recipe memory (search by ingredients)
- Study assistant (course notes by topic)
- Code snippets library (search by concept)
- Movie/book tracker (find by mood/theme)
- Gaming achievements log
- Workout tracker (exercises by muscle group)

**Optional extensions:**
- Metadata filtering (by date, category, etc.)
- Delete memory tool
- List all memories
- Export/import data
- Multiple collections
- Qdrant integration (see [Advanced Guide](docs/07-advanced-production.md))

---

## 🆘 Getting Help

**When stuck:**
1. Check the docs - Guides for common issues
2. Compare with [working example](examples/basic-typescript-example/)
3. Ask your teacher
4. Pair program with classmates

**Resources:**
- [MCP Documentation](https://modelcontextprotocol.io/)
- [Platform Setup Guides](docs/)
- [Troubleshooting](docs/02-setup-guide.md#troubleshooting) (in each guide)

---

## 📤 Submission

**What to submit:**
- GitHub repository link
- README with setup instructions
- Working code (tested on fresh install)

**Your teacher will specify:**
- Submission deadline
- Presentation format (demo, video, live presentation, etc.)
- Grading criteria
- Any additional requirements

---

## ✅ Before You Submit

- [ ] Code runs on a fresh clone
- [ ] README instructions are clear
- [ ] All core requirements work
- [ ] Tested semantic search (finds by meaning)
- [ ] Tested persistence (survives restart)
- [ ] No personal paths or secrets in code

---

## 🌟 Why This Matters

**You're learning real-world skills:**
- AI integration (embeddings, vector search)
- API design (building tools AI can use)
- Database management (SQLite, optionally Qdrant/Neo4j)
- System integration (MCP protocol)

**These technologies are used by:**
- AI startups building RAG systems
- Companies adding AI memory to products
- Open source AI projects
- Enterprise knowledge management systems

Add this to your portfolio - it shows you can build production-grade AI tools! 💼

---

**Ready to start?** Head to [Understanding the Concepts](docs/01-concepts.md) → 🚀
