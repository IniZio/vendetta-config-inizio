---
name: usage-logger
description: Track and analyze agent usage for productivity optimization
version: 1.0.0
---

# Usage Logger and Metrics

Track all agent operations (skills, commands, rules) to enable:
- Performance benchmarking
- Productivity analytics
- Usage pattern analysis
- Rule effectiveness measurement

## CLI Commands

### Daily Summary
```bash
vendatta usage summary [date]
```
Generate daily summary of usage metrics and insights.

**Example**:
```bash
# Today's summary
vendatta usage summary

# Specific date
vendatta usage summary 2025-01-12
```

### Productivity Metrics
```bash
vendatta usage metrics [days]
```
Calculate detailed productivity metrics for specified time period.

**Example**:
```bash
# Last 7 days (default)
vendatta usage metrics

# Last 30 days
vendatta usage metrics 30
```

### Usage Patterns
```bash
vendatta usage patterns [days]
```
Analyze usage patterns and trends over time.

**Example**:
```bash
# Last 7 days
vendatta usage patterns

# Last 30 days
vendatta usage patterns 30
```

### Benchmark Comparison
```bash
vendatta usage benchmark <baseline-days> <current-days>
```
Compare productivity between baseline and current periods.

**Example**:
```bash
# Compare first 7 days vs last 7 days
vendatta usage benchmark 7 7
```

## Metrics Tracked

- **Agent Performance**: Total invocations, skill/command/rule breakdown
- **Time Metrics**: Average duration, total duration, invocations/hour
- **Success Metrics**: Success rate, failure rate
- **Cost Metrics**: Total tokens used, total cost, tokens/task
- **Efficiency Metrics**: Invocations per hour, tokens per task

## Storage

Usage logs stored in JSON format at:
```
.vendatta/logs/usage.json
```

## Integration

Works seamlessly with:
- Task #15: Agent Usage Logging System
- Vendatta CLI workspace commands
- AI agent configurations

## Best Practices

1. **Track Everything**: Log all skill/command/rule invocations
2. **Rich Context**: Include task, project, files in logs
3. **Precise Timing**: Use accurate timestamps
4. **Error Tracking**: Log failures with context
5. **Regular Analysis**: Review metrics weekly/monthly
6. **Benchmark Baselines**: Establish before major changes
7. **Measure Impact**: Compare before/after optimizations

## Optimization Goals

Based on metrics analysis:
- Reduce average task duration by 20%
- Achieve 90%+ success rate
- Reduce tokens per task by 15%
- Increase invocations per hour from 5 to 10+
