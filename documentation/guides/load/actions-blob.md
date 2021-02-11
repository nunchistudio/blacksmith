---
title: Blob storage
enterprise: false
---

# Actions: Blob storages

An action can load data into a blob storage, by writing into a specific key of
a bucket.

The best way to get started with a blob storage as a destination is to generate
it using the CLI:
```bash
$ blacksmith generate destination --name mydestination \
  --starter gocloud/blob \
  --driver <driver>

```

The starter `gocloud/blob` generates a destination using the
[Go Cloud Development Kit](https://gocloud.dev/) library for communicating with
blob storages.

Available drivers:
- `aws` for using Amazon S3.
- `azure` for using Azure Blog Storage.
- `google` for using Google Cloud Storage.
- `file` for using local file storage.

## Create a blob action

Once a destination with the starter `gocloud/blob` has been created, you can create
as much actions as you need within the destination based on the same starter. No
need to pass the driver used by the destination.
```bash
$ blacksmith generate action --name myaction \
  --starter gocloud/blob

```

This will generate the recommended files for a blob action, inside the working
directory.

If you prefer, you can generate the action inside a directory with the `--path`
flag:
```bash
$ blacksmith generate action --name myaction \
  --starter gocloud/blob \
  --path ./destinations/mydestination

```

## Usage of a blob action

In the following example, the action accesses every jobs of every events. This
can be used to parse the JSON-encoded data and stores the result using the
`Unmarshal` function of the package `encoding/json`. This way we write to the
bucket everytime we need to.

The scheduler will automatically mark the jobs as "succeeded", "failed", or
"discarded" given the error and the number of retries for each job.

```go
package mydestination

import (
  "context"
  "time"

  "github.com/nunchistudio/blacksmith/adapter/store"
  "github.com/nunchistudio/blacksmith/flow/destination"
  "github.com/nunchistudio/blacksmith/helper/errors"

  "gocloud.dev/blob"
)

type MyAction struct {
  context context.Context
  bucket  *blob.Bucket

  Version string        `json:"version,omitempty"`
  Context *Context      `json:"context"`
  Data    *MyActionData `json:"data"`
  SentAt  *time.Time    `json:"sent_at,omitempty"`
}

func (a MyAction) Load(tk *destination.Toolkit, queue *store.Queue, then chan<- destination.Then) {

  // We can go through every events received from the queue and their
  // related jobs. The queue can contain one or many events. The jobs
  // present in the events are specific to this action only.
  //
  // This allows to parse everything needed. This can also be useful for
  // making a request to the destination for each event / job if the
  // destination does not allow batch loads.
  //
  // Do not forget to add the job ID everytime you write in Then.
  // Otherwise the scheduler can not be aware of the job status and will
  // mark it as "unknown".
  for _, event := range queue.Events {
    for _, job := range event.Jobs {

      // ...

      // Try to write some data in the bucket.
      writer, _ := a.bucket.NewWriter(a.context, "content", nil)
      err := writer.Close()
      if err != nil {
        then <- destination.Then{
          Jobs:  []string{job.ID},
          Error: &errors.Error{
            Message: err.Error(),
          },
          ForceDiscard: false,
          OnFailed:     []destination.Action{},
          OnDiscarded:  []destination.Action{},
        }

        continue
      }

      // Finally, inform the scheduler about the success.
      then <- destination.Then{
        Jobs:        []string{job.ID},
        OnSucceeded: []destination.Action{},
      }
    }
  }
}

```
