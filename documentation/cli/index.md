---
title: blacksmith
enterprise: false
docker: false
---

# `blacksmith`

The Blacksmith CLI allows developers and operators to manage and run a Blacksmith
application.

**Example:**
```bash
$ blacksmith --help

```

**Related ressources:**
- Getting started >
  [Installation](/blacksmith/introduction/start/install)

## Global optional flags

The following flags are shared across every commands.

- `--appdir [path]`: Set a custom relative path indicating where to find the
  Blacksmith application.

  **Default value:** `./`

  **Example:**
  ```bash
  $ blacksmith build --appdir ./app

  ```

- `--plugin [path]`: Set a custom relative path indicating where to build the Blacksmith
  application as a Go plugin.

  **Default value:** `.blacksmith/application.so`

  **Example:**
  ```bash
  $ blacksmith build --plugin .tmp/app.so

  ```

- `--help`: Prevents the command to be executed and display the help of the said
  command instead.

  **Example:**
  ```bash
  $ blacksmith migrations ack --help

  ```
