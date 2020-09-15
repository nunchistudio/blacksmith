---
title: CRON schedules
enterprise: false
---

# Triggers: CRON schedules

Triggers of mode `cron` allow data extraction from scheduled tasks. It is useful
to extract data on recurring interval.

## Configuration

To make a source's trigger handles CRON schedules, you first need to specifiy the
mode to `cron` in the trigger's options.

The following schedule will trigger the `Extract` function every minute:
```go
func (t MyTrigger) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeCRON,
    UsingCRON: &source.Schedule{
      Interval: "@every 1m",
    },
  }
}
```

## Usage

If the trigger mode is `cron`, it must respect the interface
[`source.TriggerCRON`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#TriggerCRON).

The signature of the `Extract` function is:
```go
Extract(*source.Toolkit) (*source.Payload, error)
```

## Example

```go
/*
Extract will be executed every minute given the interval set in
the scheduling options.
*/
func (t MyTrigger) Extract(tk *source.Toolkit) (*source.Payload, error) {

  // ...

  // Return a marshalled context and data, along a collection of
  // flows to execute against destinations with transformed data
  // More details about data flowing in the "Flows" guide.
  return &source.Payload{
    Context: ctx,
    Data:    data,
    Flows:   []flow.Flow{},
  }, nil
}
```
