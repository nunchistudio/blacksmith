---
title: CDC notifications
enterprise: false
---

# Triggers: CDC notifications

Triggers of mode `cdc` allow data extraction from Change-Data-Capture notifications.
The most interesting use case is for capturing changes from databases. This way,
whenever a condition is met in one of your databases, you can automatically listen
for the changes and act on it.

## Create a CDC trigger

A CDC trigger can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode cdc

```

This will generate the recommended files for a CDC trigger, inside the working
directory.

If you prefer, you can generate the trigger inside a directory with the `--path`
flag:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode cdc \
  --path ./sources/mysource

```

If you need to [handle data migrations](/blacksmith/guides/practices/migrations)
within the trigger, you can also add the `--migrations` flag:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode cdc \
  --path ./sources/mysource \
  --migrations

```

## Usage of a CDC trigger

If the trigger mode is `cdc`, it must respect the interface
[`source.TriggerCDC`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#TriggerCDC).

The signature of the `Extract` function is:
```go
Extract(*source.Toolkit, *source.Notifier)

```

Since this mode is asynchronous, there is no way for the gateway to know when the
trigger is done. To gracefully shutdown like in other trigger modes, the function
receives a message on `notifier.IsShuttingDown` and must write to `notifier.Done`
whenever the function is ready to exit. Otherwise, the gateway will block until
`true` is received on `notifier.Done`.
