package pubsub

import (
	"context"
)

/*
AvailableAdapters is a list of available pubsub adapters.
*/
var AvailableAdapters = map[string]bool{
	"aws":      true,
	"azure":    true,
	"google":   true,
	"kafka":    true,
	"nats":     true,
	"rabbitmq": true,
}

/*
Defaults are the defaults options set for the pubsub. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{
	Context:      context.Background(),
	Topic:        "blacksmith",
	Broker:       "blacksmith",
	Subscription: "blacksmith",
}

/*
Options is the options a user can pass to use the pubsub adapter.
*/
type Options struct {

	// From is used to set the desired pubsub adapter. It must be one of
	// AvailableAdapters.
	From string `json:"from,omitempty"`

	// Context is a free key-value dictionary that will be passed to the underlying
	// adapter.
	Context context.Context `json:"-"`

	// Connection is the connection string to connect to the pubsub.
	Connection string `json:"-"`

	// Topic is the topic name the pubsub adapter will use to publish messages to.
	//
	// Example for Kafka: "<topic>"
	// Example for NATS: "<subject>"
	// Example for RabbitMQ: "<exchange>"
	// Example for Amazon Web Services: "arn:aws:sns:<region>:<id>:<topic>"
	// Example for Google Cloud: "projects/<project>/topics/<topic>"
	// Example for Microsoft Azure: "<topic>"
	Topic string `json:"topic"`

	// Broker is the middleman's name to use for pushing or subscribing to messages.
	// It is not applicable for every adapters. It can be used to group messages per
	// queue and therefore help the adapter create a load balancing or ensure a
	// single active consumer.
	//
	// Example for Kafka: "<consumer-group>"
	// Example for NATS: "<queue>"
	// Example for RabbitMQ: N/A
	// Example for Amazon Web Services: N/A
	// Example for Google Cloud: N/A
	// Example for Microsoft Azure: N/A
	Broker string `json:"broker,omitempty"`

	// Subscription is the subscription name the pubsub adapter will use to subscribe
	// to messages when different from the topic.
	//
	// Example for Kafka: N/A
	// Example for NATS: N/A
	// Example for RabbitMQ: "<queue>"
	// Example for Amazon Web Services: "arn:aws:sqs:<region>:<id>:<queue>"
	// Example for Google Cloud: "projects/<project>/subscriptions/<subscription>"
	// Example for Microsoft Azure: "<subscription>"
	Subscription string `json:"subscription,omitempty"`
}
