# Error Recovery and Resilience

**Purpose**: Ensure agents recover from errors gracefully without human intervention.

## Core Principles

### 1. Fail-Fast, Fail-Loudly

Catch errors early and report clearly:

```typescript
// ✅ Good - Clear error messages
async function saveOrder(order: Order): Promise<void> {
  if (!order.items || order.items.length === 0) {
    throw new ValidationError(
      'Order must contain at least one item',
      { field: 'items', value: order.items }
    );
  }

  await database.save(order);
}

// ❌ Bad - Silent failures
async function saveOrder(order: Order): Promise<void> {
  if (!order.items || order.items.length === 0) {
    return; // Silent failure
  }
  await database.save(order);
}
```

### 2. Retry with Backoff

Transient failures should be retried with exponential backoff:

```typescript
async function executeWithRetry<T>(
  operation: () => Promise<T>,
  options: {
    maxAttempts?: number;
    baseDelay?: number;
    backoffMultiplier?: number;
  } = {}
): Promise<T> {
  const {
    maxAttempts = 3,
    baseDelay = 1000,
    backoffMultiplier = 2
  } = options;

  let attempt = 0;

  while (attempt < maxAttempts) {
    try {
      return await operation();
    } catch (error) {
      attempt++;

      if (attempt >= maxAttempts) {
        throw new RetryExhaustedError(
          `Operation failed after ${maxAttempts} attempts`,
          { cause: error, attempts: maxAttempts }
        );
      }

      if (!isTransientError(error)) {
        throw error; // Don't retry non-transient errors
      }

      const delay = baseDelay * Math.pow(backoffMultiplier, attempt - 1);
      await sleep(delay);
    }
  }

  throw new Error('Unreachable');
}
```

### 3. Circuit Breaker

Stop calling failing tools after N consecutive failures:

```typescript
class CircuitBreaker {
  private failures = 0;
  private lastFailureTime = 0;
  private cooldownPeriod = 60000; // 1 minute

  async execute<T>(
    tool: MCPTool,
    params: any
  ): Promise<T> {
    const now = Date.now();

    if (this.failures >= 3 && (now - this.lastFailureTime) < this.cooldownPeriod) {
      throw new CircuitOpenError('Tool circuit is open');
    }

    try {
      const result = await tool.execute(params);
      this.failures = 0; // Reset on success
      return result;
    } catch (error) {
      this.failures++;
      this.lastFailureTime = now;
      throw error;
    }
  }
}
```

### 4. Fallback Mechanisms

Provide alternative approaches when primary fails:

```typescript
async function generateWithFallback(
  description: string
): Promise<string> {
  try {
    // Primary: Use specialized AI model
    return await ai.primary.generate(description);
  } catch (primaryError) {
    log.warn('Primary model failed, trying fallback', { error: primaryError });

    try {
      // Fallback 1: Use general-purpose model
      return await ai.general.generate(description);
    } catch (fallbackError) {
      // Fallback 2: Use template-based approach
      return await generateFromTemplate(description);
    }
  }
}
```

### 5. Atomic Operations

Changes should be atomic and reversible:

```typescript
interface FileChange {
  path: string;
  originalContent: string;
  newContent: string;
}

async function applyChangeAtomic(change: FileChange): Promise<void> {
  try {
    // Apply change
    await writeFile(change.path, change.newContent);

    // Verify
    const actual = await readFile(change.path);
    if (actual !== change.newContent) {
      throw new Error('File content mismatch after write');
    }
  } catch (error) {
    // Rollback: Restore original content
    await writeFile(change.path, change.originalContent);
    throw new RollbackError('Change failed, rolled back', { cause: error });
  }
}
```

## Error Classification

### Transient Errors (Retry)

- Network timeouts
- Rate limiting (429, 503)
- Temporary service unavailability (502, 504)
- Lock contention (database)

### Permanent Errors (Don't Retry)

- Authentication failures (401, 403)
- Not found (404)
- Validation errors (400, 422)
- Permission denied (403)

### Application Errors (Fallback)

- Model output parsing errors
- Invalid tool parameters
- Unexpected tool responses

## Error Handling Patterns

### 1. Try-Catch-Finally

Ensure cleanup even on errors:

```typescript
async function processOrder(order: Order): Promise<void> {
  let connection: DatabaseConnection;

  try {
    connection = await database.connect();
    await connection.save(order);
  } catch (error) {
    throw new OrderProcessingError('Failed to process order', { cause: error });
  } finally {
    if (connection) {
      await connection.close(); // Always close
    }
  }
}
```

### 2. Error Aggregation

Collect multiple errors and report together:

```typescript
interface BatchError extends Error {
  errors: Error[];
  summary: string;
}

async function processBatch<T>(items: T[]): Promise<T[]> {
  const errors: Error[] = [];
  const results: T[] = [];

  for (const item of items) {
    try {
      const result = await processItem(item);
      results.push(result);
    } catch (error) {
      errors.push(error);
      results.push(null); // Preserve order
    }
  }

  if (errors.length > 0) {
    const batchError = new BatchError(
      `Failed to process ${errors.length}/${items.length} items`,
      { errors, summary: `${errors.length} errors in batch` }
    );
    // Return partial results with error info
    (batchError as any).partialResults = results;
    throw batchError;
  }

  return results;
}
```

### 3. Context Recovery

Save state before operations, restore on error:

