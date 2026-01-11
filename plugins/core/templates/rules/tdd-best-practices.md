---
title: TDD Best Practices
description: Guidelines for Test-Driven Development
globs: ["**/*_test.go", "**/*.test.ts"]
source: https://arxiv.org/abs/2510.23761
---

# TDD BEST PRACTICES

1. **RED**: Write a failing test first.
2. **GREEN**: Write the minimal code needed to make the test pass.
3. **REFACTOR**: Clean up the code while keeping the tests green.
- Aim for 80%+ coverage on new logic.
- Use `testify/assert` and `testify/require` in Go.
