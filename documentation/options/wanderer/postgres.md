---
title: PostgreSQL wanderer
enterprise: true
---

# PostgreSQL wanderer

The PostgreSQL adapter is compatible with any PostgreSQL wire compatible database
and can work with any kind of extensions.

## Options

- `Connection`: The connection string to use for the PostgreSQL wanderer. When set,
  this will override the `POSTGRES_WANDERER_URL` environment variable. **We strongly
  recommend the use of the `POSTGRES_WANDERER_URL` environment variable to avoid
  connection strings in your code.**

  **Required:** no

  **Example:** `postgres://user:password@127.0.0.1/database`

## Environment variables

Some options can be loaded from the environment variables.

- `POSTGRES_WANDERER_URL`: The PostgreSQL URL to use for the wanderer adapter. If
  `Options.Wanderer.Connection` is set, it will override and be used in replacement
  of this environment variable.

  **Required:** yes (if `Options.Wanderer.Connection` is not set)

  **Example:** `postgres://user:password@127.0.0.1/database`

  **Order:** options, environment variable

## Example

```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/adapter/wanderer"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Wanderer: &wanderer.Options{
      From:       "postgres",
      Connection: "postgres://user:password@127.0.0.1/database",
    },
  }

  return options
}

```

## SQL migration

Before using the adapter, you first need to run the following migration:

```sql
CREATE SCHEMA IF NOT EXISTS blacksmith_wanderer;

CREATE TABLE IF NOT EXISTS blacksmith_wanderer.migrations (
  id VARCHAR(27) PRIMARY KEY,
  version TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  scope TEXT NOT NULL,
  name TEXT NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blacksmith_wanderer.transitions (
  id VARCHAR(27) PRIMARY KEY,
  state_before TEXT,
  state_after TEXT NOT NULL,
  error JSONB,
  migration_id VARCHAR(27) NOT NULL REFERENCES blacksmith_wanderer.migrations (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX migrations_version
  ON blacksmith_wanderer.migrations (version, scope);

```
