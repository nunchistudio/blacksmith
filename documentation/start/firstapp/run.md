---
title: Running an instance
enterprise: false
---

# Running an instance

## Docker stack for development

A `Docker-compose.yml` is generated along your application. This file is not
required for running the application but is here for convenience to help you get
started faster in development. It contains:
- A NATS server for using the `nats` driver for the `pubsub` adapter.
- A PostgreSQL database for using the `postgres` driver for the `store`,
  `supervisor`, and `wanderer` adapters. This database is also used as a data
  warehouse alter in the guides. We use the same database for everything here
  aiming for simplicity, but it should not be the case when running in production.

You can customize the stack as much as you need, and then run it with:
```bash
$ docker compose up -d

```

## A few words about ports

Blacksmith leverages Docker for most of its operations, like running a service.
By default and for security reasons, Blacksmith does not bind the container ports.
So the `gateway` and `scheduler` services running in Docker are not accessible
from your local machine.

It is however possible to bind the desired ports as you might already do with the
Docker CLI. The `--bind` flag acts the same way as the Docker CLI does. The following
command binds the port `9090` of the container to the port `1234` of the host, and
the port `9091` to the port `1214`:
```bash
$ blacksmith start --bind 1213:9090 --bind 1214:9091

```

In production, it is highly recommended to run the gateway and scheduler on separate
*applications* for better security and scalability. This is made possible by passing
the desired service to run in the `--service` flag.

You can therefore run the `gateway` service with:
```bash
$ blacksmith start --service gateway -bind 1213:9090

```

And run the `scheduler` service with:
```bash
$ blacksmith start --service scheduler --bind 1214:9091

```

## Starting the services

To keep things as clear as possible, we'll assume you run the `gateway` on port
`9090` and the `scheduler` on port `9091`, with:
```bash
$ blacksmith start --bind 9090:9090 --bind 9091:9091

```

This starts both services and should output this:
```bash
INFO[2021-06-17T13:31:04+02:00] service/gateway: Server listening on `:9090`
INFO[2021-06-17T13:31:04+02:00] service/scheduler: Server listening on `:9091`

```

Now that the application is up and running, let's dive into some details to help
you discover step-by-step.
