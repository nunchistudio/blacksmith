---
title: Pub / Sub with AWS SNS / SQS
enterprise: false
---

# Pub / Sub with AWS SNS / SQS

The AWS SNS / SQS driver as the `pubsub` adapter allows to subscribe to SQS queues
and therefore extract data from incoming messages.

The adapter is also used for realtime communication between the gateway and scheduler
services, [as described in the onboarding](/blacksmith/start/onboarding/how).

## Application configuration

To use AWS as the Pub / Sub adapter for your application, you must set the `From`
key to `aws/snssqs` in `*pubsub.Options`:
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
      From:         pubsub.DriverAWSSNSSQS,
      Topic:        "arn:aws:sns:<region>:<id>:<topic>",
      Subscription: "arn:aws:sqs:<region>:<id>:<queue>",
    },
  }

  return options
}

```

### Application options

- `Topic`: The AWS ARN used by the gateway to forward jobs in realtime to the
  scheduler. It can either be a SNS or a SQS ARN depending on your needs.

  **Required:** yes

- `Subscription`: The AWS SQS ARN used by the scheduler to receive the jobs forwarded
  by the gateway.

  **Required:** yes

### Environment variables

Additional details must be passed to the AWS driver. They will be loaded from the
environment variables.

- `AWS_ACCESS_KEY_ID`: The AWS access key identifier to use.
  
  **Type:** `string`

  **Required:** yes

- `AWS_SECRET_ACCESS_KEY`: The AWS secret access key to use.
  
  **Type:** `string`

  **Required:** yes

- `AWS_REGION`: The AWS region to use.
  
  **Type:** `string`

  **Required:** yes

## Trigger configuration

Using the trigger mode `source.ModeSubscription`, a trigger can extract events from
AWS SQS:
```go
func (t MyTrigger) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeSubscription,
    UsingSubscription: &source.Subscription{
      Subscription: "arn:aws:sqs:<region>:<id>:<queue>",
    },
  }
}

```

The trigger will receive in realtime every events of a queue registered in AWS SQS.
Each event can then be transformed and execute actions or flows to load data to
destinations.

### Subscription options

- `Subscription`: The AWS SQS ARN to subscribe to for receiving messages.

  **Required:** yes
