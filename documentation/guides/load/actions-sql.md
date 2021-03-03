---
title: SQL databases
enterprise: false
---

# Actions: SQL databases

An action can load data to any SQL database, both entry-per-entry or within a
transaction depending on your needs.

The best way to get started with a SQL database as a destination is to generate
it using the CLI:
```bash
$ blacksmith generate destination --name mydestination \
  --starter database/sql

```

The starter `database/sql` generates a destination using the Go standard library
for communicating with SQL databases.

## Create a SQL action

Once a destination with the starter `database/sql` has been created, you can create
as much actions as you need within the destination based on the same starter.
```bash
$ blacksmith generate action --name myaction \
  --starter database/sql

```

This will generate the recommended files for a SQL action, inside the working
directory.

If you prefer, you can generate the action inside a directory with the `--path`
flag:
```bash
$ blacksmith generate action --name myaction \
  --starter database/sql \
  --path ./destinations/mydestination

```

If you need to [handle data migrations](/blacksmith/practices/management/migrations)
within the action, you can also add the `--migrations` flag:
```bash
$ blacksmith generate action --name myaction \
  --starter database/sql \
  --path ./destinations/mydestination \
  --migrations

```

## Usage of a SQL action

In the following example, we first initialize a new transaction using the SQL
client passed by the parent destination.

Then, the action accesses every jobs of every events. This can be used to parse
the JSON-encoded data and stores the result using the `Unmarshal` function of the
package `encoding/json`. This way we can add new query to execute inside the 
transaction based on some data.

Finally, we commit the transaction, and inform the scheduler of the success or
failure of the load. The scheduler will automatically mark the jobs as "succeeded",
"failed", or "discarded" given the error and the number of retries for each job.

```go
package mydestination

import (
  "context"
  "database/sql"
  "time"

  "github.com/nunchistudio/blacksmith/adapter/store"
  "github.com/nunchistudio/blacksmith/flow/destination"
  "github.com/nunchistudio/blacksmith/helper/errors"
)

type MyAction struct {
  context context.Context
  db      *sql.DB

  Version string        `json:"version,omitempty"`
  Context *Context      `json:"context"`
  Data    *MyActionData `json:"data"`
  SentAt  *time.Time    `json:"sent_at,omitempty"`
}

func (a MyAction) Load(tk *destination.Toolkit, queue *store.Queue, then chan<- destination.Then) {

  // Start a new transaction using the SQL driver.
  tx, err := a.db.Begin()
  if err != nil {
    then <- destination.Then{
      Error: &errors.Error{
        Message: err.Error(),
      },
      ForceDiscard: false,
      OnFailed:     []destination.Action{},
      OnDiscarded:  []destination.Action{},
    }

    return
  }

  // We can go through every events received from the queue and their
  // related jobs. The queue can contain one or many events. The jobs
  // present in the events are specific to this action only.
  //
  // This allows to parse everything needed. This can also be useful for
  // making a request to the destination for each event / job if the
  // destination does not allow batch loads.
  for _, event := range queue.Events {
    for _, job := range event.Jobs {

      // ...

      // For each job, add a query in the transaction.
      tx.Exec("")
    }
  }

  // When the data is parsed and ready, we can commit the transaction.
  err = tx.Commit()
  if err != nil {
    then <- destination.Then{
      Error: &errors.Error{
        Message:    err.Error(),
      },
      ForceDiscard: false,
      OnFailed:     []destination.Action{},
      OnDiscarded:  []destination.Action{},
    }

    return
  }

  // Finally, inform the scheduler about the success.
  then <- destination.Then{
    OnSucceeded: []destination.Action{},
  }
}

```
