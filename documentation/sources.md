---
title: Sources and triggers
enterprise: false
---

# Sources and triggers

A source is a collection of triggers emitted from a same source. For example, a
database could be used as a source and register:
- CRON triggers for running recurring tasks;
- CDC triggers for listening for notifications.

## Sources

### Create a source

A source is of type [`source.Source`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#Source).

Example:
```go
package mysource

import (
  "github.com/nunchistudio/blacksmith/flow/source"
)

/*
MySource implements the source.Source interface.
*/
type MySource struct {
  options *source.Options
}

/*
New returns a valid Blacksmith source.
*/
func New() source.Source {
  return &Source{
    options: &source.Options{},
  }
}

/*
String returns the string representation of the source.
*/
func (s *MySource) String() string {
  return "my-source"
}

/*
Options returns common source options. They will be shared across every
triggers of this source, except when overridden.
*/
func (s *MySource) Options() *source.Options {
  return s.options
}

/*
Triggers return a list of triggers the source is able to handle. Since
we do not have a trigger yet, return an empty map for now.
*/
func (s *MySource) Triggers() map[string]source.Trigger {
  return map[string]source.Trigger{}
}
```

### Register a source

Once a source is created, it must be registered in the Blacksmith options before
being used.

You can add a source as follow:
```go
package main

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Sources: []*source.Options{
      {
        Load: mySource.New(),
      },
    },
  }

  return options
}
```

## Triggers

### Create a trigger

A trigger is of type [`source.Trigger`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#Trigger).

A trigger is one of HTTP, CRON, or CDC. A source can have multiple triggers of
different modes.

Example:
```go
package mysource

import (
  "time"

  "github.com/nunchistudio/blacksmith/flow"
  "github.com/nunchistudio/blacksmith/flow/source"
)

/*
User holds information about a user for this source.
*/
type User struct {
  Username  string `json:"username"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  Email     string `json:"email"`
}

/*
TriggerIdentify is the payload structure sent by an event and that
will be received by the gateway. Blacksmith needs 'Context', 'Data',
and 'SentAt' keys to ensure consistency across incoming events.
*/
type TriggerIdentify struct {
  Context *MyContext `json:"context"`
  Data    *User      `json:"data"`
  SentAt  *time.Time `json:"sent_at"`
}

/*
String returns the string representation of the trigger.
*/
func (t TriggerIdentify) String() string {
  return "identify"
}

/*
Mode indicates the mode of the trigger. Usage for each mode
are detailed below.
*/
func (t TriggerIdentify) Mode() *source.Mode {
  return &source.Mode{}
}
```

#### HTTP triggers

To make a source's trigger handles HTTP request, you first need to specifiy the
mode to `http` in the trigger's options. The following route will trigger the
`Marshal()` function every time a HTTP request comes in `/identify` using a `POST`
method:
```go
func (t TriggerIdentify) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeHTTP,
    UsingHTTP: &source.Route{
      Methods: []string{"POST"},
      Path:    "/identify",
    },
  }
}
```

If the trigger is of kind HTTP, it must respect the interface
[`source.TriggerHTTP`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#TriggerHTTP)
by having a `Marshal` function as follow:
```go
Marshal(*source.Toolkit, *http.Request) (*source.Payload, error)
```

Example:
```go
/*
Marshal is the function being executed when the HTTP route is
called. The function gives access to the original HTTP request.

The gateway ensures the content type returned is "application/json".
It also includes information inside the response body such as
the jobs created by the flows called.
*/
func (t TriggerIdentify) Marshal(tk *source.Toolkit, req *http.Request) (*source.Payload, error) {

  // Create an empty payload, catch unwanted fields, and
  // unmarshal it. Return an error if any occured.
  var payload TriggerIdentify
  decoder := json.NewDecoder(req.Body)
  decoder.DisallowUnknownFields()
  err := decoder.Decode(&payload)
  if err != nil {
    return nil, err
  }

  // Try to marshal the context from the request payload.
  ctx, err := json.Marshal(&payload.Context)
  if err != nil {
    return nil, err
  }

  // Try to marshal the data from the request payload.
  data, err := json.Marshal(&payload.Data)
  if err != nil {
    return nil, err
  }

  // Return the context, data, and a collection of flows to execute
  // against destinations with transformed data. More details about
  // data flowing in the "Flow automation" guide.
  return &source.Payload{
    Context: ctx,
    Data:    data,
    Flows:   []flow.Flow{},
  }, nil
}
```

#### CRON triggers

To make a source's trigger handles CRON schedules, you first need to specifiy the
mode to `cron` in the trigger's options. The following schedule will trigger the
`Marshal()` function every minute:
```go
func (t TriggerIdentify) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeCRON,
    UsingCRON: &source.Schedule{
      Interval: "@every 1m",
    },
  }
}
```

If the trigger is of kind CRON, it must respect the interface
[`source.TriggerCRON`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#TriggerCRON)
by having a `Marshal` function as follow:
```go
Marshal(*source.Toolkit) (*source.Payload, error)
```

Example:
```go
/*
Marshal will be executed every minute given the interval set in
the scheduling options.
*/
func (t TriggerIdentify) Marshal(tk *source.Toolkit) (*source.Payload, error) {

  // ...

  // Return a marshalled context and data, along a collection of
  // flows to execute against destinations with transformed data
  // More details about data flowing in the "Flow automation" guide.
  return &source.Payload{
    Context: ctx,
    Data:    data,
    Flows:   []flow.Flow{},
  }, nil
}
```

#### CDC triggers

To make a source's trigger handles CDC notifications, you first need to specifiy the
mode to `cdc` in the trigger's options.
```go
func (t TriggerIdentify) Mode() *source.Mode {
  return &source.Mode{
    Mode: source.ModeCDC,
  }
}
```

If the trigger is of kind CDC, it must respect the interface
[`source.TriggerCDC`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#TriggerCDC)
by having a `Marshal` function as follow:
```go
Marshal(*source.Toolkit, *source.Notifier)
```

Example:
```go
/*
Marshal runs in its own go routine. It is up to the function body
to include the forever loop. When set to this mode, the function
gives you access to channels to either return the payload or an
error whenever needed.

Also, since this mode is asynchronous, there is no way for the
gateway to know when the trigger is done. To gracefully shutdown
like in a synchronous mode, the function receives a message on
`IsShuttingDown` and must write to `Done` whenever the function is
ready to exit. Otherwise, the gateway will block until `true` is
received on `Done`.
*/
func (t TriggerIdentify) Marshal(tk *source.Toolkit, notifier *source.Notifier) {
  
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

### Register a trigger

Once a trigger is created, it must be registered in the parent source triggers
before being used.

You can add a trigger to a source as follow:
```go
func (m *MySource) Triggers() map[string]source.Trigger {
  return map[string]source.Trigger{
    "identify": TriggerIdentify{},
  }
}
```
