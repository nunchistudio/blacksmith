---
title: Triggers
enterprise: false
---

# Triggers

A trigger allows data extraction from a source. A trigger can be *triggered* when:
- a HTTP route is called;
- a CRON schedule is met;
- a CDC notification arrived;
- a message is published in a Pub / Sub mechanism.

## Create a trigger

Please refer to the proper guide for each trigger mode.

## Register a trigger

Once a trigger is created, it must be registered in the parent source triggers
before being used.

You can add a trigger to a source as follow:
```go
func (s *MySource) Triggers() map[string]source.Trigger {
  return map[string]source.Trigger{
    "my-trigger": MyTrigger{},
  }
}

```

## Notes about triggers

### Payload

When a trigger is executed, it can return a payload (in `Payload` of type
[`*source.Payload`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/destination?tab=doc#Payload)).
The payload contains:
- `Context` is a dictionary of information that provides useful context about an
  event. The context should be used inside every triggers for consistency.
  It must be a valid JSON since it will be used using `encoding/json` `Marshal`
  and `Unmarshal` functions.
- `Data` is the byte representation of the data sent by the event. It must be a
  valid JSON since it will be used using `encoding/json` `Marshal` and `Unmarshal`
  functions.
- `Flows` defines the flows of actions to run when this trigger is executed. They
  will only be executed if no error has been returned (`nil`).
- `SentAt` allows you to keep track of the timestamp when the event was originally
  sent.

### Error handling

When a trigger is executed, it can return any kind of error, as long as it respects
the builtin `error` type. However, we strongly recommend to use the
[`errors.Error`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/helper/errors?tab=doc)
type for better consistency between your application and the Blacksmith platform.

If `error` is not `nil`, the gateway considers the event as untrusted and will not
continue. Therefore, no jobs will be created.
