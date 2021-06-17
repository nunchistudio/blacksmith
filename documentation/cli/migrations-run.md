---
title: blacksmith migrations run
enterprise: true
docker: true
---

# `blacksmith migrations run`

This command runs acknowledged migrations by executing their `up` logic.

**Example:**
```bash
$ blacksmith migrations run --scope destination:warehouse

```

**Related ressources:**
- Advanced practices >
  [Migrations management](/blacksmith/practices/management/migrations)
- CLI reference >
  [`generate migration`](/blacksmith/cli/generate-migration)
- CLI reference >
  [`migrations ack`](/blacksmith/cli/migrations-ack)
- CLI reference >
  [`migrations rollback`](/blacksmith/cli/migrations-rollback)

## Optional flags

- `--version [time]`: Time representation up to which the migrations shall run.
  In the following example, every migrations acknowledged and rollbacked (but not
  discarded) up to the version `20210422135835` will run, ordered by version from
  oldest to newest.

  **Default value:** *Current timestamp formatted.*

  **Example:**
  ```bash
  $ blacksmith migrations run --version 20210422135835

  ```

- `--scope [scope]`: Scope(s) to run the migrations for. Multiple scopes can be
  passed if needed. If no scope is provided, migrations of all sources, triggers,
  destinations, and actions will run. In the following example, only the migrations
  acknowledged for the destinations `mypostgres` and `warehouse` will run.

  **Example:**
  ```bash
  $ blacksmith migrations run \
    --scope "destination:sqlike(mypostgres)" \
    --scope "destination:sqlike(warehouse)"

  ```

- `--build`: Build the application before running migrations. This is useful if
  you registered new sources, triggers, destinations, or actions leveraging
  migrations that were not registered at the last build.

  **Example:**
  ```bash
  $ blacksmith migrations run --build

  ```

- `--no-cache`: Do not use the Docker cache when building the application.

  **Example:**
  ```bash
  $ blacksmith migrations run --build --no-cache

  ```
