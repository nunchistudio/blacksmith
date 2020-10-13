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

A destination can be generated with the `generate` command, as follow:
```bash
$ blacksmith generate destination --name mydestination
```

This will generate the recommended files for a destination, inside the working
directory.

If you prefer, you can generate a destination inside a directory with the `--path`
flag:
```bash
$ blacksmith generate destination --name mydestination --path ./destinations/mydestination
```

If you need to handle data migrations within the destination, you can also add the
`--migrations` flag:
```bash
$ blacksmith generate destination --name mydestination --path ./destinations/mydestination --migrations
```

### Register a destination

Once a destination is created, it must be registered in the Blacksmith options before
being used.

You can add a destination as follow:
```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/flow/destination"

  "github.com/<org>/<app>/mydestination"
)

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
