---
title: Gateway
enterprise: false
---

# Gateway

The gateway is the service extracting every events from sources' triggers.

## Options

- `Address`: HTTP address for the server to listen to.

  **Required:** no

  **Defaults:** `:9090`

- `KeyFile`: Path to the SSL / TLS key file.

  **Required:** no

- `CertFile`: Path to the SSL / TLS certificate file.

  **Required:** no

- `Middleware`: Go HTTP middleware to attach to the server. It will be applied to
  every routes registered in the sources' triggers using the `http` mode.

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

    Gateway: &service.Options{
      Address:  ":9090",
      KeyFile:  "cert.key",
      CertFile: "cert.crt",
    },
  }

  return options
}
```
