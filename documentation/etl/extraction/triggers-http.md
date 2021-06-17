---
title: HTTP requests
enterprise: false
---

# Triggers: HTTP requests

Triggers of mode `http` allow data Extraction from a source when an API route is
requested. The most interesting use case is for capturing webhooks from third-party
services. This way, whenever a condition is met in one of your applications, it
can automatically make a HTTP request to one of the triggers registered in `http`
mode.

## Create a HTTP trigger

A HTTP trigger can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode http

```

This will generate the recommended files for a HTTP trigger, inside the working
directory.

If you prefer, you can generate the trigger inside a directory with the `--path`
flag:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode http \
  --path ./sources/mysource

```

## Usage of a HTTP trigger

If the trigger mode is `http`, it must respect the interface
[`source.TriggerHTTP`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/source?tab=doc#TriggerHTTP).

The signature of the `Extract` function is:
```go
Extract(*source.Toolkit, *http.Request) (*source.Event, error)

```

The gateway ensures the content type returned is `application/json`. It can also
includes information inside the response body such as the jobs created by the
flows called.
