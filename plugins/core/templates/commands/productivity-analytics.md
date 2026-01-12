---
name: productivity-analytics
description: Generate insights and recommendations from usage logs
version: 1.0.0
---

# Productivity Analytics and Insights

Generate actionable insights from agent usage data to optimize productivity.

## Daily Insights

Generate daily productivity insights:
```bash
vendatta usage summary
```

**Insights Include**:
- Success rate warnings (< 80% or > 90%)
- Duration analysis (high > 30s, fast < 5s)
- Invocations per hour (low < 5, high > 10)
- Most used skill identification
- Cost monitoring
- Token usage tracking

**Example Output**:
```json
{
  "insights": [
    "✅ Excellent success rate: 95.0%",
    "⚡ Fast average duration: 3.5s",
    "🚀 High invocations per hour: 12.5",
    "🏆 Most used skill: analyze-codebase (15 times)",
    "💰 Total cost: $0.25",
    "📊 Total tokens used: 250K"
  ]
}
```

## Pattern Analysis

Analyze usage patterns over time:
```bash
vendatta usage patterns 30
```

**Patterns Identified**:
- Most frequently used skills
- Average duration per skill
- Success rate per skill
- Peak productivity times (time-of-day)

**Example Output**:
```json
{
  "patterns": [
    {
      "skill": "analyze-codebase",
      "frequency": 45,
      "averageDuration": 2200.0,
      "successRate": 95.0,
      "timeOfDay": "morning"
    },
    {
      "skill": "refactor-code",
      "frequency": 23,
      "averageDuration": 3500.0,
      "successRate": 88.0,
      "timeOfDay": "afternoon"
    }
  ]
}
```

## Benchmark Comparison

Compare performance before and after optimization:
```bash
vendatta usage benchmark 7 7
```

**Metrics Compared**:
- Average duration improvement
- Success rate improvement
- Tokens per task reduction

**Example Output**:
```json
{
  "percentImprovement": {
    "duration": 20.0,
    "successRate": 7.5,
    "tokensPerTask": 20.0
  }
}
```

## Actionable Recommendations

Based on metrics analysis, prioritize:

### 1. High-Impact Skills
Focus on skills with:
- High usage frequency
- Low success rates
- Long average durations

### 2. Peak Productivity Times
Schedule work during peak hours identified by:
- Time-of-day pattern analysis
- Energy level optimization
- Max invocations per hour

### 3. Token Optimization
- Monitor tokens per task
- Identify token-heavy operations
- Optimize prompts for efficiency
- Goal: 15% reduction

### 4. Success Rate Improvement
- Analyze failed operations
- Identify error patterns
- Update rules for better guidance
- Goal: 90%+ success rate

### 5. Regular Benchmarking
- Establish baselines before changes
- Measure impact of optimizations
- Track progress over time
- Validate rule effectiveness

## Usage Workflow

1. **Initialize**: `vendatta init` (metrics automatically enabled)
2. **Work**: Use agents normally (logging happens automatically)
3. **Review**: Check daily/weekly summaries
4. **Analyze**: Look for patterns and insights
5. **Optimize**: Make data-driven improvements
6. **Benchmark**: Measure impact of changes
7. **Repeat**: Continuous improvement cycle

## Benefits

- **Data-Driven Decisions**: Make informed choices based on metrics
- **Performance Monitoring**: Track and improve over time
- **Cost Awareness**: Monitor and optimize token usage
- **Pattern Discovery**: Identify productivity insights
- **Benchmark Validation**: Measure effectiveness of changes
- **Continuous Improvement**: Ongoing optimization loop
