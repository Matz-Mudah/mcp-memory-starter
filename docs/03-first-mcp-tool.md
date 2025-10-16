# Building Your First MCP Tool

This guide will walk you through creating your first MCP (Model Context Protocol) tool from scratch.

## What You'll Learn

- How MCP servers work
- Creating tool definitions
- Handling tool requests
- Testing with MCP Inspector

## Quick Start: Use the Working Example!

The fastest way to understand MCP tools is to look at the working example:

👉 **[See the Complete Example](../examples/basic-typescript-example/)**

The example includes:
- ✅ Full MCP server setup
- ✅ Two working tools (`store_memory` and `search_memory`)
- ✅ Detailed code comments
- ✅ Complete implementation

## Understanding the Template

The starter template at `starter-templates/typescript-template/` has TODOs for you to implement.

### Key Files to Understand

**1. [src/index.ts](../starter-templates/typescript-template/src/index.ts)**
- Main MCP server entry point
- Registers tools with the MCP SDK
- Handles incoming requests

**2. [src/tools/store-memory.ts](../starter-templates/typescript-template/src/tools/store-memory.ts)**
- Tool definition (what parameters it accepts)
- Handler function (what it does when called)

**3. [src/tools/search-memory.ts](../starter-templates/typescript-template/src/tools/search-memory.ts)**
- Similar structure to store-memory
- Shows how to return formatted results

## MCP Tool Anatomy

Every MCP tool has two parts:

### 1. Tool Definition
Tells the AI what the tool does and what parameters it needs:

```typescript
export const myTool = {
  name: 'my_tool',
  description: 'What this tool does',
  inputSchema: {
    type: 'object',
    properties: {
      param1: { 
        type: 'string',
        description: 'What this parameter is for'
      }
    },
    required: ['param1']
  }
}
```

### 2. Handler Function
Does the actual work when the tool is called:

```typescript
export async function handleMyTool(
  args: { param1: string },
  config: Config
): Promise<string> {
  // 1. Validate input
  if (!args.param1) {
    throw new Error('param1 is required');
  }

  // 2. Do something
  const result = doSomething(args.param1);

  // 3. Return result
  return `Success: ${result}`;
}
```

## Step-by-Step Guide

### 1. Start with the Template

```bash
cd starter-templates/typescript-template
npm install
npm run build
```

### 2. Look at the Working Example

Open `examples/basic-typescript-example/src/tools/store-memory.ts` and compare it to the template version.

**Notice:**
- Template has TODOs
- Example has working code
- Both have the same structure!

### 3. Implement Following the TODOs

The template guides you with hints:
```typescript
// TODO: Validate input
//   if (!args.text || args.text.trim().length === 0) {
//     throw new Error('Memory text cannot be empty');
//   }
```

### 4. Test Your Tool

```bash
npm run inspect
```

This opens MCP Inspector where you can test your tools!

## Testing Your First Tool

### Using MCP Inspector

1. Run: `npm run inspect`
2. Find your tool in the list
3. Enter test parameters:
   ```json
   {
     "text": "Test memory"
   }
   ```
4. Click "Call Tool"
5. See the result!

### Common Issues

**"Tool not found"**
- Check tool is registered in `src/index.ts`
- Make sure you rebuilt (`npm run build`)

**"Handler error"**
- Look at the error message
- Check your implementation matches the hints
- Compare with the working example

## Next Steps

1. ✅ Get your first tool working
2. 📚 **[Add Memory Storage](04-memory-storage.md)** - Make it persistent
3. 🔍 **[Implement Search](05-semantic-search.md)** - Add semantic search

## Additional Resources

- [MCP Documentation](https://modelcontextprotocol.io/)
- [Working Example Code](../examples/basic-typescript-example/src/)
- [Platform Setup Guides](choosing-ai-platform.md)

---

**Remember:** The goal isn't to write perfect code, but to understand how MCP tools work. Use the example as your guide! 🚀
