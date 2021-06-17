---
title: Pub / Sub messages
enterprise: false
---

# Triggers: Pub / Sub messages

Triggers of mode `sub` allow data Extraction from messages received in a Pub / Sub
subscription. This way, whenever a message is published on a given topic or for
a given subscription, it will automatically be received by the `sub` trigger.

This mode is only available if the `pubsub` adapter is configured for the application.

Available `pubsub` adapters:
- [AWS SNS / SQS](/blacksmith/options/pubsub/aws) (`aws/snssqs`)
- [Azure Service Bus](/blacksmith/options/pubsub/azure) (`azure/servicebus`)
- [Google Pub / Sub](/blacksmith/options/pubsub/google) (`google/pubsub`)
- [Apache Kafka](/blacksmith/options/pubsub/kafka) (`kafka`)
- [NATS](/blacksmith/options/pubsub/nats) (`nats`)
- [RabbitMQ](/blacksmith/options/pubsub/rabbitmq) (`rabbitmq`)

## Create a subscription trigger

A subscription trigger can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode sub

```

This will generate the recommended files for a subscription trigger, inside the
working directory.

If you prefer, you can generate the trigger inside a directory with the `--path`
flag:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode sub \
  --path ./sources/mysource

```

## Usage of a subscription trigger

If the trigger mode is `subscription`, it must respect the interface
[`source.TriggerSubscription`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/source?tab=doc#TriggerSubscription).

The signature of the `Extract` function is:
```go
Extract(*source.Toolkit, *pubsub.Message) (*source.Event, error)

```

Please refer to your Pub / Sub adapter configuration page for details about trigger
options.
