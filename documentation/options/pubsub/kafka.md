---
title: Pub / Sub with Apache Kafka
enterprise: false
---

# Pub / Sub with Apache Kafka

The Apache Kafka pub / sub adapter allows to subscribe to topics and therefore
extract data from incoming messages.

The adapter is also used for realtime communication between the gateway and scheduler
services, [as described in the introduction](/blacksmith/introduction/what/overview).

## Application configuration

To use Apache Kafka as the pub / sub adapter for your application, you must set
the `From` key to `kafka` in `*pubsub.Options`:
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
      From:         "kafka",
      Topic:        "blacksmith",
      Subscription: "blacksmith",
      Connection:   "127.0.0.1:9092,127.0.0.1:9093,127.0.0.1:9094",
    },
  }

  return options
}
```

### Application options

- `Topic`: The topic used by the gateway to forward jobs in realtime to the
  scheduler.

  **Required:** yes

- ` Subscription`: The consumer group used by the scheduler to receive the jobs
  forwarded by the gateway. We leverage Kafka consumer groups to benefit a single
  active consumer when deploying in distributed environments. This allows to receive
  messages only once, and therefore do not have duplicated data.

  **Required:** yes

- `Connection`: Comma-delimited list of hosts. When set, this will override the
  `KAFKA_BROKERS` environment variable. **We strongly recommend the use of the
  `KAFKA_BROKERS` environment variable to avoid connection strings in your code.**

  **Required:** no

### Environment variables

Additional details must be passed to the Apache Kafka adapter. They will be loaded
from the environment variables.

- `KAFKA_BROKERS`: The Kafka broker URLs to dial to leverage pub / sub. If
  `Options.PubSub.Connection` is set, it will override and be used in replacement
  of the existing environment variable.

  **Type:** `string`

  **Required:** yes (if `Options.PubSub.Connection` is not set)

  **Example:** `127.0.0.1:9092,127.0.0.1:9093,127.0.0.1:9094`

  **Order:** options, environment variable

## Trigger configuration

Using the trigger mode `source.ModeSubscriber`, a trigger can extract events from
Apache Kafka:
```go
func (t MyTrigger) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeSubscriber,
    UsingSubscriber: &source.Subscription{
      Topic:        "my-topic",
      Subscription: "my-consumer-group",
    },
  }
}
```

The trigger will receive in realtime every events of a Kafka consumer group for
a given topic. Each event can then be transformed and loaded to destinations.

### Subscription options

- `Topic`: The topic to subscribe to for receiving messages.

  **Required:** yes

- `Subscription`: The consumer group used for receiving messages.

  **Required:** yes
