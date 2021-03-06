---
title: blacksmith generate migration
enterprise: true
docker: false
---

# `blacksmith generate migration`

This command generates `up` and `down` files for a SQL migration.

**Example:**
```bash
$ blacksmith generate migration --name add_user_id

```

**Related ressources:**
- Advanced practices >
  [Migrations management](/blacksmith/practices/management/migrations)
- CLI reference >
  [`migrations ack`](/blacksmith/cli/migrations-ack)
- CLI reference >
  [`migrations run`](/blacksmith/cli/migrations-run)
- CLI reference >
  [`migrations rollback`](/blacksmith/cli/migrations-rollback)

## Required flags

- `--name [migration]`: Set the name of the migration. It shall be an overview of
  the migration's changes. It shall be a valid name only containing lowercase letters
  (`a-z`), underscores (`_`), and dashes (`-`).

  **Example:**
  ```bash
  $ blacksmith generate migration --name add_user_id

  ```

## Optional flags

- `--path [path]`: Relative path where the migration's files will be generated.
  If the directories don't exist, they will automatically be created (if possible).
  If files already exist at the path with the same name of the ones generated by
  the CLI, an error will be prompted. The CLI will never override existing files.

  **Example:**
  ```bash
  $ blacksmith generate migration --name add_user_id --path ./destinations/crm/migrations

  ```

- `--no-comments`: Remove the comments on the files generated.

  **Example:**
  ```bash
  $ blacksmith generate migration --name add_user_id --no-comments

  ```
