#### Usage

```go
ctx := context.Background()
ctx = context.WithValue(ctx, "CONSUL_DATACENTER", "eu-west-1")
ctx = context.WithValue(ctx, "CONSUL_HTTP_TOKEN", "my-private-token")

// ...

&blacksmith.Options{
  Supervisor: &supervisor.Options{
    From:    "consul",
    Enabled: true,
    Context: ctx,
    Join: &supervisor.Node{
      Address: "127.0.0.1:8500",
    }
  },
}
```

#### Environment variables

- `CONSUL_ADDRESS`: The Consul agent URL to dial to leverage distributed locks.
  If `Options.Supervisor.Join.Address` is set, it will override and be used in
  replacement of this environment variable.

  Example: `127.0.0.1:8500`


#### Additional context

- `CONSUL_DATACENTER`: The Consul datacenter to use.

  Required: no

  Type: `string`

  Defaults: `dc1`

- `CONSUL_SCHEME`: The Consul scheme to use.

  Required: no

  Type: `string`

  Defaults: `http`

- `CONSUL_HTTP_TOKEN`: The Consul token to use. It must have the read / write
  permissions for keys using the `blacksmith/` prefix.

  Required: no

  Type: `string`

- `CONSUL_NAMESPACE`: The Consul namespace to use.

  Required: no

  Type: `string`
