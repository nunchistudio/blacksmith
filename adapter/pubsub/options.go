package pubsub

import (
	"context"
)

/*
AvailableAdapters is a list of available pubsub adapters.
*/
var AvailableAdapters = map[string]bool{
	"kafka":    true,
	"nats":     true,
	"rabbitmq": true,
}

/*
Defaults are the defaults options set for the pubsub. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{
	Context: context.Background(),
	Enabled: false,
	Topic:   "blacksmith",
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

	// Enabled enables the PubSub interface to distribute jobs to destinations in
	// realtime. If disabled, the scheduler will load jobs to destinations given the
	// schedule of each destination and action (if applicable).
	Enabled bool `json:"enabled"`

	// Connection is the connection string to connect to the pubsub.
	Connection string `json:"-"`

	// Topic is the topic name the pubsub adapter will use to publish and subscribe
	// messages to.
	Topic string `json:"topic"`
}
