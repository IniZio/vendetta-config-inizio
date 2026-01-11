---
name: test-suite
description: Run all tests in the project (Go and Node.js)
steps:
  - name: Go Tests
    command: "go test ./..."
  - name: Node Tests
    command: "if [ -f package.json ]; then npm test; fi"
---

# Test Suite
Composite command to run all tests in the repository.
