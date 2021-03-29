---
title: Pub / Sub messages
enterprise: false
---

# Actions: Pub / Sub messages

An action can load data to a Pub / Sub mechanism. In other words, an action can
publish messages to subscribers.

The best way to get started with a topic as a destination is to generate it using
the CLI:
```bash
$ blacksmith generate destination --name mydestination \
  --starter gocloud/pubsub \
  --driver <driver>

```

The starter `gocloud/pubsub` generates a destination using the
[Go Cloud Development Kit](https://gocloud.dev/) library for communicating with
topics.

Available drivers:
- `aws/snssqs` for using AWS SNS / SQS.
- `azure/servicebus` for using Azure Service Bus.
- `google/pubsub` for using Google Pub / Sub.
- `kafka` for using Apache Kafka.
- `nats` for using NATS.
- `rabbitmq` for using RabbitMQ.

Make sure to have the required dependencies. The `go.mod` file should at least
have the followings:
```go
module github.com/<org>/<app>

go 1.16

require (
  github.com/nunchistudio/blacksmith v0.16.0
  gocloud.dev v0.22.0
)

```

If you rely on the `kafka` adapter you will also need the dedicated module:
```go
module github.com/<org>/<app>

go 1.16

require (
  github.com/nunchistudio/blacksmith v0.16.0
  gocloud.dev v0.22.0
  gocloud.dev/pubsub/kafkapubsub v0.22.0
)

```

If you rely on the `nats` adapter you will also need the dedicated module:
```go
module github.com/<org>/<app>

go 1.16

require (
  github.com/nunchistudio/blacksmith v0.16.0
  gocloud.dev v0.22.0
  gocloud.dev/pubsub/natspubsub v0.22.0
)

```

If you rely on the `rabbitmq` adapter you will also need the dedicated module:
```go
module github.com/<org>/<app>

go 1.16

require (
  github.com/nunchistudio/blacksmith v0.16.0
  gocloud.dev v0.22.0
  gocloud.dev/pubsub/rabbitpubsub v0.22.0
)

```

## Create a topic action

Once a destination with the starter `gocloud/pubsub` has been created, you can
create as much actions as you need within the destination based on the same starter.
No need to pass the driver used by the destination.
```bash
$ blacksmith generate action --name myaction \
  --starter gocloud/pubsub

```

This will generate the recommended files for a topic action, inside the working
directory.

If you prefer, you can generate the action inside a directory with the `--path`
flag:
```bash
$ blacksmith generate action --name myaction \
  --starter gocloud/pubsub \
  --path ./destinations/mydestination

```

## Usage of a topic action

In the following example, the action accesses every jobs of every events. This
can be used to parse the JSON-encoded data and stores the result using the
`Unmarshal` function of the package `encoding/json`. This way we publish to the
topic everytime we need to.

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

  "gocloud.dev/pubsub"
)

type MyAction struct {
  context context.Context
  topic   *pubsub.Topic

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

      // Try to publish a message in the topic.
      err := a.topic.Send(a.context, &pubsub.Message{})
      if err != nil {
        then <- destination.Then{
          Jobs:  []string{job.ID},
          Error: &errors.Error{
            Message:    err.Error(),
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
