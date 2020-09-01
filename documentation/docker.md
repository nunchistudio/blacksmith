---
title: Docker environment
enterprise: false
---

# Docker environment

## Custom Docker image

As specified before, an application can only run inside a supported Docker image.

Before running the environment, we can create a custom `Dockerfile` at the root
directory of your application based on one of the image. This allows to make some
customization, but more importantly to add the root directory of the application
to the container. Without this, the Blacksmith CLI will not find the directory
containing the `Init()` function required.

Example:
```dockerfile
FROM nunchistudio/blacksmith-enterprise:0.11.0-alpine

ADD ./ /app

WORKDIR /app

RUN go mod tidy
```

## In development

The simplest way to get started quickly in development is to start both the gateway
and scheduler in the same stack. Assuming you have a `Dockerfile` like we just
wrote, we can create a `Docker-compose.yml` file to build our image and add it
to a stack of services, such as a PostgreSQL database (for the store) and a NATS
server (for the pub / sub):
```yml
version: "3"

services:
  blacksmith_gateway:
    build: "./"
    restart: "unless-stopped"
    entrypoint: ["blacksmith", "start", "--service", "gateway"]
    environment:
      NATS_SERVER_URL: "nats://pubsub:4222"
      POSTGRES_STORE_URL: "postgres://app:qwerty@datastore:5432/app?sslmode=disable"
      POSTGRES_WANDERER_URL: "postgres://app:qwerty@datastore:5432/app?sslmode=disable"
    ports:
      - "8080:8080"
    depends_on:
      - "datastore"
      - "pubsub"

  blacksmith_scheduler:
    build: "./"
    restart: "unless-stopped"
    entrypoint: ["blacksmith", "start", "--service", "scheduler"]
    environment:
      NATS_SERVER_URL: "nats://pubsub:4222"
      POSTGRES_STORE_URL: "postgres://app:qwerty@datastore:5432/app?sslmode=disable"
      POSTGRES_WANDERER_URL: "postgres://app:qwerty@datastore:5432/app?sslmode=disable"
    ports:
      - "8081:8081"
    depends_on:
      - "datastore"
      - "pubsub"

  datastore:
    image: "postgres:12-alpine"
    restart: "unless-stopped"
    environment:
      POSTGRES_DB: "app"
      POSTGRES_USER: "app"
      POSTGRES_PASSWORD: "qwerty"
    volumes:
      # - "./migrations:/docker-entrypoint-initdb.d"
      - "app:/var/lib/postgresql/data"
    ports:
      - "5432:5432"

  pubsub:
    image: "nats:2-alpine"
    restart: "unless-stopped"
    ports:
      - "4222:4222"
      - "8222:8222"

volumes:
  app:
```

The complete example including the PostgreSQL migrations lives in our demo project
[smithy](https://github.com/nunchistudio/smithy).

## In production

In production, it is highly recommended to run the gateway and scheduler on separate
machines for better security and scalability. Using the same image, a `Docker-compose.yml`
only running the gateway will look like this:
```yml
version: "3"

services:
  blacksmith_gateway:
    build: "./"
    restart: "unless-stopped"
    entrypoint: ["blacksmith", "start", "--service", "gateway"]
    environment:
      NATS_SERVER_URL: ""
      POSTGRES_STORE_URL: ""
      POSTGRES_WANDERER_URL: ""
    ports:
      - "8080:8080"
```
