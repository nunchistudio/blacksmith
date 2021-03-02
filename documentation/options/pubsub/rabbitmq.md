---
title: Pub / Sub with RabbitMQ
enterprise: false
---

# Pub / Sub with RabbitMQ

The RabbitMQ pub / sub adapter allows to subscribe to queues and therefore extract
data from incoming messages.

The adapter is also used for realtime communication between the gateway and scheduler
services, [as described in the introduction](/blacksmith/introduction/what/how).

## Application configuration

To use RabbitMQ as the pub / sub adapter for your application, you must set the
`From` key to `rabbitmq` in `*pubsub.Options`:
```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/adapter/pubsub"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    PubSub: &pubsub.Options{
      From:         "rabbitmq",
      Topic:        "blacksmith",
      Subscription: "blacksmith",
      Connection:   "amqp://guest:guest@127.0.0.1:5672/",
    },
  }

  return options
}

```

### Application options

- `Topic`: The producer exchange used by the gateway to forward jobs in realtime
  to the scheduler.

  **Required:** yes

- ` Subscription`: The queue used to receive messages. In distributed environments,
  you first need to make sure the queue is configured to have a single active consumer.
  This allows to receive messages only once, and therefore do not have duplicated
  data.

  **Required:** yes

- `Connection`: The RabbitMQ server URL. When set, this will override the
  `RABBIT_SERVER_URL` environment variable. **We strongly recommend the use of the
  `RABBIT_SERVER_URL` environment variable to avoid connection strings in your
  code.**

  **Required:** no

### Environment variables

Additional details must be passed to the Apache Kafka adapter. They will be loaded
from the environment variables.

- `RABBIT_SERVER_URL`: The RabbitMQ server URL to dial to leverage pub / sub. If
  `Options.PubSub.Connection` is set, it will override and be used in replacement
  of the existing environment variable.

  **Type:** `string`

  **Required:** yes (if `Options.PubSub.Connection` is not set)

  **Example:** `amqp://guest:guest@127.0.0.1:5672/`

  **Order:** options, environment variable

## Trigger configuration

Using the trigger mode `source.ModeSubscription`, a trigger can extract events from
RabbitMQ:
```go
func (t MyTrigger) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeSubscription,
    UsingSubscription: &source.Subscription{
      Subscription: "my-queue",
    },
  }
}

```

The trigger will receive in realtime every events of a RabbitMQ queue. Each event
can then be transformed and loaded to destinations.

### Subscription options

- `Subscription`: The queue to subscribe to for receiving messages.

  **Required:** yes
