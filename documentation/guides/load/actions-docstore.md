---
title: NoSQL databases
enterprise: false
---

# Actions: NoSQL databases

An action can load data to a NoSQL database, both entry-per-entry or in bulk
depending on your needs and the capabilities of the destination.

The best way to get started with a NoSQL database as a destination is to generate
it using the CLI:
```bash
$ blacksmith generate destination --name mydestination \
  --starter gocloud/docstore \
  --driver <driver>

```

The starter `gocloud/docstore` generates a destination using the
[Go Cloud Development Kit](https://gocloud.dev/) library for communicating with
NoSQL databases.

Available drivers:
- `firestore` for using Google Firestore.
- `dynamodb` for using Amazon DynamoDB.
- `mongodb` for using MongoDB.
- `cosmosdb` for using MongoDB on Azure with CosmosDB.

Make sure to have the required dependencies. The `go.mod` file should at least
have the followings:
```go
module github.com/<org>/<app>

go 1.16

require (
  github.com/nunchistudio/blacksmith v0.15.0
  gocloud.dev v0.22.0
)

```

If you rely on the `mongodb` or `cosmosdb` adapter you will also need the dedicated
module for MongoDB:
```go
module github.com/<org>/<app>

go 1.16

require (
  github.com/nunchistudio/blacksmith v0.15.0
  gocloud.dev v0.22.0
  gocloud.dev/docstore/mongodocstore v0.22.0
)

```

## Create a NoSQL action

Once a destination with the starter `gocloud/docstore` has been created, you can create
as much actions as you need within the destination based on the same starter. No
need to pass the driver used by the destination.
```bash
$ blacksmith generate action --name myaction \
  --starter gocloud/docstore

```

This will generate the recommended files for a NoSQL action, inside the working
directory.

If you prefer, you can generate the action inside a directory with the `--path`
flag:
```bash
$ blacksmith generate action --name myaction \
  --starter gocloud/docstore \
  --path ./destinations/mydestination

```

If you need to [handle data migrations](/blacksmith/guides/practices/migrations)
within the action, you can also add the `--migrations` flag:
```bash
$ blacksmith generate action --name myaction \
  --starter gocloud/docstore \
  --path ./destinations/mydestination \
  --migrations

```

## Usage of a NoSQL action

In the following example, we first initialize a new bulk using the collection
passed by the parent destination.

Then, the action accesses every jobs of every events. This can be used to parse
the JSON-encoded data and stores the result using the `Unmarshal` function of the
package `encoding/json`. This way we can add new query to execute inside the 
bulk based on some data.

Finally, we load the the batch of documents, and inform the scheduler of the
success or failure of the load. The scheduler will automatically mark the jobs as
"succeeded", "failed", or "discarded" given the error and the number of retries
for each job.

```go
package mydestination

import (
  "context"
  "time"

  "github.com/nunchistudio/blacksmith/adapter/store"
  "github.com/nunchistudio/blacksmith/flow/destination"
  "github.com/nunchistudio/blacksmith/helper/errors"

  "gocloud.dev/docstore"
)

type MyAction struct {
  context    context.Context
  collection *docstore.Collection

  Version string        `json:"version,omitempty"`
  Context *Context      `json:"context"`
  Data    *MyActionData `json:"data"`
  SentAt  *time.Time    `json:"sent_at,omitempty"`
}

func (a MyAction) Load(tk *destination.Toolkit, queue *store.Queue, then chan<- destination.Then) {

  // Create an action list, allowing to optimize bulk load.
  bulk := a.collection.Actions()
  
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

      // For each job, create a new document.
      bulk.Create(&MyActionData{})
    }
  }

  // When the data is parsed and ready, we can load it into the destination.
  err := bulk.Do(a.context)
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

  // Finally, inform the scheduler about the success.
  then <- destination.Then{
    OnSucceeded: []destination.Action{},
  }
}

```
