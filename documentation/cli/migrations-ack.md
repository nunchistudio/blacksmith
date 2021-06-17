---
title: blacksmith migrations ack
enterprise: true
docker: true
---

# `blacksmith migrations ack`

This command acknowledges migrations into the `wanderer` adapter. Once acknowledged,
they can be run (executing the `up` logic) and rollbacked (using the `down` logic).

**Example:**
```bash
$ blacksmith migrations ack --scope "destination:sqlike(warehouse)"

```

**Related ressources:**
- Advanced practices >
  [Migrations management](/blacksmith/practices/management/migrations)
- CLI reference >
  [`generate migration`](/blacksmith/cli/generate-migration)
- CLI reference >
  [`migrations run`](/blacksmith/cli/migrations-run)
- CLI reference >
  [`migrations rollback`](/blacksmith/cli/migrations-rollback)

## Optional flags

- `--scope [scope]`: Scope(s) to acknowledge the migrations for. Multiple scopes
  can be passed if needed. If no scope is provided, migrations of all sources,
  triggers, destinations, and actions will be acknowledged. In the following
  example, only the new migrations for the destinations `mypostgres` and
  `warehouse` will be acknowledged.

  **Example:**
  ```bash
  $ blacksmith migrations ack --scope "destination:sqlike(mypostgres)" \
    --scope "destination:sqlike(warehouse)"

  ```

- `--build`: Build the application before acknowledging migrations. This is useful
  if you registered new sources, triggers, destinations, or actions leveraging
  migrations that were not registered at the last build.

  **Example:**
  ```bash
  $ blacksmith migrations ack --build

  ```

- `--no-cache`: Do not use the Docker cache when building the application.

  **Example:**
  ```bash
  $ blacksmith migrations ack --build --no-cache

  ```
