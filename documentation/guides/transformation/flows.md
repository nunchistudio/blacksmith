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

A flow can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate flow --name myflow

```

This will generate the recommended files for a flow, inside the working
directory.

If you prefer, you can generate a flow inside a directory with the `--path` flag:
```bash
$ blacksmith generate flow --name myflow \
  --path ./flows/myflow

```

## Call a flow from a trigger

Given a trigger (here is of mode HTTP), we can now call the flow from the `Payload`
returned by the `Extract` function like this:
```go
func (t MyTrigger) Extract(tk *source.Toolkit, req *http.Request) (*source.Payload, error) {

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
