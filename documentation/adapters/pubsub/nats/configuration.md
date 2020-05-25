#### Usage

```go
&blacksmith.Options{
  PubSub: &pubsub.Options{
    From:       "nats",
    Enabled:    true,
    Topic:      "blacksmith",
    Connection: "nats://127.0.0.1:4222",
  },
}
```

#### Environment variables

- `NATS_SERVER_URL`: The NATS server URL to dial to leverage pub / sub. If
  `Options.PubSub.Connection` is set, it will override and be used in replacement
  of this environment variable.

  Example: `nats://127.0.0.1:4222`
