---
title: Actions
enterprise: false
---

# Actions

An action is in charge of loading some data to a destination.

## Create an action

An action is an interface of type
[`destination.Action`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/destination?tab=doc#Action).

## Example

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
MyAction is the payload structure received by an
action and that will be sent by the scheduler. Blacksmith
needs 'Context', 'Data', and 'SentAt' keys to ensure
consistency across events.

The 'Version' key allows schema versioning, following
best practices for production. More details in the
next guide.
*/
type MyAction struct {
  Version string     `json:"version,omitempty"`
  Context *MyContext `json:"context"`
  Data    *User      `json:"data"`
  SentAt  *time.Time `json:"sent_at"`
}

/*
String returns the string representation of the action.
*/
func (a MyAction) String() string {
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
func (a MyAction) Schedule() *destination.Schedule {
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
func (a MyAction) Marshal(tk *destination.Toolkit) (*destination.Payload, error) {

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
Load is the function being run when the event is loaded into
the destination by the scheduler. It is in charge of the "L"
in the ETL process: it Loads the data to the destination.

In this case, since there is an error, the job will fail and
will be marked as "discarded" after 3 retries. Otherwise, the
scheduler will consider the job has "succeeded".
*/
func (a MyAction) Load(tk *destination.Toolkit, queue *store.Queue, then chan<- destination.Then) {

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

## Register an action

Once an action is created, it must be registered in the parent destination triggers
before being used.

You can add an action to a destination as follow:
```go
func (md *MyDestination) Actions() map[string]destination.Action {
  return map[string]destination.Action{
    "my-action": MyAction{},
  }
}
```

## Add an action within a flow

Given the flow we created before, we can now add the action within it:
```go
func (mf *MyFlow) Transform(tk *flow.Toolkit) destination.Actions {
  return map[string][]destination.Action{
    "my-destination": []destination.Action{
      &mydestination.MyAction{
        Version: "2020-10-01",
        Data: &mydestination.User{
          FullName: mf.FullName,
          Email:    mf.Email,
        },
      },
    },
  }
}
```

Every time the flow is executed, a *job* will be created for the action.

The `Version` key introduced in this last example is optional. It leverages schema
versioning, following production best practices as explained in the next guide.
