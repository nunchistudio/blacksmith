#### Usage

```go
ctx := context.Background()
ctx = context.WithValue(ctx, "RABBIT_PRODUCER_EXCHANGE", "producer-exchange-name")

// ...

&blacksmith.Options{
  PubSub: &pubsub.Options{
    From:       "rabbitmq",
    Enabled:    true,
    Topic:      "blacksmith",
    Connection: "amqp://guest:guest@127.0.0.1:5672/",
    Context:    ctx,
  },
}
```

#### Environment variables

- `RABBIT_SERVER_URL`: The RabbitMQ server URL to dial to leverage pub / sub. If
  `Options.PubSub.Connection` is set, it will override and be used in replacement
  of this environment variable.

  Example: `amqp://guest:guest@127.0.0.1:5672/`


#### Additional context

- `RABBIT_PRODUCER_EXCHANGE`: The RabbitMQ producer exchange to use when publishing
  events.

  Required: no
  
  Type: `string`

  Defaults: `nunchi`
