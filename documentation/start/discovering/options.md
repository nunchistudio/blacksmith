---
title: Application's options
enterprise: false
---

# Application's options

## The `Init` function

A Blacksmith application is compiled and run as a Go plugin by the CLI.

Go plugins act like Go applications but can be loaded by external programs. The
CLI compiles your Go code as a Go plugin. Go plugins are `main` packages without
a `main` function.

To achieve this, the `main` package of your application must have a function with
the following signature:
```go
func Init() *blacksmith.Options

```

Blacksmith options is of type
[`*blacksmith.Options`](https://pkg.go.dev/github.com/nunchistudio/blacksmith?tab=doc#Options).
It allows to configure the different components needed by the platform to successfully
run an application.

So, the entrypoint of an application shall look like this:
```go
package main

import (
  "github.com/nunchistudio/blacksmith"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

  }

  return options
}

```

## From development to production

When generating an application, all the options are already set to work in a
development environment. 

Please refer to [the configuration reference](/blacksmith/options) to properly
configure your application for a non-local environment, depending on your needs.
This should only takes a few minutes.
