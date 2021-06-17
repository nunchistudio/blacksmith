---
title: blacksmith migrations rollback
enterprise: true
docker: true
---

# `blacksmith migrations rollback`

This command rollbacks previously run migrations by executing their `down` logic.

**Example:**
```bash
$ blacksmith migrations rollback --version 20210422135835

```

**Related ressources:**
- Advanced practices >
  [Migrations management](/blacksmith/practices/management/migrations)
- CLI reference >
  [`generate migration`](/blacksmith/cli/generate-migration)
- CLI reference >
  [`migrations ack`](/blacksmith/cli/migrations-ack)
- CLI reference >
  [`migrations run`](/blacksmith/cli/migrations-run)

## Required flags

- `--version [time]`: Time representation down to which the migrations shall be
  rollbacked. In the following example, every migrations previously run down to
  the version `20210422135835` will be rollbacked, ordered by version from
  newest to oldest.

  **Example:**
  ```bash
  $ blacksmith migrations rollback --version 20210422135835

  ```

## Optional flags

- `--discard`: Mark the migrations within the scope(s) and down to the specified
  version as `discarded`. Whereas a rollbacked migration can run again when the
  `run` command is executed, a discarded migration will not. In the following
  example, every migrations down to the version `20210422135835` will be discarded,
  meaning their `up` logic will not be executed again.

  **Example:**
  ```bash
  $ blacksmith migrations rollback --version 20210422135835 --discard

  ```

- `--scope [scope]`: Scope(s) to rollback the migrations for. Multiple scopes can
  be passed if needed. If no scope is provided, migrations of all sources, triggers,
  destinations, and actions will rollback down until the specified version is
  reached. In the following example, only the migrations previously run for the
  destinations `mypostgres` and `warehouse` will be rollbacked, until the version
  `20210422135835` is reached.

  **Example:**
  ```bash
  $ blacksmith migrations rollback --version 20210422135835 \
    --scope "destination:sqlike(mypostgres)" \
    --scope "destination:sqlike(warehouse)"

  ```

- `--build`: Build the application before rolling back migrations. This is useful
  if you registered new sources, triggers, destinations, or actions leveraging
  migrations that were not registered at the last build.

  **Example:**
  ```bash
  $ blacksmith migrations rollback --version 20210422135835 \
    --build

  ```

- `--no-cache`: Do not use the Docker cache when building the application.

  **Example:**
  ```bash
  $ blacksmith migrations rollback --version 20210422135835 \
    --build --no-cache

  ```
