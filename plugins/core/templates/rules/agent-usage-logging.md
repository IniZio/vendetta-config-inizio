# Agent Usage Logging System

**Purpose**: Track and analyze skill, command, and rule invocations for productivity optimization.

## Overview

This system provides comprehensive logging of all agent operations to enable:
- Performance benchmarking
- Productivity analytics
- Usage pattern analysis
- Rule effectiveness measurement
- Skill invocation tracking

## Data Model

### Usage Log Entry

```typescript
interface UsageLog {
  id: string;                    // Unique identifier
  timestamp: string;             // ISO 8601 format
  agent: string;                 // Agent name (sisyphus, oracle, etc.)
  invocation: {
    type: 'skill' | 'command' | 'rule';
    name: string;                // Skill/command/rule name
    category: string;             // Category for grouping
  };
  context: {
    task: string;                // What agent was doing
    project: string;             // Current project
    files: string[];            // Files affected
  };
  outcome: {
    success: boolean;
    duration: number;            // Execution time in ms
    tokensUsed: number;          // Tokens consumed
    cost: number;                // Cost if applicable
  };
  metadata?: {
    model?: string;              // Model used (opencode/glm-4.7-free, etc.)
    sessionId?: string;          // Session identifier
    toolCalls?: string[];       // MCP tools invoked
    errors?: Error[];           // Errors encountered
  };
}
```

## Logging Implementation

### Log Storage

```typescript
// logs/usage.json
interface UsageLogStore {
  entries: UsageLog[];
  metadata: {
    version: string;
    lastUpdated: string;
  };
}

class UsageLogger {
  private store: UsageLogStore;
  private filePath: string = '.vendatta/logs/usage.json';

  async load(): Promise<UsageLogStore> {
    try {
      const data = await fs.readFile(this.filePath, 'utf-8');
      return JSON.parse(data);
    } catch (error) {
      if (error.code === 'ENOENT') {
        // New log file
        return { entries: [], metadata: { version: '1.0', lastUpdated: new Date().toISOString() } };
      }
      throw error;
    }
  }

  async log(entry: UsageLog): Promise<void> {
    const store = await this.load();
    entry.id = generateId();
    entry.timestamp = new Date().toISOString();
    store.entries.push(entry);
    await this.save(store);
  }

  async save(store: UsageLogStore): Promise<void> {
    store.metadata.lastUpdated = new Date().toISOString();
    await fs.writeFile(this.filePath, JSON.stringify(store, null, 2), 'utf-8');
  }

  async query(filters: LogFilters): Promise<UsageLog[]> {
    const store = await this.load();
    return store.entries.filter(entry => this.matches(entry, filters));
  }

  private matches(entry: UsageLog, filters: LogFilters): boolean {
    if (filters.agent && entry.agent !== filters.agent) return false;
    if (filters.startTime && new Date(entry.timestamp) < filters.startTime) return false;
    if (filters.endTime && new Date(entry.timestamp) > filters.endTime) return false;
    if (filters.category && entry.invocation.category !== filters.category) return false;
    return true;
  }
}
```

## Agent Categories

### Sisyphus (Primary Agent)
- Multi-agent orchestration
- Parallel execution management
- Context-aware decision making
- Tool delegation (oracle, explore, librarian)

### Oracle (Architecture Advisor)
- Complex system design guidance
- Performance optimization
- Code review and refactoring
- Deep technical analysis

### Explore (Codebase Search)
- Contextual grep operations
- File structure analysis
- Pattern discovery

### Librarian (Reference Search)
- External repository research
- Documentation retrieval
- Best practices discovery

### Frontend-UI-UX-Engineer (Visual Development)
- Component styling and layout
- Design system implementation
- Animation and responsive design

### Document-Writer (Documentation)
- README and guide creation
- API documentation
- Technical writing

## Skill Invocations

### Common Skills Tracked

1. **Code Analysis**
   - `analyze-codebase`: Deep code structure analysis
   - `refactor-code`: Code refactoring with optimization
   - `review-code`: Code review and quality assessment

2. **File Operations**
   - `find-files`: Search and locate files
   - `read-files`: Read multiple files in parallel
   - `edit-files`: Multi-file edits with context

3. **External Research**
   - `search-repos`: Search external repositories
   - `fetch-docs`: Retrieve documentation
   - `find-examples`: Discover implementation examples

