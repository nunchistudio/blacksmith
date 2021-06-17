---
title: Triggers
enterprise: false
---

# Triggers

A trigger allows data Extraction from a source's event. An event can be triggered
when:
- a HTTP route is requested (mode `http`);
- a CRON schedule is met (mode `cron`);
- a CDC notification arrived (mode `cdc`);
- a message is received in a Pub / Sub subscription (mode `sub`).

## Create a trigger

Please refer to the proper guide for each trigger *mode*.

## Register a trigger

Once a trigger is created, regardless its mode, it must be registered in its parent
source before being used.

You can register a trigger to a source as follow:
```go
func (s *MySource) Triggers() map[string]source.Trigger {
  return map[string]source.Trigger{
    "my-trigger": MyTrigger{},
  }
}

```

This allows to pass some source's options down to the trigger if necessary.

## Notes about triggers

### Event created

When a trigger is executed, it returns the created event(s) (in `Event` of type
[`*source.Event`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/source?tab=doc#Event)).
The event contains:
- `Context` is a dictionary of information that provides useful context about an
  event. The context should be used inside every triggers for consistency.
  It must be a valid JSON since it will be used using `encoding/json` `Marshal`
  and `Unmarshal` functions.
- `Data` is the byte representation of the data sent by the event. It must be a
  valid JSON since it will be used using `encoding/json` `Marshal` and `Unmarshal`
  functions.
- `Flows` defines the flows of actions to run when this trigger is executed. They
  will only be executed if no error has been returned (`nil`).
- `Actions` allows to create jobs directly from the event when the trigger is
  executed. Where `Flows` is used to create jobs from a common data schema, calling
	`Actions` directly not from a flow can often be simpler.
- `SubEvents` is a collection of events that need to be created and associated
	to the event being processed. This is useful for triggers accepting a batch of
	events in a single request.
- `SentAt` allows you to keep track of the timestamp when the event was originally
  sent.

### Error handling

When a trigger is executed, it can return any kind of error, as long as it respects
the builtin `error` type. However, we strongly recommend to use the
[`errors.Error`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/helper/errors?tab=doc)
type for better consistency between your application and the Blacksmith platform.

If `error` is not `nil`, the gateway considers the event as untrusted and will not
continue. Therefore, no jobs will be created.
