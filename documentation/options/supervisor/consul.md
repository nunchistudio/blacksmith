---
title: Consul supervisor
enterprise: true
---

# Consul supervisor

The use of Consul from a Blacksmith application allows to:
- Automatically register / deregister the gateway and scheduler services from the
  registry;
- Automatically handle sessions and locks across nodes to avoid access collision
  when working within a multi-node environment.

## Options

- `Join`: The node of the distributed system to join. Each running instance of
  Blacksmith must join a different Consul agent to avoid access collision.
  - `Address`: The Consul agent URL to dial to leverage distributed locks. When
  set, this will override the `CONSUL_ADDRESS` environment variable. **We strongly
  recommend the use of the `CONSUL_ADDRESS` environment variable to avoid
  connection strings in your code.**
  - `Tags`: Slice of tags related to the node.
  - `Meta`: Collection of meta-data related to the node.

## Environment variables

Some options can be loaded from the environment variables. They will be loaded
from `*pubsub.Options.Context` (or from `*pubsub.Options.Connection` for the
connection string) if not found.

- `CONSUL_ADDRESS`: The Consul agent URL to dial to leverage distributed locks.
  If `Options.Supervisor.Join.Address` is set, it will override and be used in
  replacement of this environment variable. Each running instance of Blacksmith
  must join a different Consul agent to avoid access collision.

  **Type:** `string`

  **Required:** yes (if `Options.Supervisor.Join.Address` is not set)

  **Example:** `127.0.0.1:8500`

  **Order:** environment variable, options

- `CONSUL_DATACENTER`: The Consul datacenter to use.
  
  **Type:** `string`

  **Required:** no

  **Defaults:** `dc1`

  **Order:** environment variable, context

- `CONSUL_SCHEME`: The Consul scheme to use.

  **Type:** `string`
  
  **Required:** no

  **Defaults:** `http`

  **Order:** environment variable, context

- `CONSUL_HTTP_TOKEN`: The Consul token to use. It must have the read / write
  permissions for keys using the `blacksmith/` prefix.

  **Type:** `string`
  
  **Required:** no

  **Order:** environment variable, context

- `CONSUL_NAMESPACE`: The Consul namespace to use.

  **Type:** `string`
  
  **Required:** no

  **Order:** environment variable, context

## Example

```go

package main

func Init() *blacksmith.Options {

  ctx := context.Background()
  ctx = context.WithValue(ctx, "CONSUL_DATACENTER", "eu-west-1")
  ctx = context.WithValue(ctx, "CONSUL_HTTP_TOKEN", "my-private-token")

  var options = &blacksmith.Options{

    // ...

    Supervisor: &supervisor.Options{
      From:    "consul",
      Context: ctx,
      Join: &supervisor.Node{
        Address: "127.0.0.1:8500",
      },
    },
  }

  return options
}
```
