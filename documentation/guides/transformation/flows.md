---
title: Flows
enterprise: false
---

# Flows

In the previous guides we created triggers within a source. We associate sources'
triggers to destinations' actions with flows. This *middleman* is useful because
it allows to:
- Enable or disable flows whenever we want;
- Share flows between multiple triggers;
- Share data structures between triggers and actions;
- Add business logic specific to a flow without impacting the triggers and actions.

## Create a flow

A flow is of type
[`flow.Flow`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow?tab=doc#Flow).

## Example

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
Transform is the function called by the scheduler and
allows to run — independently from one another — a
collection of actions. Since we do not have a destination
yet, return an empty map for now.
*/
func (f *MyFlow) Transform(tk *flow.Toolkit) destination.Actions {
  return map[string][]destination.Action{

    // ...

  }
}
```

## Call a flow from a trigger

Given the HTTP action we created before, we can now call the flow from the `Payload`
returned by the `Extract` function like this:
```go
func (mt MyTrigger) Extract(tk *source.Toolkit, req *http.Request) (*source.Payload, error) {

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
