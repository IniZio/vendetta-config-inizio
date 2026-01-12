# Agent-First Development

**Purpose**: Enable AI agents to work autonomously with minimal human intervention.

## Core Principles

### 1. Autonomous Execution

Agents should:
- Complete tasks without constant user supervision
- Make decisions about file selection and execution order
- Handle errors and retry automatically
- Seek clarification only when necessary

**Example**:
```typescript
// ✅ Good - Agent decides which files to edit
"Refactor the payment flow to support new delivery options"

// ❌ Bad - Agent requires manual guidance
"Edit payment.js and then edit delivery.js and then update types.ts"
```

### 2. Multi-File Awareness

Agents must understand codebase structure and make coordinated changes across multiple files.

**Guidelines**:
- Use semantic search to find relevant files
- Understand file dependencies and relationships
- Plan changes before executing
- Verify changes don't break existing code

**Example**:
```typescript
// ✅ Good - Agent handles multi-file changes
"Add support for new cake flavors and update pricing calculations"

// ❌ Bad - Single-file thinking
"Update the pricing calculation function"
```

### 3. Long-Horizon Planning

Tasks should be broken down into subtasks with clear milestones.

**Guidelines**:
- Identify main goal and subtasks
- Track progress through each subtask
- Handle failures without losing progress
- Support resume/interruption

**Example**:
```typescript
// Agent workflow for adding feature
1. Analyze requirements
2. Identify affected files (3-5 files expected)
3. Plan changes in logical order
4. Execute changes file by file
5. Run tests after each change
6. Handle errors and adjust plan
7. Finalize and document
```

### 4. Error Recovery

Agents must handle failures gracefully and recover automatically.

**Strategies**:
- **Retry Logic**: Transient failures retry 3 times with exponential backoff
- **Rollback**: Changes are atomic and can be undone
- **Fallback**: Alternative approaches when primary fails
- **Logging**: All errors logged with context for analysis

**Example**:
```typescript
// Agent error handling
async function executeWithRetry(operation: () => Promise<void>) {
  let attempts = 0;
  const maxAttempts = 3;

  while (attempts < maxAttempts) {
    try {
      await operation();
      return;
    } catch (error) {
      attempts++;
      if (attempts >= maxAttempts) {
        throw error;
      }
      await sleep(1000 * Math.pow(2, attempts)); // Exponential backoff
    }
  }
}
```

## Tool Orchestration

### Model Context Protocol (MCP) Integration

Agents should orchestrate multiple tools effectively:

**Best Practices**:
1. **Tool Discovery**: Automatically find available tools
2. **Tool Selection**: Choose correct tool for each subtask
3. **Parameter Construction**: Invoke tools with correct parameters
4. **Cross-Tool Coordination**: Pass outputs between tools
5. **State Management**: Track tool states and intermediate results

**Example**:
```typescript
// Agent using multiple tools
async function implementFeature(featureDescription: string) {
  // Tool 1: Search for relevant files
  const files = await mcp.codebase.search(featureDescription);

  // Tool 2: Analyze codebase structure
  const structure = await mcp.codebase.analyze(files);

  // Tool 3: Generate code changes
  const changes = await mcp.codegen.generate(featureDescription, structure);

  // Tool 4: Apply changes
  for (const change of changes) {
    await mcp.file.apply(change);
  }

  // Tool 5: Run tests
  const results = await mcp.test.run();
  return results;
}
```

## Architecture Patterns

### Multi-Phase Intelligence System

