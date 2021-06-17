---
title: blacksmith run operation
enterprise: false
docker: true
---

# `blacksmith run operation`

This command runs an operation on top of a SQL database.

**Example:**
```bash
$ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
  --file "./operations/demo.sql" \
  --data "./data/somedata.json" \
  --dryrun

```

**Related ressources:**
- Guides for TLT with SQL >
  [Running operations in the data warehouse](/blacksmith/tlt/operations/warehouse)

## Required flags

- `--scope [scope]`: Scope to run the operation against. Only one scope can be
  passed when running an operation.

  **Example:**
  ```bash
  $ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
    --file "./operations/demo.sql"

  ```

- `--file [filename]`: The SQL file to compile and execute against the database.

  **Example:**
  ```bash
  $ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
    --file "./operations/demo.sql"

  ```

## Optional flags

- `--data [filename]`: The CSV or JSON file to pass down to the SQL file as data
  source.

  **Example:**
  ```bash
  $ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
    --file "./operations/demo.sql" \
    --data "./data/somedata.json"

  ```

- `--dryrun`: Only compile the SQL file into `<operation>.compiled.sql` at the
  same location of the template file. This prevents the operation to actually run.

  **Example:**
  ```bash
  $ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
    --file "./operations/demo.sql" \
    --dryrun

  ```

- `--build`: Build the application before rolling back migrations. This is useful
  if you registered new sources, triggers, destinations, or actions leveraging
  migrations that were not registered at the last build.

  **Example:**
  ```bash
  $ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
    --file "./operations/demo.sql" \
    --build

  ```

- `--no-cache`: Do not use the Docker cache when building the application.

  **Example:**
  ```bash
  $ blacksmith run operation --scope "destination:sqlike(mypostgres)" \
    --file "./operations/demo.sql" \
    --build --no-cache

  ```
