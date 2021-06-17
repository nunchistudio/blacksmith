---
title: Flows
enterprise: false
---

# Flows

In the previous guides we created triggers within a source. When returning the
processed event, a trigger can either call actions directly for Transformation or
pass through *flows*.

![Blacksmith Flows](/images/blacksmith/guides-etl.002.png)

The use of flows is optional. A trigger can call actions directly and never return
flows. However, flows can be useful because they allow to:
- Enable or disable flows whenever we want;
- Share actions between multiple triggers;
- Share data structures between multiple triggers and actions;
- Add business logic specific to a flow without impacting the triggers and actions.

## Create a flow

A flow is of type
[`flow.Flow`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow?tab=doc#Flow).

A flow can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate flow --name myflow

```

This will generate the recommended files for a flow, inside the working
directory.

If you prefer, you can generate a flow inside a directory with the `--path` flag:
```bash
$ blacksmith generate flow --name myflow \
  --path ./flows

```

## Call a flow from a trigger

Given a trigger (here is of mode `http`), we can now call the flow from the `Event`
returned by the `Extract` function like this:
```go
func (t MyTrigger) Extract(tk *source.Toolkit, req *http.Request) (*source.Event, error) {

  // ...

  return &source.Event{
    Context: ctx,
    Data:    data,
    Flows:   []flow.Flow{
      &flows.MyFlow{
        FirstName: &payload.Data.FirstName,
        LastName:  &payload.Data.LastName,
        Username:  &payload.Data.Username,
        Email:     &payload.Data.Email,
      },
    },
  }, nil
}

```