4. **Testing**
   - `run-tests`: Execute test suites
   - `generate-coverage`: Create coverage reports
   - `e2e-testing`: End-to-end test execution

## Command Invocations

### Common Commands Tracked

1. **Development Commands**
   - `create-branch`: Create feature branch
   - `merge-code`: Merge changes
   - `run-build`: Execute build process
   - `start-dev-server`: Start development environment

2. **Project Management**
   - `create-task`: Create task/ticket
   - `update-kanban`: Move items in Kanban
   - `generate-report`: Create status reports

3. **Analysis Commands**
   - `analyze-metrics`: Calculate productivity metrics
   - `compare-benchmarks`: Compare against baseline
   - `generate-insights`: Create usage insights

## Rule Applications

### Common Rules Tracked

1. **Development Rules**
   - `typescript-conventions`: Type safety enforcement
   - `git-workflow`: Branch and commit standards
   - `agent-first-development`: Autonomous execution patterns

2. **Quality Rules**
   - `error-recovery`: Retry and fallback strategies
   - `tool-orchestration`: MCP tool management
   - `code-review`: Quality standards and patterns

3. **Performance Rules**
   - `parallel-execution`: Concurrent operation optimization
   - `context-management`: Efficient context handling
   - `error-suppression`: Prevent silent failures

## Metrics Collection

### Productivity Metrics

```typescript
interface ProductivityMetrics {
  // Agent Performance
  totalInvocations: number;
  skillInvocations: number;
  commandInvocations: number;
  ruleApplications: number;

  // Time Metrics
  averageDuration: number;
  totalDuration: number;

  // Success Metrics
  successRate: number;
  failureRate: number;

  // Cost Metrics
  totalTokensUsed: number;
  totalCost: number;

  // Efficiency Metrics
  invocationsPerHour: number;
  tokensPerTask: number;
}

class MetricsCalculator {
  async calculate(metrics: UsageLog[]): Promise<ProductivityMetrics> {
    const total = metrics.length;
    const successful = metrics.filter(m => m.outcome.success);
    const skillInvocations = metrics.filter(m => m.invocation.type === 'skill');
    const commandInvocations = metrics.filter(m => m.invocation.type === 'command');
    const ruleApplications = metrics.filter(m => m.invocation.type === 'rule');

    const totalDuration = metrics.reduce((sum, m) => sum + m.outcome.duration, 0);
    const averageDuration = totalDuration / total;

    return {
      totalInvocations: total,
      skillInvocations: skillInvocations.length,
      commandInvocations: commandInvocations.length,
      ruleApplications: ruleApplications.length,
      averageDuration,
      totalDuration,
      successRate: (successful.length / total) * 100,
      failureRate: ((total - successful.length) / total) * 100,
      totalTokensUsed: metrics.reduce((sum, m) => sum + (m.outcome.tokensUsed || 0), 0),
      totalCost: metrics.reduce((sum, m) => sum + (m.outcome.cost || 0), 0),
      invocationsPerHour: this.calculateInvocationsPerHour(metrics),
      tokensPerTask: totalTokensUsed / total,
    };
  }

  private calculateInvocationsPerHour(metrics: UsageLog[]): number {
    const timeRange = this.getTimeRange(metrics);
    const hours = timeRange.duration / (1000 * 60 * 60);
    return metrics.length / hours;
  }

  private getTimeRange(metrics: UsageLog[]): { start: Date; end: Date; duration: number } {
    const timestamps = metrics.map(m => new Date(m.timestamp));
    const start = new Date(Math.min(...timestamps));
    const end = new Date(Math.max(...timestamps));
    const duration = end.getTime() - start.getTime();
    return { start, end, duration };
  }
}
```

## Analytics and Insights

### Usage Patterns

