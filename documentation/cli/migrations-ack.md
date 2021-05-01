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
$ blacksmith migrations ack --scope destination:warehouse

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
  example, only the new migrations for the source `postgres` and for the destination
  `warehouse` will be acknowledged.

  **Example:**
  ```bash
  $ blacksmith migrations ack --scope source:postgres --scope destination:warehouse

  ```

- `--no-cache`: Do not use the Docker cache when building the application.

  **Example:**
  ```bash
  $ blacksmith migrations ack --no-cache

  ```