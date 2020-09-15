---
title: Creating a project
enterprise: false
---

# Creating a project

> As of now, we assume you are familiar with Go and Docker, and already have
  installed these technologies on your machine.

## The `Init` function

A Blacksmith application is compiled and run as a Go plugin by the CLI.

Go plugins act like Go applications but can be loaded by external programs, like
the Blacksmith CLI does. Go plugins are `main` packages without `init` and
`main` functions.

The CLI needs to validate, load, and run an application. To achieve this, it must
have a `main` package including the following function signature:
```go
func Init() *blacksmith.Options
```

Blacksmith options is of type
[`*blacksmith.Options`](https://pkg.go.dev/github.com/nunchistudio/blacksmith?tab=doc#Options).
It allows to configure the different components needed by the platform to successfully
run an application.

So, the entrypoint of an application looks like this:
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

## Go modules

Blacksmith leverages Go modules for managing dependencies. Before continuing, make
sure you have `go.mod` with the required dependencies within it:
```go
module github.com/<org>/<app>

go 1.15

require github.com/nunchistudio/blacksmith v0.11.0

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd
```

Validate and lock the dependencies by executing the command:
```bash
$ go mod tidy
```

## File tree

Inside a dedicated directory, you should now have:
- `application.go` for initializing the application (the name does not matter);
- `go.mod` and `go.sum` for managing Go dependencies.
