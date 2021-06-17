---
title: Sources
enterprise: false
---

# Sources

A source is a collection of triggers Extracting events' data from a same source.
For example, a database could be used as a source and register:
- CRON triggers for running recurring tasks;
- CDC triggers for listening for notifications.

As you can notice, a source can register triggers of different *modes*.

The triggers registered by a source return Extracted events' data along other
informations. These informations are then processed by the `gateway` service:

![Blacksmith Sources](/images/blacksmith/guides-etl.001.png)

## Create a source

A source must respect the interface
[`source.Source`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/source?tab=doc#Source).

The recommended way to create a source is by using the `generate` command, as
follow:
```bash
$ blacksmith generate source --name mysource

```

This will generate the recommended files for a source, inside the working
directory.

If you prefer, you can generate a source inside a directory with the `--path` flag:
```bash
$ blacksmith generate source --name mysource \
  --path ./sources/mysource

```

## Register a source

Once a source is created, it must be registered in the Blacksmith options before
being used.

You can register a source to a Blacksmith application as follow:
```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/source"

  "github.com/<org>/<app>/mysource"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Sources: []source.Source{
      mysource.New(&mysource.Options{
        // ...
      }),
    },
  }

  return options
}

```
