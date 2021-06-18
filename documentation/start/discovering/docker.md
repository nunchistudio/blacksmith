---
title: Docker environment
enterprise: false
---

# Docker environment

Blacksmith leverages Docker for environment parity. When running most of the
commands, the Blacksmith CLI will run itself inside a Docker container. It uses
the `Dockerfile` located at the root directory of the application, created when
generating the application.

This allows to make some customization, but more importantly to add the root
directory of the application to the container. Building and running a Blacksmith
application outside an official Docker image is not supported.

One should look like this:
```dockerfile
FROM nunchistudio/blacksmith-enterprise:0.18.0-alpine

ADD ./ /app
WORKDIR /app

RUN rm -rf go.sum
RUN go mod tidy

EXPOSE 9090 9091

```
