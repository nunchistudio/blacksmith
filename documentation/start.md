---
title: Starting the server
enterprise: false
---

# Starting the server

When starting the application, Blacksmith will automatically download, install,
and update the adapters you passed in the options. This will create a `.blacksmith`
directory at the top of your application, including the Go plugin for each adapter.

## In development

The best way to run the application is to run the `development` directory. It is
a `main` package, that loads the data pipeline previously created, and starts the
server.

It starts both the gateway and scheduler.
```bash
GO111MODULE=on go run ./development/main.go
```

## In production

In production, it is higlhy recommended to run the gateway and scheduler on separate
machines for better security and scalability.

Running the gateway:
```bash
GO111MODULE=on go run ./gateway/main.go
```

Running the scheduler:
```bash
GO111MODULE=on go run ./scheduler/main.go
```
