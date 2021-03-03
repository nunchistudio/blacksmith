---
title: Configuration
enterprise: true
---

# Configuration

The admin API can be attached to at most one of the `gateway` or `scheduler`
service. By default, it is attached to the `scheduler` service.

The admin can take a Go standard HTTP middleware. It will be added to the stack
of middleware passed to the parent service's options. This way, any HTTP logic
can be added for authentication, authorization, or other security requirements.

Example:
```go
package main

import (
  "fmt"
  "http"

  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/service"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Scheduler: &service.Options{
      Address:  ":9091",
      Admin: &service.Admin{
        Enabled: true,
        Middleware: func(next http.Handler) http.Handler {

          return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
            res.Header().Set("Content-Type", "application/json")
            res.Header().Set("Access-Control-Allow-Origin", "*")

            fmt.Println("This happens before serving the admin endpoints")

            next.ServeHTTP(res, req)

            fmt.Println("This happens after serving the admin endpoints")
          })
        },
      },
    },
  }

  return options
}

```
