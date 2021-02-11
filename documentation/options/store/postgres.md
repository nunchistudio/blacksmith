---
title: PostgreSQL store
enterprise: false
---

# PostgreSQL store

The PostgreSQL adapter is compatible with any PostgreSQL wire compatible database
and can work with any kind of extensions.

## Options

- `Connection`: The connection string to use for the PostgreSQL store. When set,
  this will override the `POSTGRES_STORE_URL` environment variable. **We strongly
  recommend the use of the `POSTGRES_STORE_URL` environment variable to avoid
  connection strings in your code.**

  **Required:** no

  **Example:** `postgres://user:password@127.0.0.1/database`

## Environment variables

Some options can be loaded from the environment variables.

- `POSTGRES_STORE_URL`: The PostgreSQL URL to use for the store adapter. If
  `Options.Store.Connection` is set, it will override and be used in replacement
  of this environment variable.

  **Required:** yes (if `Options.Store.Connection` is not set)

  **Example:** `postgres://user:password@127.0.0.1/database`

## Example

```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/adapter/store"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Store: &store.Options{
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
CREATE SCHEMA IF NOT EXISTS blacksmith_store;

CREATE TABLE IF NOT EXISTS blacksmith_store.events (
  id VARCHAR(27) PRIMARY KEY,
  source TEXT NOT NULL,
  trigger TEXT NOT NULL,
  version TEXT,
  context JSONB,
  data JSONB,
  sent_at TIMESTAMP WITHOUT TIME ZONE,
  received_at TIMESTAMP WITHOUT TIME ZONE,
  ingested_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blacksmith_store.jobs (
  id VARCHAR(27) PRIMARY KEY,
  destination TEXT NOT NULL,
  action TEXT NOT NULL,
  version TEXT,
  context JSONB,
  data JSONB,
  parent_job_id VARCHAR(27) REFERENCES blacksmith_store.jobs (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  event_id VARCHAR(27) NOT NULL REFERENCES blacksmith_store.events (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blacksmith_store.transitions (
  id VARCHAR(27) PRIMARY KEY,
  attempt INT4 NOT NULL,
  state_before TEXT,
  state_after TEXT NOT NULL,
  error JSONB,
  event_id VARCHAR(27) NOT NULL REFERENCES blacksmith_store.events (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  job_id VARCHAR(27) NOT NULL REFERENCES blacksmith_store.jobs (id)
    ON UPDATE CASCADE ON DELETE CASCADE
    DEFERRABLE INITIALLY DEFERRED,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

```
