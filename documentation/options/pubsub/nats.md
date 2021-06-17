---
title: Pub / Sub with NATS
enterprise: false
---

# Pub / Sub with NATS

The NATS driver as the `pubsub` adapter allows to subscribe to topics and therefore
extract data from incoming messages.

The adapter is also used for realtime communication between the gateway and scheduler
services, [as described in the onboarding](/blacksmith/start/onboarding/how).

## Application configuration

To use NATS as the Pub / Sub adapter for your application, you must set the `From`
key to `nats` in `*pubsub.Options`:
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
      From:         pubsub.DriverNATS,
      Topic:        "blacksmith",
      Subscription: "blacksmith",
      Connection:   "nats://127.0.0.1:4222",
    },
  }

  return options
}

```

### Application options

- `Topic`: The subject used by the gateway to forward jobs in realtime to the
  scheduler.

  **Required:** yes

- ` Subscription`: The queue used by the scheduler to receive the jobs forwarded
  by the gateway. We leverage NATS queues to benefit a single active consumer
  when deploying in distributed environments. This allows to receive messages
  only once, and therefore do not have duplicated data.

  **Required:** yes

- `Connection`: The connection string of the NATS server. When set, this will
  override the `NATS_SERVER_URL` environment variable. **We strongly recommend
  the use of the `NATS_SERVER_URL` environment variable to avoid connection
  strings in your code.**

  **Required:** no

### Environment variables

Additional details must be passed to the Apache Kafka driver. They will be loaded
from the environment variables.

- `NATS_SERVER_URL`: The NATS server URL to dial to leverage Pub / Sub. If
  `Options.PubSub.Connection` is set, it will override and be used in replacement
  of the existing environment variable.

  **Type:** `string`

  **Required:** yes (if `Options.PubSub.Connection` is not set)

  **Example:** `nats://127.0.0.1:4222`

  **Order:** options, environment variable

## Trigger configuration

Using the trigger mode `source.ModeSubscription`, a trigger can extract events from
NATS:
```go
func (t MyTrigger) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeSubscription,
    UsingSubscription: &source.Subscription{
      Topic:        "my-subject",
      Subscription: "my-queue",
    },
  }
}

```

The trigger will receive in realtime every events of a NATS topic in a given queue.
Each event can then be transformed and execute actions or flows to load data to
destinations.

### Subscription options

- `Topic`: The subject to subscribe to for receiving messages.

  **Required:** yes

- `Subscription`: The queue group used for receiving messages.

  **Required:** yes
