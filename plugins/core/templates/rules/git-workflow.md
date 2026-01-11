---
title: Git Workflow
description: Standardized git workflow for Vendatta projects
globs: ["**/*"]
alwaysApply: true
source: https://www.conventionalcommits.org/
---

# GIT WORKFLOW

- Use feature branches: `feature/xxx`
- Commit messages MUST follow Conventional Commits: `feat:`, `fix:`, `docs:`, `chore:`, etc.
- Always use `git worktree` via Vendatta for isolation (Ref: [git-worktree-runner](https://github.com/coderabbitai/git-worktree-runner)).
- NEVER push to `main` directly.
