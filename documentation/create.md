---
title: Creating a project
enterprise: false
---

# Creating a project

> As of now, we assume you have a better understanding of how Blacksmith works
  and you are familiar with the Go language.

## Template repository

The simplest way to create a new project is to download and unzip the
[`smithy` template repository from GitHub](https://github.com/nunchistudio/smithy).
This is the project we start from to demo Blacksmith features and benefits.

It comes with:
- A `development` package for starting the gateway and scheduler together in the
  same process.
- A `gateway` and `scheduler` packages for starting the gateway and scheduler
  in different processes (suited for production).
- A `builder` package allowing the Blacksmith CLI (Enterprise version only) to
  compile the application as a Go plugin and therefore being able to manage
  migrations across adapters.
- A `Docker-compose.yml` setup for running the gateway and the scheduler using
  the [Blacksmith Docker image](https://github.com/nunchistudio/blacksmith-docker).
  It also runs a PostgreSQL store and a NATS pub / sub.
- A `Vagrantfile` for running the stack in a virtual machine using the
  [Blacksmith Vagrant box](https://github.com/nunchistudio/blacksmith-vagrant).
- Sample adapters for `source` and `destination` interfaces showing how the data
  pipeline works.

## Packages documentation

The following guides will drive you how to read, configure, run, and maintain a
Blacksmith application, based on the template and the 
[documentation of the Go packages](https://pkg.go.dev/github.com/nunchistudio/blacksmith?tab=doc).

## New pipeline

The first file to look at is `load.go`. It is in charge of loading a new Blacksmith
data pipeline from a set of configuration:

```go
p, _ := blacksmith.New(&blacksmith.Options{})
```

To run correctly, an application needs an adapter for the `gateway`, `scheduler`
and `store` interfaces.

When not set, the `standard` adapters will be use for the gateway and scheduler.
They are the default implementations and the ones we use at Nunchi. The store
is required and we will use the `postgres` store adapter for the demo:

```go
p, _ := blacksmith.New(&blacksmith.Options{
  Gateway: &gateway.Options{},
  Scheduler: &scheduler.Options{},
  Store: &store.Options{
    From: "postgres",
  },
})
```

It may seem light, but the above configuration is the only thing needed to setup
your data pipeline.

To enable realtime with pub / sub features, simply add the `pubsub` options with
the adapter you need:

```go
p, _ := blacksmith.New(&blacksmith.Options{
  Gateway: &gateway.Options{},
  Scheduler: &scheduler.Options{},
  Store: &store.Options{
    From: "postgres",
  },
  PubSub: &pubsub.Options{
    From:    "nats",
    Enabled: true,
  },
})
```

## Sources and destinations

Sources and destinations also need to be added directly to your options as follow:

```go
p, _ := blacksmith.New(&blacksmith.Options{
  Gateway: &gateway.Options{},
  Scheduler: &scheduler.Options{},
  Store: &store.Options{
    From: "postgres",
  },
  PubSub: &pubsub.Options{
    From:    "nats",
    Enabled: true,
  },
  Sources: []*source.Options{},
  Destinations: []*destination.Options{},
})
```
