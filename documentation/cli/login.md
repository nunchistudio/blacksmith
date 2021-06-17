---
title: blacksmith login
enterprise: true
docker: true
---

# `blacksmith login`

This command generates and outputs a unique link to open the Nunchi Customer Portal.
It uses the Blacksmith Enterprise License details of your application.

**Example:**
```bash
$ blacksmith login

```

## Optional flags

- `--build`: Build the application before login into the Nunchi Customer Portal.
  This is only useful if you updated the `*blacksmith.Options` with new license
  details.

  **Example:**
  ```bash
  $ blacksmith login --build

  ```

- `--no-cache`: Do not use the Docker cache when building the application.

  **Example:**
  ```bash
  $ blacksmith login --build --no-cache

  ```
