---
title: PostgreSQL supervisor
enterprise: true
---

# PostgreSQL supervisor

The use of PostgreSQL as the `supervisor` adapter allows to leverage distributed
locks mechanism to avoid access collision when working within a multi-node
environment.

The PostgreSQL adapter is compatible with any PostgreSQL wire compatible database
and can work with any kind of extensions.

## Options

- `Join`: The node of the distributed system to join. Each running instance of
  Blacksmith can join a different PostgreSQL node of a highly-available cluster.
  - `Address`: The connection string to use for the PostgreSQL node to join. When
    set, this will override the `POSTGRES_SUPERVISOR_URL` environment variable.
    **We strongly recommend the use of the `POSTGRES_SUPERVISOR_URL` environment
    variable to avoid connection strings in your code.**

  **Required:** no

  **Example:** `postgres://user:password@127.0.0.1/database`

## Environment variables

Some options can be loaded from the environment variables.

- `POSTGRES_SUPERVISOR_URL`: The PostgreSQL URL to use for the supervisor adapter. 
  If `Options.Supervisor.Join.Address` is set, it will override and be used in
  replacement of this environment variable.

  **Required:** yes (if `Options.Supervisor.Join.Address` is not set)

  **Example:** `postgres://user:password@127.0.0.1/database`

  **Order:** options, environment variable

## Example

```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/adapter/supervisor"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Supervisor: &supervisor.Options{
      From: "postgres",
    },
  }

  return options
}

```

## SQL migration

Before using the adapter, you first need to run the following migration:

```sql
CREATE SCHEMA IF NOT EXISTS blacksmith_supervisor;

CREATE TABLE IF NOT EXISTS blacksmith_supervisor.locks (
  key TEXT PRIMARY KEY,
  is_acquired BOOL NOT NULL DEFAULT FALSE,
  session_id VARCHAR(27),
  acquirer_name TEXT,
  acquirer_address TEXT,
  acquired_at TIMESTAMP WITHOUT TIME ZONE
);

```
