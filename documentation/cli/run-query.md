---
title: blacksmith run query
enterprise: false
docker: true
---

# `blacksmith run query`

This command runs an query on top of a SQL database.

**Example:**
```bash
$ blacksmith run query --scope "destination:sqlike(mypostgres)" \
  --file "./queries/demo.sql" \
  --data "./data/somedata.json" \
  --csv --json

```

**Related ressources:**
- Guides for TLT with SQL >
  [Running queries in the data warehouse](/blacksmith/tlt/queries/warehouse)

## Required flags

- `--scope [scope]`: Scope to run the query against. Only one scope can be passed
  when running a query.

  **Example:**
  ```bash
  $ blacksmith run query --scope "destination:sqlike(mypostgres)" \
    --file "./queries/demo.sql"

  ```

- `--file [filename]`: The SQL file to compile and execute against the database.

  **Example:**
  ```bash
  $ blacksmith run query --scope "destination:sqlike(mypostgres)" \
    --file "./queries/demo.sql"

  ```

## Optional flags

- `--csv`: Download the result as a CSV file. The file is written as
  `<query>.csv` at the same location of the SQL file. This does not apply if
  the flag `--dryrun` is passed as well.

  **Example:**
  ```bash
  $ blacksmith run query --scope "destination:sqlike(mypostgres)" \
    --file "./queries/demo.sql" \
    --csv

  ```

- `--json`: Download the result as a JSON file. The file is written as
  `<query>.json` at the same location of the SQL file. This does not apply if
  the flag `--dryrun` is passed as well.

  **Example:**
  ```bash
  $ blacksmith run query --scope "destination:sqlike(mypostgres)" \
    --file "./queries/demo.sql" \
    --json

  ```

- `--data [filename]`: The CSV or JSON file to pass down to the SQL file as data
  source.

  **Example:**
  ```bash
  $ blacksmith run query --scope "destination:sqlike(mypostgres)" \
    --file "./queries/demo.sql" \
    --data "./data/somedata.json"

  ```

- `--dryrun`: Only compile the SQL file into `<query>.compiled.sql` at the same
  location of the template file. This prevents the query to actually run.

  **Example:**
  ```bash
  $ blacksmith run query --scope "destination:sqlike(mypostgres)" \
    --file "./queries/demo.sql" \
    --dryrun

  ```

- `--build`: Build the application before rolling back migrations. This is useful
  if you registered new sources, triggers, destinations, or actions leveraging
  migrations that were not registered at the last build.

  **Example:**
  ```bash
  $ blacksmith run query --scope "destination:sqlike(mypostgres)" \
    --file "./queries/demo.sql" \
    --build

  ```

- `--no-cache`: Do not use the Docker cache when building the application.

  **Example:**
  ```bash
  $ blacksmith run query --scope "destination:sqlike(mypostgres)" \
    --file "./queries/demo.sql" \
    --build --no-cache

  ```
