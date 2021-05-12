---
title: Pub / Sub with Google Cloud
enterprise: false
---

# Pub / Sub with Google Cloud

The Google Cloud Pub / Sub driver as the `pubsub` adapter allows to subscribe to
Google Cloud subscriptions and therefore extract data from incoming messages.

The adapter is also used for realtime communication between the gateway and scheduler
services, [as described in the introduction](/blacksmith/introduction/what/how).

## Application configuration

To use Google Cloud as the Pub / Sub adapter for your application, you must set
the `From` key to `google/pubsub` in `*pubsub.Options`:
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
      From:         pubsub.DriverGooglePubSub,
      Topic:        "projects/<project>/topics/<topic>",
      Subscription: "projects/<project>/subscriptions/<subscription>",
    },
  }

  return options
}

```

### Application options

- `Topic`: The Google Cloud topic used by the gateway to forward jobs in realtime
  to the scheduler.

  **Required:** yes

- `Subscription`: The Google Cloud subscription used by the scheduler to receive
  the jobs forwarded by the gateway.

  **Required:** yes

### Environment variables

Additional details must be passed to the Google Cloud driver. They will be loaded
from the environment variables.

- `GOOGLE_APPLICATION_CREDENTIALS`: The file path of the JSON file that contains
  your service account key.
  
  **Type:** `string`

  **Required:** yes

## Trigger configuration

Using the trigger mode `source.ModeSubscription`, a trigger can extract events from
Google Pub / Sub:
```go
func (t MyTrigger) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeSubscription,
    UsingSubscription: &source.Subscription{
      Subscription: "<subscription>",
    },
  }
}

```

The trigger will receive in realtime every events of a subscription registered in
Google Pub / Sub. Each event can then be transformed and execute actions or flows to load data to
destinations.

### Subscription options

- `Subscription`: The Google Cloud subscription to subscribe to for receiving
  messages.

  **Required:** yes
