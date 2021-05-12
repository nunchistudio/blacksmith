---
title: Actions
enterprise: false
---

# Actions

An action is in charge of loading some data to a destination. When an action is
processed by a flow or directly by a trigger, this creates a *job*.

## Create an action

An action is an interface of type
[`destination.Action`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/destination?tab=doc#Action).

An action can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate action --name myaction

```

This will generate the recommended files for an action, inside the working
directory.

If you prefer, you can generate an action inside a directory with the `--path` flag:
```bash
$ blacksmith generate action --name myaction \
  --path ./destinations/mydestination

```

## Register an action

Once an action is created, it must be registered in its parent destination before
being used.

You can register an action to a destination as follow:
```go
func (d *MyDestination) Actions() map[string]destination.Action {
  return map[string]destination.Action{
    "my-action": MyAction{},
  }
}

```

This allows to pass some destiation's options down to the action if necessary.

## Add an action within a flow

Given the flow created before, we can now add the action within it:
```go
func (f *MyFlow) Transform(tk *flow.Toolkit) destination.Actions {
  return map[string][]destination.Action{
    "my-destination": []destination.Action{
      &mydestination.MyAction{
        Version: "2020-10-01",
        Data: &mydestination.User{
          FullName: f.FullName,
          Email:    f.Email,
        },
      },
    },
  }
}

```

Every time the flow is executed, a *job* will be created for the action.

The `Version` key introduced in this last example is optional. It leverages schema
versioning, following production best practices as explained in 
[one of the "Advanced practices" guides](/blacksmith/practices/management/versioning).
