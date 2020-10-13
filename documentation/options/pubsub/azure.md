---
title: Pub / Sub with Service Bus
enterprise: false
---

# Pub / Sub with Service Bus

The Azure pub / sub adapter allows to connect to Service Bus subscriptions and
therefore extract data from incoming messages.

The adapter is also used for realtime communication between the gateway and scheduler
services, [as described in the introduction](/blacksmith/introduction/what/overview).

## Application configuration

To use Azure as the pub / sub adapter for your application, you must set the `From`
key to `azure` in `*pubsub.Options`:
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
      From:         "azure",
      Topic:        "<topic>",
      Subscription: "<subscription>",
    },
  }

  return options
}
```

### Application options

- `Topic`: The Service Bus topic used by the gateway to forward jobs in realtime
  to the scheduler.

  **Required:** yes

- `Subscription`: The Service Bus subscription used by the scheduler to receive
  the jobs forwarded by the gateway.

  **Required:** yes

### Environment variables

Additional details must be passed to the Azure adapter. They will be loaded from
the environment variables, or from the `*pubsub.Options.Context` if not found.

- `SERVICEBUS_CONNECTION_STRING`: The Service Bus connection string to use. It can
  be the one from the parent namespace.
  
  **Type:** `string`

  **Required:** yes

  **Order:** environment variable, context

## Trigger configuration

Using the trigger mode `source.ModeSubscriber`, a trigger can extract events from
Azure Service Bus:
```go
func (t MyTrigger) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeSubscriber,
    UsingSubscriber: &source.Subscription{
      Topic:       "<topic>",
      Subscription: "<subscription>",
    },
  }
}
```

The trigger will receive in realtime every events of a queue registered in Service
Bus. Each event can then be transformed and loaded to destinations.

### Subscription options

- `Topic`: The Service Bus topic the subscription listens to.

  **Required:** yes

- `Subscription`: The Service Bus subscription to subscribe to for receiving
  messages.

  **Required:** yes
