---
title: Docker environment
enterprise: false
---

# Docker environment

Blacksmith leverages Docker for environment parity. When running most of the
commands, the Blacksmith CLI will run itself inside a Docker container, created
with the `Dockerfile` located in the root directory of application.

## Custom Docker image

When generating an application, a `Dockerfile` is created as well. This allows to
make some customization, but more importantly to add the root directory of the
application to the container. Building and running a Blacksmith application outside
an official Docker image is not supported.

One should looks like this:
```dockerfile
FROM nunchistudio/blacksmith-enterprise:0.17.1-alpine

ADD ./ /app
WORKDIR /app

RUN rm -rf go.sum
RUN go mod tidy

EXPOSE 9090 9091

```

## Docker stack for development

A `Docker-compose.yml` is also generated along your application. This file is not
required for running the application but is here for convenience to help you get
started even faster in development. It contains:
- a PostgreSQL database for using the `postgres` driver for the `store`,
  `supervisor`, and `wanderer` adapters;
- a NATS server for using the `nats` driver for the `pubsub` adapter.

You can customize the stack as much as you need, and run it with:
```bash
$ docker-compose up -d

```
