---
title: Registering a SQL destination
enterprise: false
---

# Registering a SQL destination

Any Go SQL driver built on top of the standard library with `database/sql` is
supported. This includes PostgreSQL-compatible, MySQL-compatible, ClickHouse,
Snowflake, and more.

Before using a SQL database, you must first register one (or multiple) in the
[`*blacksmith.Options`](https://pkg.go.dev/github.com/nunchistudio/blacksmith?tab=doc#Options).
This is done by leveraging the `sqlike` module, which offers a unique and
consistent way to work with SQL databases.

```go
package main

import (
  "database/sql"

  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/flow/destination"

  "github.com/nunchistudio/blacksmith-modules/sqlike/sqlikedestination"
)

func Init() *blacksmith.Options {

  clientA, _ := sql.Open("client", "connection")
  clientB, _ := sql.Open("client", "connection")

  var options = &blacksmith.Options{

    // ...

    Destinations: []destination.Destination{
      sqlikedestination.New(&sqlikedestination.Options{
        DB:         clientA,
        Name:       "mypostgres",
        Migrations: []string{"mypostgres", "migrations"},
      }),
      sqlikedestination.New(&sqlikedestination.Options{
        DB:         clientB,
        Name:       "mysnowflake",
        Migrations: []string{"mysnowflake", "migrations"},
      }),
    },
  }

  return options
}

```

The destinations are now accessible by using the `sqlike(mypostgres)` and
`sqlike(mysnowflake)` identifiers when one is required.
