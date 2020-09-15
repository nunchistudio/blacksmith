---
title: Sources
enterprise: false
---

# Sources

A source is a collection of triggers emitted from a same source. For example, a
database could be used as a source and register:
- CRON triggers for running recurring tasks;
- CDC triggers for listening for notifications.

## Create a source

A source is an interface of type
[`source.Source`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#Source).

## Example

```go
package mysource

import (
  "github.com/nunchistudio/blacksmith/flow/source"
)

/*
MySource implements the source.Source interface.
*/
type MySource struct {
  options *source.Options
}

/*
New returns a valid Blacksmith source.
*/
func New() source.Source {
  return &Source{
    options: &source.Options{},
  }
}

/*
String returns the string representation of the source.
*/
func (s *MySource) String() string {
  return "my-source"
}

/*
Options returns common source options. They will be shared
across every triggers of this source, except when overridden.
*/
func (s *MySource) Options() *source.Options {
  return s.options
}

/*
Triggers return a list of triggers the source is able to
handle. Since we do not have a trigger yet, return an empty
map for now.

Their respective Extract function will automatically be
triggered by the gateway given their Mode.
*/
func (s *MySource) Triggers() map[string]source.Trigger {
  return map[string]source.Trigger{

    // ...

  }
}
```

## Register a source

Once a source is created, it must be registered in the Blacksmith options before
being used.

You can add a source as follow:
```go
package main

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Sources: []*source.Options{
      {
        Load: mysource.New(),
      },
    },
  }

  return options
}
```
