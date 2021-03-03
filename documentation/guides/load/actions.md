---
title: Actions
enterprise: false
---

# Actions

An action is in charge of loading some data to a destination.

## Create an action

> If the action you wish to generate relies on a *starter*, please refer to the
  appropriate guide on the left navigation for details and examples.

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

If you need to [handle data migrations](/blacksmith/practices/management/migrations)
within the action, you can also add the `--migrations` flag:
```bash
$ blacksmith generate action --name myaction \
  --path ./destinations/mydestination \
  --migrations

```

## Register an action

Once an action is created, it must be registered in the parent destination triggers
before being used.

You can add an action to a destination as follow:
```go
func (d *MyDestination) Actions() map[string]destination.Action {
  return map[string]destination.Action{
    "my-action": MyAction{},
  }
}

```

## Add an action within a flow

Given the flow we created before, we can now add the action within it:
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
[one of the "Best practices" guides](/blacksmith/practices/management/versioning).
