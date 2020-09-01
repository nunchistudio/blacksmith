---
title: Destinations and actions
enterprise: false
---

# Destinations and actions

A destination is a collection of actions that load data to a same destination.
For example, a database could be used as a data warehouse to centrally store all
the data of an organization. It would have multiple actions according to the data
to load.

## Destinations

### Create a destination

A destination is of type
[`destination.Destination`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/destination?tab=doc#Destination).

Example:
```go
package mydestination

import (
  "github.com/nunchistudio/blacksmith/flow/destination"
)

/*
MyDestination implements the destination.Destination interface.
*/
type MyDestination struct {
  options *destination.Options
}

/*
New returns a valid Blacksmith destination.

For the purpose of the guide, we will use a destination that
does not support realtime loading, but with little interval. In
other words, we do not want events to be loaded into this
destination in realtime. This can be overridden by each action
if necessary.

In case of failure, we specify to retry every 2 minutes with a
limit of 20 retries. After that, if the jobs still fail they
will be marked as 'discarded'.
*/
func New() destination.Destination {
  return &MyDestination{
    options: &destination.Options{
      DefaultSchedule: &destination.Schedule{
        Realtime:   false,
        Interval:   "@every 2m",
        MaxRetries: 20,
      },
    },
  }
}

/*
String returns the string representation of the destination.
*/
func (d *MyDestination) String() string {
  return "my-destination"
}

/*
Options returns common destination options. They will be shared
across every actions of this destination, except when overridden.
*/
func (d *MyDestination) Options() *destination.Options {
  return d.options
}

/*
Actions return a list of actions the destination is able to
handle. Since we do not have an action yet, return an empty map
for now.
*/
func (d *MyDestination) Actions() map[string]destination.Action {
  return map[string]destination.Action{}
}
```

### Register a destination

Once a destination is created, it must be registered in the Blacksmith options before
being used.

You can add a destination as follow:
```go
package main

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Destinations: []*destination.Options{
      {
        Load: myDestination.New(),
      },
    },
  }

  return options
}
```

## Actions

### Create an action

An action is of type
[`destination.Action`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/destination?tab=doc#Action).

Example:
```go
package mydestination

import (
  "encoding/json"
  "time"

  "github.com/nunchistudio/blacksmith/adapter/store"
  "github.com/nunchistudio/blacksmith/flow/destination"
)

/*
User holds information about a user for this destination.
*/
type User struct {
  FullName string `json:"full_name"`
  Email    string `json:"email"`
}

/*
ActionRegister is the payload structure received by an
action and that will be sent by the scheduler. Blacksmith
needs 'Context', 'Data', and 'SentAt' keys to ensure
consistency across events.
*/
type ActionRegister struct {
  Context *MyContext `json:"context"`
  Data    *User      `json:"data"`
  SentAt  *time.Time `json:"sent_at"`
}

/*
String returns the string representation of the action.
*/
func (a ActionRegister) String() string {
  return "register"
}

/*
Schedule allows destinations' actions to override the
schedule options of their related destination.

For the purpose of the demo, we override the original
schedule of the destination for a specific one. The data
will be loaded in realtime with a retry every minute in
case of failure with a maximum of 3 retries before being
discarded.
*/
func (a ActionRegister) Schedule() *destination.Schedule {
  return &destination.Schedule{
    Realtime:   true,
    Interval:   "@every 1m",
    MaxRetries: 3,
  }
}

/*
Marshal is the function being run to transform the data
received. Like for a source's trigger, it is also in charge
of the "T" in the ETL process: it can Transform (if needed)
the payload to the given data structure.
*/
func (a ActionRegister) Marshal(tk *destination.Toolkit) (*destination.Payload, error) {

  // Try to marshal the action data passed directly to the
  // struct.
  buff, err := json.Marshal(&a.Data)
  if err != nil {
    return nil, err
  }

  // Create a payload with the data. Since the 'Context' key
  // is not set, the one from the event will automatically be
  // applied.
  p := &destination.Payload{
    Data: buff,
  }

  // Return the payload with the marshaled data.
  return p, nil
}

/*
Run is the function being run when the event is loaded into
the destination by the scheduler. It is in charge of the "L"
in the ETL process: it Loads the data to the destination.

In this case, since there is an error, the job will fail and
will be marked as "discarded" after 3 retries. Otherwise, the
scheduler will consider the job has succeeded.
*/
func (a ActionRegister) Run(tk *destination.Toolkit, queue *store.Queue, then chan<- destination.Then) {

  // Go through each event of the queue. The queue only contains
  // the jobs needed to be run against this destination's action.
  for _, event := range queue.Events {
    for _, job := range event.Jobs {
      var u User
      json.Unmarshal(job.Data, &u)

      // ...

      // Once our business logic to load the data to the
      // destination is done we can inform the scheduler
      // about the job status. If Jobs is nil or empty, all
      // jobs from the queue will be affected by the result.
      // This allows to either load the data entry-per-entry
      // or in batch if the destination allows it. If a job ID
      // is not returned, the scheduler will not be aware of
      // its status and will mark it as 'unknown'.
      //
      // OnSucceeded, OnFailed, and OnDiscarded allows to run
      // actions from the same destination given the job status.
      then <- destination.Then{
        Jobs: []string{job.ID},
        Error: &errors.Error{
          StatusCode: 400,
          Message:    "Internal Server Error",
        },
        OnSucceeded: []destination.Action{},
        OnFailed:    []destination.Action{},
        OnDiscarded: []destination.Action{},
      }
    }
  }
}
```

### Register an action

Once an action is created, it must be registered in the parent destination triggers
before being used.

You can add an action to a destination as follow:
```go
func (md *myDestination) Actions() map[string]destination.Action {
  return map[string]destination.Action{
    "identify": ActionIdentify{},
  }
}
```
