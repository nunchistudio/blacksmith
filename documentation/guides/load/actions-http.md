---
title: HTTP requests
enterprise: false
---

# Actions: HTTP requests

An action can load data to any HTTP API, both entry-per-entry or in bulk depending
on the capabilities of the destination.

The best way to get started with a HTTP API as a destination is to generate it
using the CLI:
```bash
$ blacksmith generate destination --name mydestination \
  --starter net/http

```

The starter `net/http` generates a destination using the Go standard library for
communicating with HTTP APIs.

## Create a HTTP action

Once a destination with the starter `net/http` has been created, you can create
as much actions as you need within the destination based on the same starter.
```bash
$ blacksmith generate action --name myaction \
  --starter net/http

```

This will generate the recommended files for a HTTP action, inside the working
directory.

If you prefer, you can generate the action inside a directory with the `--path`
flag:
```bash
$ blacksmith generate action --name myaction \
  --starter net/http \
  --path ./destinations/mydestination

```

## Usage of a HTTP action

In the following example, the action accesses every jobs of every events. This can
be used to parse the JSON-encoded data and stores the result using the `Unmarshal`
function of the package `encoding/json`.

Then, a HTTP request is made using the client (of type `*http.Client`) passed by
the parent destination.

Finally, given the response returned by the HTTP API, we inform the scheduler of
the success or failure of the load. The scheduler will automatically mark the
jobs as "succeeded", "failed", or "discarded" given the error and the number of
retries for each job.

```go
package mydestination

import (
  "bytes"
  "context"
  "net/http"
  "time"

  "github.com/nunchistudio/blacksmith/adapter/store"
  "github.com/nunchistudio/blacksmith/flow/destination"
  "github.com/nunchistudio/blacksmith/helper/errors"
)

type MyAction struct {
  context context.Context
  client  *http.Client

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
  // destination does not allow batch loads. If so, do not forget to add
  // the job ID everytime you write in Then. Otherwise the scheduler can
  // not be aware of the job status and will mark it as "unknown".
  for _, event := range queue.Events {
    for _, job := range event.Jobs {

      // ...

    }
  }

  // When the data is parsed and ready, we can send it to the destination.
  // Inform the scheduler and stop to load data if an error occured.
  req, _ := http.NewRequest("POST", "http://example.com", nil)
  req.Header.Set("Content-Type", "application/json")
  res, err := a.client.Do(req)
  if err != nil {
    then <- destination.Then{
      Error: &errors.Error{
        StatusCode: 500,
        Message:    err.Error(),
      },
      ForceDiscard: true,
      OnFailed:     []destination.Action{},
      OnDiscarded:  []destination.Action{},
    }

    return
  }

  // Since a non-2xx status code doesn't cause an error, catch HTTP status
  // code to ensure nothing bad happened.
  if res.StatusCode >= 300 {
    buf := new(bytes.Buffer)
    buf.ReadFrom(res.Body)
    then <- destination.Then{
      ForceDiscard: false,
      Error: &errors.Error{
        StatusCode: res.StatusCode,
        Message:    buf.String(),
      },
    }

    return
  }

  // Finally, inform the scheduler about the success.
  then <- destination.Then{
    OnSucceeded: []destination.Action{},
  }
}

```
