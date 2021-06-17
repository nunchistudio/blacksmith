---
title: Actions
enterprise: false
---

# Actions

An action is in charge of Loading some data to a destination. When an action is
processed by a flow or directly by a trigger, this creates a *job*.

## Add an action within a flow

Given the flow created before, we can now add the action within it. In this example,
we call the action `RunOperation` from the destination `sqlike(warehouse)` which
is based on the `sqlike` module and its `sqlikedestination` package:
```go
func (f *MyFlow) Transform(tk *flow.Toolkit) destination.Actions {
  return map[string][]destination.Action{
    "sqlike(warehouse)": []destination.Action{
      sqlikedestination.RunOperation{
        Filename: "./warehouse/operations/insert-user.sql",
        Data: map[string]interface{}{
          "user": map[string]interface{}{
            "first_name": &f.FirstName,
            "last_name":  &f.LastName,
            "username":   &f.Username,
            "email":      &f.Email,
            "created_at": time.Now().UTC(),
          },
        },
      },
    },
  }
}

```

Since this destination is a SQL database, you can Transform and Load data with
SQL, [as described in the dedicated guides](/blacksmith/tlt). The `user` will be
accessible like this:
```sql
INSERT INTO users (id, first_name, last_name, username, email, created_at)
  VALUES (
    '{% ksuid %}',
    '{{ user.first_name | capfirst }}',
    '{{ user.last_name | upper }}',
    '{{ user.username | slugify }}',
    '{{ user.email | lower }}',
    '{{ user.created_at }}'
  );

```

Every time the flow is executed, a *job* will be created for the action.
