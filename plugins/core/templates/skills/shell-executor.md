---
name: shell-executor
description: Execute shell commands with safety checks and logging
parameters:
  type: object
  properties:
    command: { type: string, description: "Command to execute" }
execute:
  command: "bash"
  args: ["-c", "{{.command}}"]
---

# Shell Executor
Standard tool for running shell commands within the isolated environment.
