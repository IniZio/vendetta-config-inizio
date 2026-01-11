---
title: Docker Best Practices
description: Standardized Docker configuration for Vendatta
globs: ["**/Dockerfile", "**/docker-compose.yml"]
source: https://github.com/docker/cagent
---

# DOCKER BEST PRACTICES

- Use multi-stage builds to minimize image size.
- Always include health checks in `docker-compose.yml`.
- Use specific image tags instead of `latest`.
- Avoid running processes as root inside containers.
