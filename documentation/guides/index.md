---
title: Guides & Tutorials
enterprise: false
---

# Guides & Tutorials

Following is a cheatsheet of the methods used across sources, flows, and destinations
for ETL operations.

Please refer to each section on the left navigation for details and examples.

## Data Extraction

### From HTTP requests

```go
func (t MyTrigger) Extract(tk *source.Toolkit, req *http.Request) (*source.Payload, error) {

  // ...
  
  return &source.Payload{
    Context: ctx,
    Data:    data,
    Flows:   []flow.Flow{},
  }, nil
}
```

### From CRON schedules

```go
func (t MyTrigger) Extract(tk *source.Toolkit) (*source.Payload, error) {

  // ...

  return &source.Payload{
    Context: ctx,
    Data:    data,
    Flows:   []flow.Flow{},
  }, nil
}
```

### From CDC notifications

```go
func (t MyTrigger) Extract(tk *source.Toolkit, notifier *source.Notifier) {
  
  for {
    select {
    case <-notification:

      // ...

      notifier.Payload <- &source.Payload{
        Context: ctx,
        Data:    data,
        Flows:   []flow.Flow{},
      }
    
    case <-notifier.IsShuttingDown:
      notifier.Done <- true
    }
  }
}
```

### From Pub / Sub messages

```go
func (t MyTrigger) Extract(tk *source.Toolkit, msg *pubsub.Message) (*source.Payload, error) {

  // ...
  
  return &source.Payload{
    Context: ctx,
    Data:    data,
    Flows:   []flow.Flow{},
  }, nil
}
```

## Data Transformation

```go
func (f *MyFlow) Transform(tk *flow.Toolkit) destination.Actions {
  return map[string][]destination.Action{

    // ...

  }
}
```

## Data Load

### Step 1: Marshal the data

```go
func (a MyAction) Marshal(tk *destination.Toolkit) (*destination.Payload, error) {

  // ...

  return &destination.Payload{
    Context: ctx,
    Data:    data,
  }, nil
}
```

### Step 2: Load the data

```go
func (a MyAction) Run(tk *destination.Toolkit, queue *store.Queue, then chan<- destination.Then) {

  for _, event := range queue.Events {
    for _, job := range event.Jobs {

      // ...

      then <- destination.Then{
        Jobs:         []string{job.ID},
        Error:        nil,
        ForceDiscard: false
        OnSucceeded:  []destination.Action{},
        OnFailed:     []destination.Action{},
        OnDiscarded:  []destination.Action{},
      }
    }
  }
}
```