```typescript
interface UsagePattern {
  skill: string;
  frequency: number;
  averageDuration: number;
  successRate: number;
  timeOfDay: 'morning' | 'afternoon' | 'evening' | 'night';
}

class UsageAnalyzer {
  async analyzePatterns(metrics: UsageLog[]): Promise<UsagePattern[]> {
    const skills = this.groupBySkill(metrics);
    const patterns: UsagePattern[] = [];

    for (const [skill, skillMetrics] of Object.entries(skills)) {
      const frequency = skillMetrics.length;
      const averageDuration = skillMetrics.reduce((sum, m) => sum + m.outcome.duration, 0) / frequency;
      const successRate = skillMetrics.filter(m => m.outcome.success).length / frequency * 100;

      const hour = parseInt(skillMetrics[0].timestamp.split('T')[1]);
      const timeOfDay = this.getTimeOfDay(hour);

      patterns.push({
        skill: skill as string,
        frequency,
        averageDuration,
        successRate,
        timeOfDay
      });
    }

    return patterns.sort((a, b) => b.frequency - a.frequency);
  }

  private groupBySkill(metrics: UsageLog[]): Map<string, UsageLog[]> {
    const groups = new Map();
    for (const m of metrics) {
      if (m.invocation.type === 'skill') {
        const skill = m.invocation.name;
        if (!groups.has(skill)) {
          groups.set(skill, []);
        }
        groups.get(skill)!.push(m);
      }
    }
    return groups;
  }

  private getTimeOfDay(hour: number): 'morning' | 'afternoon' | 'evening' | 'night' {
    if (hour >= 5 && hour < 12) return 'morning';
    if (hour >= 12 && hour < 17) return 'afternoon';
    if (hour >= 17 && hour < 21) return 'evening';
    return 'night';
  }
}
```

## Benchmark Comparison

### Baseline Establishment

```typescript
interface BenchmarkBaseline {
  skillsWithoutNewRules: number;
  averageDuration: number;
  successRate: number;
  tokensPerTask: number;
}

class BenchmarkComparator {
  async compare(
    baseline: UsageLog[],
    current: UsageLog[]
  ): Promise<BenchmarkComparison> {
    const baselineMetrics = await this.metricsCalculator.calculate(baseline);
    const currentMetrics = await this.metricsCalculator.calculate(current);

    return {
      baseline: baselineMetrics,
      current: currentMetrics,
      improvements: {
        duration: baselineMetrics.averageDuration - currentMetrics.averageDuration,
        successRate: currentMetrics.successRate - baselineMetrics.successRate,
        tokensPerTask: baselineMetrics.tokensPerTask - currentMetrics.tokensPerTask,
      },
      percentImprovement: {
        duration: ((baselineMetrics.averageDuration - currentMetrics.averageDuration) / baselineMetrics.averageDuration) * 100,
        successRate: currentMetrics.successRate - baselineMetrics.successRate,
        tokensPerTask: ((baselineMetrics.tokensPerTask - currentMetrics.tokensPerTask) / baselineMetrics.tokensPerTask) * 100,
      },
    };
  }
}
```

## Integration with Opencode CLI

### Usage Tracking Hook

```typescript
// Automatically invoked before/after skill execution
export class OpencodeUsageTracker {
  private logger: UsageLogger;
  private currentTask?: string;

  async beforeExecution(skill: string, task: string): Promise<void> {
    this.currentTask = task;
    await this.logger.log({
      agent: 'sisyphus',
      invocation: {
        type: 'skill',
        name: skill,
        category: this.categorizeSkill(skill),
      },
      context: {
        task,
        project: process.cwd(),
      },
      outcome: { success: false, duration: 0 },
    });
  }

  async afterExecution(skill: string, success: boolean, duration: number, tokensUsed?: number): Promise<void> {
    if (!this.currentTask) return;

    await this.logger.log({
      agent: 'sisyphus',
      invocation: {
        type: 'skill',
        name: skill,
        category: this.categorizeSkill(skill),
      },
      context: {
        task: this.currentTask,
        project: process.cwd(),
      },
      outcome: {
        success,
        duration,
        tokensUsed,
      },
    });

    this.currentTask = undefined;
  }

  private categorizeSkill(skill: string): string {
    const categoryMap: Record<string, string> = {
      'analyze-codebase': 'code-analysis',
      'refactor-code': 'code-analysis',
      'review-code': 'code-review',
      'find-files': 'file-operations',
      'read-files': 'file-operations',
      'edit-files': 'file-operations',
      'search-repos': 'external-research',
      'fetch-docs': 'external-research',
      'find-examples': 'external-research',
      'run-tests': 'testing',
      'generate-coverage': 'testing',
      'e2e-testing': 'testing',
    };

    return categoryMap[skill] || 'other';
  }
}
```

## Reporting

### Daily Summary Report

