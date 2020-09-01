package pubsub

import (
	"github.com/nunchistudio/blacksmith/adapter/store"
)

/*
InterfacePubSub is the string representation for the pubsub interface.
*/
var InterfacePubSub = "pubsub"

/*
PubSub is the interface used to load events' jobs to destinations in realtime. When
disabled, the gateway and scheduler will work as expected but will load jobs to
destinations given the configured schedule.
*/
type PubSub interface {

	// String returns the string representation of the adapter.
	//
	// Example: "nats"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// Publisher returns the interface in charge of publishing messages in realtime.
	// Can be nil if PubSub is disabled.
	Publisher() Publisher

	// Subscriber returns the interface in charge of subscribing to messages in
	// realtime. Can be nil if PubSub is disabled.
	Subscriber() Subscriber
}

/*
Publisher is in charge of creating topics and sending messages to the Subscriber.
*/
type Publisher interface {

	// Init let you initialize the Publisher.
	Init(*Toolkit) error

	// Send publishes a queue. It only returns after the queue has been sent, or
	// failed to be sent.
	Send(*Toolkit, *store.Queue) error

	// Shutdown flushes pending message sends and disconnects the Publisher. It only
	// return after all pending messages have been sent.
	Shutdown(*Toolkit) error
}

/*
Subscriber is in charge of receiving messages on given topics.
*/
type Subscriber interface {

	// Init let you initialize the Subscriber.
	Init(*Toolkit) error

	// Receive receives and returns the next queue from the Subscriber, blocking and
	// polling if none are available.
	Receive(*Toolkit) (*store.Queue, error)

	// Shutdown flushes pending ack sends and disconnects the Subscriber.
	Shutdown(*Toolkit) error
}
