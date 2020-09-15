---
title: Destinations
enterprise: false
---

# Destinations

A destination is a collection of actions that load data to a same destination.
For example, a database could be used as a data warehouse to centrally store all
the data of an organization. It would have multiple actions according to the data
to load.

## Create a destination

A destination is an interface of type
[`destination.Destination`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/destination?tab=doc#Destination).

## Example

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

For the purpose of the guide, we will use a destination
that does not support realtime loading, but with little
interval. In other words, we do not want events to be
loaded into this destination in realtime. This can be
overridden by each action if necessary.

In case of failure, we specify to retry every 2 minutes
with a limit of 20 retries. After that, if the jobs still
fail they will be marked as "discarded".
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
String returns the string representation of the
destination.
*/
func (d *MyDestination) String() string {
  return "my-destination"
}

/*
Options returns common destination options. They will be
shared across every actions of this destination, except
when overridden.
*/
func (d *MyDestination) Options() *destination.Options {
  return d.options
}

/*
Actions return a list of actions the destination is able
to handle. Since we do not have an action yet, return an
empty map for now.
*/
func (d *MyDestination) Actions() map[string]destination.Action {
  return map[string]destination.Action{

    // ...

  }
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
        Load: mydestination.New(),
      },
    },
  }

  return options
}
```
