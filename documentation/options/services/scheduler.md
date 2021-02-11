---
title: Scheduler
enterprise: false
---

# Scheduler

The scheduler is the service loading jobs from destinations' triggers.

## Options

- `Address`: HTTP address for the server to listen to.

  **Required:** no

  **Defaults:** `:9091`

- `KeyFile`: Path to the SSL / TLS key file.

  **Required:** no

- `CertFile`: Path to the SSL / TLS certificate file.

  **Required:** no

- `Middleware`: Go HTTP middleware to attach to the server. It will be applied to
  every routes exposed by the REST API in the Enterprise Edition.

  **Required:** no

- `Attach`: Go HTTP handler (`http.Handler`) to attach to the server.

  **Required:** no

## Example

```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/service"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Scheduler: &service.Options{
      Address:  ":9091",
      KeyFile:  "cert.key",
      CertFile: "cert.crt",
    },
  }

  return options
}

```
