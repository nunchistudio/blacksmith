---
title: Pub / Sub with Service Bus
enterprise: false
---

# Pub / Sub with Service Bus

The Azure Service Bus driver as the `pubsub` adapter allows to subscribe to Service
Bus subscriptions and therefore extract data from incoming messages.

The adapter is also used for realtime communication between the gateway and scheduler
services, [as described in the introduction](/blacksmith/introduction/what/how).

## Application configuration

To use Azure as the Pub / Sub adapter for your application, you must set the `From`
key to `azure/servicebus` in `*pubsub.Options`:
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
      From:         pubsub.DriverAzureServiceBus,
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

Additional details must be passed to the Azure driver. They will be loaded from
the environment variables.

- `SERVICEBUS_CONNECTION_STRING`: The Service Bus connection string to use. It can
  be the one from the parent namespace.
  
  **Type:** `string`

  **Required:** yes

## Trigger configuration

Using the trigger mode `source.ModeSubscription`, a trigger can extract events from
Azure Service Bus:
```go
func (t MyTrigger) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeSubscription,
    UsingSubscription: &source.Subscription{
      Topic:        "<topic>",
      Subscription: "<subscription>",
    },
  }
}

```

The trigger will receive in realtime every events of a queue registered in Service
Bus. Each event can then be transformed and execute actions or flows to load data to
destinations.

### Subscription options

- `Topic`: The Service Bus topic the subscription listens to.

  **Required:** yes

- `Subscription`: The Service Bus subscription to subscribe to for receiving
  messages.

  **Required:** yes
