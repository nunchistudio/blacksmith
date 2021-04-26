---
title: blacksmith build
enterprise: false
docker: true
---

# `blacksmith build`

This command builds your application as a Go plugin to be executed by the CLI. The
build process is not required since it is automatically executed before starting
the application. This command is here for convenience so it can be used to validate
an application without starting it.

**Example:**
```bash
$ blacksmith build

```

**Related ressources:**
- Getting started >
  [Docker environment](/blacksmith/introduction/start/docker)
- Getting started >
  [Running an instance](/blacksmith/introduction/start/run)

## Optional flags

- `--no-cache`: Do not use the Docker cache when building the application.

  **Example:**
  ```bash
  $ blacksmith build --no-cache

  ```
