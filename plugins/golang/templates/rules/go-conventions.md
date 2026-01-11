---
title: Go Conventions
description: Standardized Go coding style for Vendatta
globs: ["**/*.go"]
source: https://github.com/golang/go/wiki/CodeReviewComments
---

# GO CONVENTIONS

- Go 1.24+ features preferred (Ref: [Go ADK](https://developers.googleblog.com/announcing-the-agent-development-kit-for-go-build-powerful-ai-agents-with-your-favorite-languages/)).
- Use `fmt.Errorf("...: %w", err)` for error wrapping.
- Table-driven tests are encouraged.
- Follow `internal/` package pattern for private logic.
- Avoid `interface{}` where possible; use Generics or specific interfaces.