Inspired by Apex2-Terminal-Bench-Agent (#1 on Stanford Terminal Bench):

```
┌─────────────────────────────────────┐
│  Predictive Intelligence Layer        │  <- Plans strategy
└──────────────┬──────────────────┘
               │
┌──────────────▼──────────────────┐
│   Deep Strategy Generator          │  <- Detailed plans
└──────────────┬──────────────────┘
               │
┌──────────────▼──────────────────┐
│     Strategy Synthesizer          │  <- Combines results
└─────────────────────────────────┘
```

**Implementation**:
1. **Predictive Layer**: Predicts what tools and files are needed
2. **Strategy Generator**: Creates detailed execution plans
3. **Synthesizer**: Combines intermediate results into final output

### Parallel Execution

Run multiple independent agents simultaneously:

```typescript
// Parallel agent workflow
async function parallelFeatureImplementation() {
  const [backend, frontend, tests] = await Promise.all([
    runAgent('backend', implementBackendAPI),
    runAgent('frontend', implementUIComponents),
    runAgent('testing', writeIntegrationTests)
  ]);

  return { backend, frontend, tests };
}
```

## Coding Standards

### Type Safety

Always use TypeScript with strict mode:

```typescript
// ✅ Good
interface CakeConfig {
  flavor: FlavorType;
  size: SizeType;
  decoration: DecorationType;
}

function calculatePrice(config: CakeConfig): number {
  // Implementation
}

// ❌ Bad
function calculatePrice(config: any): number {
  // Implementation
}
```

### Error Handling

Never suppress errors:

```typescript
// ✅ Good - Proper error handling
try {
  await saveOrder(order);
} catch (error) {
  if (error instanceof DatabaseError) {
    throw new OrderSaveError('Failed to save order', { cause: error });
  }
  throw error;
}

// ❌ Bad - Suppressing errors
try {
  await saveOrder(order);
} catch (e) {
  // Silently fail
}
```

### Test Coverage

Every feature must have tests:

```typescript
// Feature implementation
export function calculateCakePrice(config: CakeConfig): number {
  // Implementation
}

// Test suite
describe('calculateCakePrice', () => {
  it('applies flavor multiplier', () => {
    // Test
  });

  it('applies size base price', () => {
    // Test
  });

  it('handles custom decorations', () => {
    // Test
  });
});
```

## Performance Optimization

### Context Management

- Use semantic search for relevant context
- Maintain 200K+ token context window (like Claude 4.1)
- Cache frequently accessed files
- Incrementally update context

### Execution Speed

Target: Complete 90% of agentic turns within 30 seconds (Cursor Composer benchmark):

```typescript
// Fast execution patterns
async function quickEdit(file: string, change: CodeChange) {
  // Direct edit without full re-read
  await mcp.file.edit(file, change);

  // Run only affected tests
  await mcp.test.run(file);

  // Skip full rebuild unless needed
  // if (!change.affectsBuild) return;
  await mcp.build.incremental();
}
```

## Evaluation Metrics

### Success Criteria

An agent is effective if it:
- Completes tasks autonomously (minimal user intervention)
- Achieves 90%+ on SWE-Bench Verified
- Handles multi-file refactors correctly
- Recovers from errors automatically
- Uses tools correctly (MCP Atlas score > 60%)

### Benchmark Targets

- **SWE-Bench Verified**: 77.2% (Claude 4 Sonnet)
- **MCP Atlas**: 62.3% (top performer)
- **Terminal Bench**: 64.5% (Apex2 with Claude Sonnet 4.5)

## Anti-Patterns

### ❌ What to Avoid

1. **Single-File Focus**: Only editing one file when multiple are needed
2. **Manual Guidance**: Asking user which files to edit
3. **No Error Recovery**: Failing on first error without retry
4. **No Planning**: Making changes without understanding impact
5. **Poor Tool Usage**: Calling wrong tools with wrong parameters
6. **Context Dumping**: Loading entire codebase instead of relevant files
7. **Suppressed Errors**: Using `@ts-ignore` or empty catch blocks

### ✅ What to Do

1. **Multi-File Awareness**: Understand codebase structure
2. **Autonomous Execution**: Make decisions without asking
3. **Error Recovery**: Retry and handle failures gracefully
4. **Long-Horizon Planning**: Break complex tasks into subtasks
5. **Tool Orchestration**: Use tools correctly and combine results
6. **Type Safety**: Use TypeScript with strict mode
7. **Test Coverage**: Write tests for all changes

## References

- [Cursor Composer](https://learn-cursor.com/en/docs/composer)
- [MCP Atlas](https://scale.com/leaderboard/mcp_atlas)
- [Apex2 Terminal Bench](https://github.com/heartyguy/Apex2-Terminal-Bench-Agent)
- [SWE-Bench](https://scale.com/blog/swe-bench-pro)
- [Agentic Coding Leaderboards](/docs/AGENTIC_CODING_LEADERBOARDS.md)

---

**Version**: 1.0
**Last Updated**: 2025-01-12
