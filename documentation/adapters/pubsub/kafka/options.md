#### Configuration

```go
ctx := context.Background()
ctx = context.WithValue(ctx, "KAFKA_CONSUMER_GROUP", "consumer-group-name")

// ...

&blacksmith.Options{
  PubSub: &pubsub.Options{
    From:       "kafka",
    Context:    ctx,
    Enabled:    true,
    Topic:      "blacksmith",
    Connection: "127.0.0.1:9092",
  },
}
```

#### Environment variables

- `KAFKA_BROKERS`: The Kafka broker URLs to dial to leverage pub / sub. If
  `Options.PubSub.Connection` is set, it will override and be used in replacement
  of this environment variable. The value is a comma-delimited list of hosts.

  Example: `127.0.0.1:9092,127.0.0.2:9092`


#### Additional context

- `KAFKA_CONSUMER_GROUP`: The Kafka consumer group to use when subscribing to
  events.

  Required: no
  
  Type: `string`

  Defaults: `nunchi`