```typescript
class StateManager {
  private snapshots = new Map<string, any>();

  async withState<T>(
    key: string,
    operation: () => Promise<T>
  ): Promise<T> {
    // Save state before operation
    const beforeState = await this.captureState(key);
    this.snapshots.set(key, beforeState);

    try {
      const result = await operation();
      return result;
    } catch (error) {
      // Restore state on error
      await this.restoreState(key, beforeState);
      throw new StateRestoredError('Restored state before error', { cause: error });
    }
  }

  private async captureState(key: string): Promise<any> {
    // Capture current state
    return {
      files: await this.getFileHashes(),
      database: await this.getDatabaseState(),
      environment: process.env
    };
  }

  private async restoreState(key: string, state: any): Promise<void> {
    // Restore previous state
    await this.restoreFileHashes(state.files);
    await this.restoreDatabaseState(state.database);
    Object.assign(process.env, state.environment);
  }
}
```

### 4. Progressive Degradation

Fall back to simpler approaches:

```typescript
async function smartAnalysis(codebase: Codebase): Promise<Analysis> {
  try {
    // Try full analysis
    return await deepAnalysis(codebase);
  } catch (error) {
    log.warn('Full analysis failed, degrading to surface analysis', { error });

    try {
      // Fallback: Surface analysis only
      return await surfaceAnalysis(codebase);
    } catch (surfaceError) {
      log.warn('Surface analysis failed, using basic analysis', { error: surfaceError });

      // Fallback: Basic metrics only
      return await basicAnalysis(codebase);
    }
  }
}
```

## Recovery Strategies by Error Type

### API Failures

```typescript
async function callAPIWithRecovery(
  endpoint: string,
  params: any
): Promise<any> {
  try {
    return await api.call(endpoint, params);
  } catch (error) {
    if (error.status === 401) {
      // Auth error: Refresh and retry
      await api.refreshToken();
      return await api.call(endpoint, params);
    }

    if (error.status === 429) {
      // Rate limit: Wait and retry
      await sleep(error.retryAfter * 1000);
      return await api.call(endpoint, params);
    }

    if (error.status >= 500) {
      // Server error: Retry with backoff
      return await executeWithRetry(
        () => api.call(endpoint, params),
        { maxAttempts: 3, baseDelay: 2000 }
      );
    }

    throw error;
  }
}
```

### Code Generation Failures

```typescript
async function generateWithRetry(
  description: string,
  maxAttempts = 3
): Promise<string> {
  let attempt = 0;

  while (attempt < maxAttempts) {
    try {
      const code = await ai.generate(description);

      // Verify generated code
      const { valid, errors } = await verifyCode(code);
      if (valid) {
        return code;
      }

      throw new CodeVerificationError('Generated code failed verification', { errors });
    } catch (error) {
      attempt++;
      if (attempt >= maxAttempts) {
        // Final fallback: Use template
        return await getTemplate(description);
      }

      // Adjust prompt and retry
      const adjustedPrompt = adjustPromptForError(description, error);
      await sleep(1000 * attempt); // Progressive delay
    }
  }

  throw new Error('Failed to generate code after all attempts');
}
```

### Test Failures

```typescript
async function runTestsWithRecovery(
  testSuite: TestSuite
): Promise<TestResults> {
  try {
    return await testRunner.run(testSuite);
  } catch (error) {
    if (error instanceof EnvironmentError) {
      // Fix environment and retry
      await setupTestEnvironment();
      return await testRunner.run(testSuite);
    }

    if (error instanceof TimeoutError) {
      // Increase timeout and retry
      return await testRunner.run(testSuite, { timeout: 2 * error.timeout });
    }

    throw error;
  }
}
```

## Logging and Monitoring

### Error Logging

```typescript
interface ErrorLog {
  timestamp: string;
  error: Error;
  context: any;
  recovery?: string;
}

class ErrorLogger {
  private logs: ErrorLog[] = [];

  log(error: Error, context: any, recovery?: string): void {
    this.logs.push({
      timestamp: new Date().toISOString(),
      error,
      context,
      recovery
    });

    // Also log to external monitoring
    monitoring.reportError(error, context, recovery);
  }

  getStats(): ErrorStats {
    return {
      total: this.logs.length,
      byType: this.groupByType(),
      recoveryRate: this.calculateRecoveryRate(),
      commonPatterns: this.identifyPatterns()
    };
  }
}
```

## Best Practices

### ✅ DO

1. **Handle errors explicitly**: Never use empty catch blocks
2. **Provide context**: Include relevant state and parameters
3. **Implement retry**: For transient errors only
4. **Use fallbacks**: Alternative approaches when primary fails
5. **Make operations atomic**: Changes can be rolled back
6. **Log everything**: Timestamp, error, context, recovery action
7. **Recover gracefully**: Restore state, don't leave system corrupted

### ❌ DON'T

1. **Suppress errors**: Empty catch blocks or `@ts-ignore`
2. **Retry everything**: Only retry transient errors
3. **Silent failures**: Return without error indication
4. **Partial recovery**: Restore some state but not all
5. **No logging**: Failures without context are undebuggable
6. **Infinite loops**: Retry without max attempts
7. **Blocking operations**: Don't await on errors that should fail fast

## Implementation Checklist

- [ ] All errors caught and logged with context
- [ ] Transient errors retried with exponential backoff
- [ ] Circuit breaker for repeated failures
- [ ] Fallback mechanisms implemented
- [ ] Atomic operations with rollback
- [ ] State saved before risky operations
- [ ] Progressive degradation implemented
- [ ] Error types classified (transient/permanent)
- [ ] Recovery rate monitored and reported
- [ ] No silent failures or suppressed errors

## References

- [Retry Patterns](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/)
- [Circuit Breaker Pattern](https://martinfowler.com/bliki/CircuitBreaker)
- [Error Handling Best Practices](https://learn.microsoft.com/en-us/azure/architecture/patterns/retry-guidance)

---

**Version**: 1.0
**Last Updated**: 2025-01-12
