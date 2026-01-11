---
name: lint-all
description: Run all linters (golangci-lint and eslint)
steps:
  - name: Go Lint
    command: "golangci-lint run"
  - name: Node Lint
    command: "if [ -f package.json ]; then npm run lint; fi"
---

# Lint All
Composite command to run all linters in the repository.
