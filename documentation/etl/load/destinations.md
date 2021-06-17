---
title: Destinations
enterprise: false
---

# Destinations

A destination is a collection of actions that Load data to a same destination.
For example, a database could be used as a data warehouse to centrally store all
the data of an organization. It would have multiple actions according to the data
to Load.

The actions registered by a destination create jobs for Loading each event's data
into the said destination. The `scheduler` service is in charge of executing these
jobs:

![Blacksmith Flows](/images/blacksmith/guides-etl.003.png)

There is an infinity of possibilities and workarounds for Loading data to
destinations. Instead of locking users into a few limited patterns and still not
covering every needs you might have, we offer a collection of [production-ready
modules](/blacksmith/modules). It simplifies development, enforces best practices,
and still allows a complete freedom on how data is Loaded.

The use of these modules is optional but higly recommended. The interfaces exposed
by the Blacksmith Go API let you create your own destinations and actions if needed.

This guide only focus on using an existing destination.

## Register a destination

In the following example, we register a new SQL database as a destination using
the `sqlike` module. This module has a `sqlikedestination` package, making any
SQL database a ready-to-use destination for Blacksmith:
```go
package main

import (
  "database/sql"

  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/flow/destination"

  "github.com/nunchistudio/blacksmith-modules/sqlike/sqlikedestination"
)

func Init() *blacksmith.Options {

  db, _ := sql.Open("<driver>", "<connection>")

  var options = &blacksmith.Options{

    // ...

    Destinations: []destination.Destination{
      sqlikedestination.New(&sqlikedestination.Options{
        Realtime:   true,
        Name:       "mypostgres",
        DB:         db,
        Migrations: []string{"migrations", "mypostgres"},
      }),
    },
  }

  return options
}

```

The destination is now accessible by using the `sqlike(mypostgres)` identifier
when one is required.
