---
title: Pub / Sub messages
enterprise: false
---

# Triggers: Pub / Sub messages

Triggers of mode `subscription` allow data extraction from messages received in a
Pub / Sub mechanism. This way, whenever a message is published on a given topic
or for a given subscription, it will automatically be received by the `subscription`
trigger.

This mode is only available if the `pubsub` adapter is configured for the application.

Available `pubsub` adapters:
- [AWS SNS / SQS](/blacksmith/options/pubsub/aws) (`aws/snssqs`)
- [Azure Service Bus](/blacksmith/options/pubsub/azure) (`azure/servicebus`)
- [Google Pub / Sub](/blacksmith/options/pubsub/google) (`google/pubsub`)
- [Kafka](/blacksmith/options/pubsub/kafka) (`kafka`)
- [NATS](/blacksmith/options/pubsub/nats) (`nats`)
- [RabbitMQ](/blacksmith/options/pubsub/rabbitmq) (`rabbitmq`)

## Create a subscription trigger

A subscription trigger can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode sub

```

This will generate the recommended files for a subscription trigger, inside the working
directory.

If you prefer, you can generate the trigger inside a directory with the `--path`
flag:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode sub \
  --path ./sources/mysource

```

If you need to [handle data migrations](/blacksmith/guides/practices/migrations)
within the trigger, you can also add the `--migrations` flag:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode sub \
  --path ./sources/mysource \
  --migrations

```

## Usage of a subscription trigger

If the trigger mode is `subscription`, it must respect the interface
[`source.TriggerSubscription`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#TriggerSubscription).

The signature of the `Extract` function is:
```go
Extract(*source.Toolkit, *pubsub.Message) (*source.Payload, error)

```

Please refer to your Pub / Sub adapter configuration page for details about trigger
options. [Go to configuration reference.](/blacksmith/options)
