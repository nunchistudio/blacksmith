---
title: CRON schedules
enterprise: false
---

# Triggers: CRON schedules

Triggers of mode `cron` allow data extraction from scheduled tasks. It is useful
to extract data on recurring interval.

## Create a CRON trigger

A CRON trigger can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode cron

```

This will generate the recommended files for a CRON trigger, inside the working
directory.

If you prefer, you can generate the trigger inside a directory with the `--path`
flag:
```bash
$ blacksmith generate trigger --name mytrigger \
  --mode cron \
  --path ./sources/mysource

```

## Usage of a CRON trigger

If the trigger mode is `cron`, it must respect the interface
[`source.TriggerCRON`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#TriggerCRON).

The signature of the `Extract` function is:
```go
Extract(*source.Toolkit) (*source.Event, error)

```
