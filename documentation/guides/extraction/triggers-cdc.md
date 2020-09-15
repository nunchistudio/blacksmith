---
title: CDC notifications
enterprise: false
---

# Triggers: CDC notifications

Triggers of mode `cdc` allow data extraction from Change-Data-Capture notifications.
The most interesting use case is for capturing changes from databases. This way,
whenever a condition is met in one of your databases, you can automatically listen
for the changes and act on it.

## Configuration

To make a source's trigger handles CDC notifications, you first need to specifiy the
mode to `cdc` in the trigger's options:
```go
func (t MyTrigger) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeCDC,
  }
}
```

## Usage

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

## Example

```go
/*
Extract runs in its own go routine. It is up to the function body
to include the forever loop. When set to this mode, the function
gives you access to channels to either return the payload or an
error whenever needed.
*/
func (t MyTrigger) Extract(tk *source.Toolkit, notifier *source.Notifier) {
  
  for {
    select {
    // case <-notification:
    //   notifier.Payload <- &source.Payload{}
    //   notifier.Error <- &errors.Error{}
    case <-notifier.IsShuttingDown:
      notifier.Done <- true
    }
  }
}
```
