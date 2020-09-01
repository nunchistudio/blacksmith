---
title: Flow automation
enterprise: false
---

# Flow automation

In the previous guides we created triggers within a source, and an action within
a destination. We associate triggers to actions with flows. This *middleman* is
useful because it allows to:
- Enable or disable flows whenever we want;
- Share flows between multiple triggers;
- Share data structures between triggers and actions;
- Add business logic specific to a flow without impacting the triggers and actions.

## Create a flow

A flow is of type
[`flow.Flow`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow?tab=doc#Flow).

Example:
```go
package myflow

import (
  "github.com/nunchistudio/blacksmith/flow"
  "github.com/nunchistudio/blacksmith/flow/destination"
)

/*
MyFlow implements the flow.Flow interface.
*/
type MyFlow struct {
  Username  string `json:"username"`
  FullName  string `json:"full_name"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  Email     string `json:"email"`
}

/*
Options returns common flow options. When the flow
is disabled, the scheduler will not go through it
and therefore no jobs related to this flow will
be created.
*/
func (f *MyFlow) Options() *flow.Options {
  return &flow.Options{
    Enabled: true,
  }
}

/*
Run is the function called by the scheduler and allows
to run — independently from one another — a collection
of actions.
*/
func (f *MyFlow) Run(tk *flow.Toolkit) destination.Actions {
  return map[string][]destination.Action{
    "my-destination": []destination.Action{
      &mydestination.ActionRegister{
        Data: &mydestination.User{
          FullName: f.FullName,
          Email:    f.Email,
        },
      },
    },
  }
}
```

## Call a flow from a trigger
Given the HTTP action we created before, we can call the flow from the `Marshal`
function like this:
```go
func (t TriggerIdentify) Marshal(tk *source.Toolkit, req *http.Request) (*source.Payload, error) {

  // ...

  return &source.Payload{
    Context: ctx,
    Data:    data,
    Flows:   []flow.Flow{
      &myflow.MyFlow{
        FullName: &payload.Data.FullName,
        Email:    &payload.Data.Email,
      },
    },
  }, nil
}
```
