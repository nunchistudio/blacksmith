---
title: Destinations
enterprise: false
---

# Destinations

A destination is a collection of actions that load data to a same destination.
For example, a database could be used as a data warehouse to centrally store all
the data of an organization. It would have multiple actions according to the data
to load.

The actions registered by a destination create jobs for loading each event's data
into the said destination. The `scheduler` service is in charge of executing these
jobs:

![Blacksmith Destinations](/images/blacksmith/how.006.png)

There is an infinity of possibilities and workarounds for loading data to
destinations. Instead of locking users into a few limited patterns and still not
covering every needs you might have, we offer a collection of [production-ready
modules](/blacksmith/modules). It simplifies development, enforces best practices,
and still allows a complete freedom on how data is loaded.

The use of these modules is optional. The interfaces exposed by the Blacksmith
Go API let you define your own destinations and actions if needed.

## Create a destination

A destination must respect the interface
[`destination.Destination`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/destination?tab=doc#Destination).

The recommended way to create a destination is by using the `generate` command,
as follow:
```bash
$ blacksmith generate destination --name mydestination

```

This will generate the recommended files for a destination, inside the working
directory.

If you prefer, you can generate a destination inside a directory with the `--path`
flag:
```bash
$ blacksmith generate destination --name mydestination \
  --path ./destinations/mydestination

```

## Register a destination

Once a destination is created, it must be registered in the Blacksmith options
before being used.

You can register a destination to a Blacksmith application as follow:
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

    Destinations: []destination.Destination{
      mydestination.New(&mydestination.Options{
        // ...
      }),
    },
  }

  return options
}

```
