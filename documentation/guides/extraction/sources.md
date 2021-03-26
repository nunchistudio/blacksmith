---
title: Sources
enterprise: false
---

# Sources

A source is a collection of triggers emitted from a same source. For example, a
database could be used as a source and register:
- CRON triggers for running recurring tasks;
- CDC triggers for watching for notifications.

## Create a source

A source is an interface of type
[`source.Source`](https://pkg.go.dev/github.com/nunchistudio/blacksmith/flow/source?tab=doc#Source).

A source can be generated with the `generate` command, as follow:
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

If you need to [handle data migrations](/blacksmith/practices/management/migrations)
within the source, you can also add the `--migrations` flag:
```bash
$ blacksmith generate source --name mysource \
  --path ./sources/mysource \
  --migrations

```

## Register a source

Once a source is created, it must be registered in the Blacksmith options before
being used.

You can add a source as follow:
```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/flow/source"

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
