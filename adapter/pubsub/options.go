package pubsub

import (
	"context"
)

/*
AvailableAdapters is a list of available pubsub adapters.
*/
var AvailableAdapters = map[string]bool{
	"aws/snssqs":       true,
	"azure/servicebus": true,
	"google/pubsub":    true,
	"kafka":            true,
	"nats":             true,
	"rabbitmq":         true,
}

/*
Defaults are the defaults options set for the pubsub. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{
	Context:      context.Background(),
	Topic:        "blacksmith",
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

	// Topic is the topic name the pubsub adapter will use to publish messages.
	//
	// Example for Amazon Web Services: "arn:aws:sns:<region>:<id>:<topic>"
	// Example for Microsoft Azure: "<topic>"
	// Example for Google Cloud: "projects/<project>/topics/<topic>"
	// Example for Kafka: "<topic>"
	// Example for NATS: "<subject>"
	// Example for RabbitMQ: "<exchange>"
	Topic string `json:"topic"`

	// Subscription is the queue or subscription name the pubsub adapter will use
	// to subscribe to messages.
	//
	// Example for Amazon Web Services: "arn:aws:sqs:<region>:<id>:<queue>"
	// Example for Microsoft Azure: "<subscription>"
	// Example for Google Cloud: "projects/<project>/subscriptions/<subscription>"
	// Example for Kafka: "<consumer-group>"
	// Example for NATS: "<queue>"
	// Example for RabbitMQ: "<queue>"
	Subscription string `json:"subscription"`
}
