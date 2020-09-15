---
title: Pub / Sub with NATS
enterprise: false
---

# Pub / Sub with NATS

The NATS pub / sub adapter allows to subscribe to topics and therefore extract data
from incoming messages.

The adapter is also used for realtime communication between the gateway and scheduler
services, [as described in the introduction](https://nunchi.studio/blacksmith/introduction/how).

## Application configuration

To use NATS as the pub / sub adapter for your application, you must set the `From`
key to `nats` in `*pubsub.Options`:
```go
package main

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    PubSub: &pubsub.Options{
      From:       "nats",
      Topic:      "blacksmith",
      Broker:     "blacksmith",
      Connection: "nats://127.0.0.1:4222",
    },
  }

  return options
}
```

### Application options

- `Topic`: The subject used by the gateway to forward jobs in realtime to the
  scheduler.

  **Required:** yes

- ` Broker`: The queue used by the scheduler to receive the jobs forwarded by the
  gateway. We leverage NATS queues to benefit a single active consumer when deploying
  in distributed environments. This allows to receive messages only once, and
  therefore do not have duplicated data.

  **Required:** yes

- `Connection`: The connection string of the NATS server. When set, this will
  override the `NATS_SERVER_URL` environment variable. **We strongly recommend
  the use of the `NATS_SERVER_URL` environment variable to avoid connection
  strings in your code.**

  **Required:** no

### Environment variables

Additional details must be passed to the Apache Kafka adapter. They will be loaded
from the environment variables.

- `NATS_SERVER_URL`: The NATS server URL to dial to leverage pub / sub. If
  `Options.PubSub.Connection` is set, it will override and be used in replacement
  of the existing environment variable.

  **Type:** `string`

  **Required:** yes (if `Options.PubSub.Connection` is not set)

  **Example:** `nats://127.0.0.1:4222`

  **Order:** options, environment variable

## Trigger configuration

Using the trigger mode `source.ModeSubscriber`, a trigger can extract events from
NATS:
```go
func (t MyTrigger) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeSubscriber,
    UsingSubscriber: &source.Subscription{
      Broker:       "my-queue",
      Subscription: "my-subject",
    },
  }
}
```

The trigger will receive in realtime every events of a NATS topic in a given queue.
Each event can then be transformed and loaded to destinations.

### Subscription options

- `Broker`: The queue group used for receiving messages.

  **Required:** yes

- `Subscription`: The subject to subscribe to for receiving messages.

  **Required:** yes
