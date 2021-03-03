---
title: HTTP middlewares
enterprise: false
---

# HTTP middlewares

The `gateway` and `scheduler` services can take a Go standard HTTP middleware.
This way, any HTTP logic can be added for authentication, authorization, or other
security requirements.

[As described in its dedicated section](/blacksmith/http/introduction/options),
the admin REST API can also take a stack of HTTP middlewares to secure its
endpoints.

A middleware is Go standard `http.Handler` wrapper. When configuring the application,
this looks like this:
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

    Gateway: &service.Options{
      Middleware: func(next http.Handler) http.Handler {

        return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
          res.Header().Set("Content-Type", "application/json")
          res.Header().Set("Access-Control-Allow-Origin", "*")

          fmt.Println("This happens before serving any HTTP endpoint")

          next.ServeHTTP(res, req)

          fmt.Println("This happens after serving any HTTP endpoint")
        })
      },
    },
  }

  return options
}

```
