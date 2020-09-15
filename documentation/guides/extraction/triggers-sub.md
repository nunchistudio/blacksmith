---
title: Pub / Sub messages
enterprise: false
---

# Triggers: Pub / Sub messages

Triggers of mode `subscriber` allow data extraction from messages received in a
Pub / Sub mechanism. This way, whenever a message is published on a given topic
or for a given subscription, it will automatically be received by the `subscriber`.

This mode is only available if the Pub / Sub adapter is configured for the application.

## Configuration

Please refer to your Pub / Sub adapter configuration page for details about trigger
options. [Go to configuration reference.](http://localhost:2000/blacksmith/options)

## Usage

If the trigger mode is `subscriber`, it must respect the interface
[`source.TriggerSubscriber`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#TriggerSubscriber).

The signature of the `Extract` function is:
```go
Extract(*source.Toolkit, *pubsub.Message) (*source.Payload, error)
```

## Example

```go
/*
Extract is the function being executed when a new message is
received. The function gives access to the message body as well
as its metadata.
*/
func (t MyTrigger) Extract(tk *source.Toolkit, msg *pubsub.Message) (*source.Payload, error) {

  // Try to unmarshal the data from the message.
  var m MyTrigger
  json.Unmarshal(msg.Body, &m)
  
  // You now have access to the data.
  tk.Logger.Info("New message received:", m)

  // Return the context, data, and a collection of flows to execute
  // against destinations with transformed data. More details about
  // data flowing in the "Flows" guide.
  return &source.Payload{
    Context: ctx,
    Data:    data,
    Flows:   []flow.Flow{},
  }, nil
}
```
