---
name: playwright
description: Browser automation and web scraping via Playwright
parameters:
  type: object
  properties:
    url: { type: string, description: "URL to visit" }
    action: { type: string, description: "Action to perform (screenshot, extract, click)" }
execute:
  command: "npx"
  args: ["playwright", "test"]
source: https://playwright.dev/
---

# Playwright Skill
Allows the agent to interact with web pages using Playwright.