```typescript
interface DailySummary {
  date: string;
  totalInvocations: number;
  skillInvocations: number;
  commandInvocations: number;
  ruleApplications: number;
  averageDuration: number;
  successRate: number;
  topSkills: { skill: string; count: number }[];
  insights: string[];
}

class ReportGenerator {
  async generateDailySummary(date: Date): Promise<DailySummary> {
    const startOfDay = new Date(date.getFullYear(), date.getMonth(), date.getDate());
    const endOfDay = new Date(startOfDay.getTime() + 24 * 60 * 60 * 1000 - 1);

    const metrics = await this.logger.query({
      startTime: startOfDay,
      endTime: endOfDay,
    });

    const summary = await this.metricsCalculator.calculate(metrics);
    const topSkills = this.getTopSkills(metrics);

    return {
      date: startOfDay.toISOString().split('T')[0],
      totalInvocations: summary.totalInvocations,
      skillInvocations: summary.skillInvocations,
      commandInvocations: summary.commandInvocations,
      ruleApplications: summary.ruleApplications,
      averageDuration: summary.averageDuration,
      successRate: summary.successRate,
      topSkills,
      insights: this.generateInsights(summary, topSkills),
    };
  }

  private getTopSkills(metrics: UsageLog[], topN = 5): { skill: string; count: number }[] {
    const skillCounts = new Map<string, number>();
    for (const m of metrics) {
      if (m.invocation.type === 'skill') {
        const count = skillCounts.get(m.invocation.name) || 0;
        skillCounts.set(m.invocation.name, count + 1);
      }
    }

    return Array.from(skillCounts.entries())
      .sort((a, b) => b[1] - a[1])
      .slice(0, topN)
      .map(([skill, count]) => ({ skill, count }));
  }

  private generateInsights(metrics: ProductivityMetrics, topSkills: any[]): string[] {
    const insights: string[] = [];

    if (metrics.successRate < 80) {
      insights.push(`⚠️ Low success rate: ${metrics.successRate.toFixed(1)}%`);
    }

    if (metrics.averageDuration > 30000) {
      insights.push(`⚠️ High average duration: ${(metrics.averageDuration / 1000).toFixed(1)}s`);
    }

    if (metrics.invocationsPerHour < 5) {
      insights.push(`ℹ️ Low invocations per hour: ${metrics.invocationsPerHour.toFixed(1)}`);
    }

    const topSkill = topSkills[0];
    insights.push(`🏆 Most used skill: ${topSkill.skill} (${topSkill.count} times)`);

    return insights;
  }
}
```

## CLI Commands

### Usage Summary

```bash
#!/usr/bin/env node
// cli/usage-summary.js

import { UsageLogger, MetricsCalculator, ReportGenerator } from './lib/usage-logger';

const logger = new UsageLogger();
const calculator = new MetricsCalculator();
const reporter = new ReportGenerator();

async function main() {
  const command = process.argv[2];

  switch (command) {
    case 'summary':
      const summary = await reporter.generateDailySummary(new Date());
      console.log(JSON.stringify(summary, null, 2));
      break;

    case 'metrics':
      const today = new Date();
      const endOfDay = new Date(today.getFullYear(), today.getMonth(), today.getDate() + 1);
      const metrics = await logger.query({ startTime: today, endTime: endOfDay });
      const calc = await calculator.calculate(metrics);
      console.log(JSON.stringify(calc, null, 2));
      break;

    case 'benchmark':
      const baselineDays = parseInt(process.argv[3]) || 7;
      const today = new Date();
      const baselineStart = new Date(today.getTime() - baselineDays * 24 * 60 * 60 * 1000);
      const baseline = await logger.query({ startTime: baselineStart, endTime: today });
      const current = await logger.query({ startTime: today, endTime: new Date(today.getTime() + 24 * 60 * 60 * 1000) });
      const comparison = await calculator.compare(baseline, current);
      console.log(JSON.stringify(comparison, null, 2));
      break;

    default:
      console.error(`Unknown command: ${command}`);
      console.log('Available commands: summary, metrics, benchmark <days>');
      process.exit(1);
  }
}

main();
```

### Usage Patterns

