---
title: Consul supervisor
enterprise: true
---

# Consul supervisor

The use of Consul from a Blacksmith application allows to automatically register
and deregister the `gateway` and `scheduler` services from the registry:

![Blacksmith with Consul](/images/blacksmith/consul.001.png)

It also automatically handle sessions and locks across nodes to avoid access
collision when working within a multi-node environment:

![Blacksmith with Consul](/images/blacksmith/consul.002.png)

## Options

- `Connection`: The Consul agent URL to dial to leverage distributed locks. When
  set, this will override the `CONSUL_ADDRESS` environment variable. **We strongly
  recommend the use of the `CONSUL_ADDRESS` environment variable to avoid connection
  strings in your code.**

## Environment variables

Some options can be loaded from the environment variables.

- `CONSUL_ADDRESS`: The Consul agent URL to dial to leverage distributed locks.
  If `Options.Supervisor.Connection` is set, it will override and be used in
  replacement of this environment variable.

  **Type:** `string`

  **Required:** yes (if `Options.Supervisor.Connection` is not set)

  **Example:** `http://127.0.0.1:8500`

  **Order:** environment variable, options

- `CONSUL_DATACENTER`: The Consul datacenter to use.
  
  **Type:** `string`

  **Required:** no

  **Defaults:** `dc1`

- `CONSUL_SCHEME`: The Consul scheme to use.

  **Type:** `string`
  
  **Required:** no

  **Defaults:** `http`

- `CONSUL_HTTP_TOKEN`: The Consul token to use. It must have the read / write
  permissions for keys using the `blacksmith/` prefix.

  **Type:** `string`
  
  **Required:** no

- `CONSUL_NAMESPACE`: The Consul namespace to use.

  **Type:** `string`
  
  **Required:** no

## Example

```go
package main

import (
  "github.com/nunchistudio/blacksmith"
  "github.com/nunchistudio/blacksmith/adapter/supervisor"
)

func Init() *blacksmith.Options {

  var options = &blacksmith.Options{

    // ...

    Supervisor: &supervisor.Options{
      From: supervisor.DriverConsul,
    },
  }

  return options
}

```
