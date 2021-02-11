---
title: Running an instance
enterprise: false
---

# Running an instance

## Building an application

You can validate and build an application without starting any service by running:
```bash
$ blacksmith build

```

This will ensure your application can be successfully loaded by the Blacksmith CLI
and pushed to production.

## Starting a service

Before starting an instance of a service, Blacksmith will automatically validate
and build your application. If the build process does not succeed, your application
can not start.

Based on what we learned in the previous guide, we can start the Docker stack with:
```bash
$ blacksmith start

```

In production, it is highly recommended to run the gateway and scheduler on separate
machines for better security and scalability. This is made possible by passing the
desired service to run to the `--service` flag.

You can therefore run the `gateway` service with:
```bash
$ blacksmith start --service gateway

```

And run the `scheduler` service with:
```bash
$ blacksmith start --service scheduler

```

## Docker container

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

Also, the Docker cache can be disabled when building your application. To do so,
simply add the `--no-cache` flag:
```bash
$ blacksmith start --bind 1213:9090 --bind 1214:9091 --no-cache

```