```bash
#!/usr/bin/env node
// cli/usage-patterns.js

import { UsageLogger, UsageAnalyzer } from './lib/usage-logger';

const logger = new UsageLogger();
const analyzer = new UsageAnalyzer();

async function main() {
  const days = parseInt(process.argv[2]) || 7;
  const today = new Date();
  const startDate = new Date(today.getTime() - days * 24 * 60 * 60 * 1000);

  const metrics = await logger.query({ startTime: startDate, endTime: today });
  const patterns = await analyzer.analyzePatterns(metrics);

  console.log('Usage Patterns:');
  for (const pattern of patterns) {
    console.log(`\n${pattern.skill}:`);
    console.log(`  Frequency: ${pattern.frequency} times`);
    console.log(`  Duration: ${(pattern.averageDuration / 1000).toFixed(1)}s`);
    console.log(`  Success: ${pattern.successRate.toFixed(1)}%`);
    console.log(`  Peak Time: ${pattern.timeOfDay}`);
  }
}

main();
```

## Configuration

### Logger Config

```yaml
# vendatta-logger.yaml
version: 1

logging:
  enabled: true
  storagePath: .vendatta/logs/usage.json
  retentionDays: 30

analytics:
  enabled: true
  metrics:
    - duration
    - successRate
    - tokensUsed
    - invocationsPerHour

  reporting:
    dailySummary: true
    generatePatterns: true
    benchmarkComparison: true

skills:
  track:
    - code-analysis
    - file-operations
    - external-research
    - testing

  commands:
    track:
      - create-branch
      - merge-code
      - run-build
      - analyze-metrics

  rules:
    track:
      - typescript-conventions
      - git-workflow
      - agent-first-development
      - error-recovery
      - tool-orchestration-mcp
```

## Usage

### Enable Logging

```bash
# Add to .vendatta/config.yaml
extends:
  - inizio/vendetta-config-inizio

plugins:
  - usage-logger
```

### Track Agent Operations

```typescript
// In agent implementation
import { OpencodeUsageTracker } from '@vendatta/usage-logger';

const tracker = new OpencodeUsageTracker();

async function implementFeature(description: string) {
  await tracker.beforeExecution('analyze-codebase', 'Implement feature');

  try {
    // Agent operations
    const code = await generateCode(description);
    await validateCode(code);
    await writeCode(code);

    await tracker.afterExecution('analyze-codebase', true, performance.now() - startTime, 1500);
  } catch (error) {
    await tracker.afterExecution('analyze-codebase', false, performance.now() - startTime, 0);
    throw error;
  }
}
```

## Benefits

### 1. Performance Tracking

- Measure skill/command execution times
- Identify bottlenecks
- Optimize agent workflows
- Reduce task completion time

### 2. Success Rate Monitoring

- Track successful vs failed operations
- Identify error-prone operations
- Improve rule effectiveness

### 3. Token Usage Analysis

- Track token consumption
- Calculate cost per task
- Optimize prompts to reduce tokens

### 4. Usage Pattern Discovery

- Identify frequently used skills/commands
- Find peak productivity times
- Optimize agent scheduling

### 5. Benchmark Comparison

- Compare with/without rank #1 rules
- Measure productivity gains
- Validate rule effectiveness

### 6. Productivity Insights

- Generate daily summaries
- Create trend reports
- Provide actionable recommendations

## Anti-Patterns

### ❌ What to Avoid

1. **Selective Logging**: Only logging successful operations
2. **Incomplete Context**: Missing task/project/context information
3. **No Timestamps**: Without accurate time tracking
4. **Silent Failures**: Not logging errors
5. **No Categories**: Skills/commands/rules not grouped

### ✅ What to Do

1. **Complete Logging**: Log all skill/command/rule invocations
2. **Rich Context**: Include task, project, files, session
3. **Precise Timing**: Start and end timestamps
4. **Error Tracking**: Log failures with context
5. **Consistent Categorization**: Use standard categories
6. **Metrics Collection**: Duration, success rate, tokens, cost

## Integration Checklist

- [ ] Usage logger module created and tested
- [ ] CLI commands implemented (summary, metrics, benchmark, patterns)
- [ ] Analytics engine (calculator, analyzer, reporter)
- [ ] Opencode CLI integration hooks
- [ ] Vendetta config YAML updated
- [ ] Documentation and examples provided
- [ ] Daily/weekly automated reports configured
- [ ] Benchmark baseline established
- [ ] Productivity dashboard enabled

## References

- [Usage Tracking](https://opencode.dev/docs/usage-logging)
- [Metrics Collection](https://opencode.dev/docs/analytics)
- [Productivity Optimization](https://opencode.dev/docs/performance)

---

**Version**: 1.0
**Last Updated**: 2025-01-12
