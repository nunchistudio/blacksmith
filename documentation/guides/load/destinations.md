---
title: Destinations
enterprise: false
---

# Destinations

A destination is a collection of actions that load data to a same destination.
For example, a database could be used as a data warehouse to centrally store all
the data of an organization. It would have multiple actions according to the data
to load.

There is an infinity of possibilities and workarounds for loading data to
destinations. Instead of locking users into a few limited patterns and still not
covering every needs you may have, we offer a few *starters*. It simplifies
development, enforces best practices, and still allows a complete freedom on how
data is loaded.

The use of starters is optional.

## Create a destination

> Once the destination is created, please refer to the appropriate guide on the
  left navigation for details and examples for each *starter*.

A destination is an interface of type
[`destination.Destination`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/destination?tab=doc#Destination).

A destination can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate destination --name mydestination

```

This will generate the recommended files for a destination, inside the working
directory.

If you prefer, you can generate a destination inside a directory with the `--path`
flag:
```bash
$ blacksmith generate destination --name mydestination \
  --path ./destinations/mydestination

```

If you need to [handle data migrations](/blacksmith/guides/practices/migrations)
within the destination, you can also add the `--migrations` flag:
```bash
$ blacksmith generate destination --name mydestination \
  --path ./destinations/mydestination \
  --migrations

```

### Starter for HTTP APIs

The starter `net/http` generates a destination using the Go standard library for
communicating with HTTP APIs:
```bash
$ blacksmith generate destination --name mydestination \
  --starter net/http

```

### Starter for SQL databases

The starter `database/sql` generates a destination using the Go standard library
for leveraging any SQL driver:
```bash
$ blacksmith generate destination --name mydestination 
  --starter database/sql

```

This allows to leverage any SQL and SQL-like databases such as PostgreSQL, MySQL,
Snowflake, ClickHouse, etc.

### Starter for NoSQL databases

The starter `gocloud/docstore` generates a destination using [Go Cloud](https://gocloud.dev/)
for leveraging supported document stores:
```bash
$ blacksmith generate destination --name mydestination \
  --starter gocloud/docstore \
  --driver <driver>

```

Available drivers:
- `aws/dynamodb` for using Amazon DynamoDB.
- `azure/cosmosdb` for using MongoDB on Azure with CosmosDB.
- `google/firestore` for using Google Firestore.
- `mongodb` for using MongoDB.

### Starter for blob storages

The starter `gocloud/blob` generates a destination using [Go Cloud](https://gocloud.dev/)
for leveraging supported blob storages:
```bash
$ blacksmith generate destination --name mydestination \
  --starter gocloud/blob \
  --driver <driver>

```

Available drivers:
- `aws/s3` for using Amazon S3.
- `azure/blob` for using Azure Blob Storage.
- `google/storage` for using Google Cloud Storage.
- `file` for using local file storage.

### Starter for Pub / Sub

The starter `gocloud/pubsub` generates a destination using [Go Cloud](https://gocloud.dev/)
for leveraging supported Pub / Sub technologies:
```bash
$ blacksmith generate destination --name mydestination \
  --starter gocloud/pubsub \
  --driver <driver>

```

This allows to publish messages on topics for the supported drivers.

Available drivers:
- `aws/snssqs` for using AWS SNS / SQS.
- `azure/servicebus` for using Azure Service Bus.
- `google/pubsub` for using Google Pub / Sub.
- `kafka` for using Apache Kafka.
- `nats` for using NATS.
- `rabbitmq` for using RabbitMQ.

## Register a destination

Once a destination is created, it must be registered in the Blacksmith options before
being used.

You can add a destination as follow:
```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/flow/destination"

  "github.com/<org>/<app>/mydestination"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Destinations: []*destination.Options{
      {
        Load: mydestination.New(),
      },
    },
  }

  return options
}

```
