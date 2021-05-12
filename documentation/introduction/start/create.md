---
title: Creating an application
enterprise: false
---

# Creating an application

> As of now, we assume you are familiar with Go and Docker, and already have
  installed these technologies on your machine along the Blacksmith CLI.

## Generate a new application

The best way to create a new Blacksmith application is by using the `generate`
command of the CLI. The following command generates all the required files in the
current directory:
```bash
$ blacksmith generate application --name myapp

```

If you prefer, you can generate a new application inside a directory with the
`--path` flag:
```bash
$ blacksmith generate application --name myapp --path ./myapp

```

The directory will be created if it does not exist yet.

Now that the application has been created, let's dive into the details to help
you understand some mechanisms.

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

## Go modules

Blacksmith leverages Go modules for managing dependencies. Before continuing, make
sure you have `go.mod` with the required dependencies within it:
```go
module github.com/<org>/<app>

go 1.16

require github.com/nunchistudio/blacksmith v0.16.0

replace golang.org/x/net => golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20210415045647-66c3f260301c

replace golang.org/x/sync => golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9

```

Validate and lock the dependencies by executing the command:
```bash
$ go mod tidy

```
