package pubsub

import (
	"fmt"

	"github.com/nunchistudio/blacksmith/adapter/store"
	"github.com/nunchistudio/blacksmith/helper/errors"
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

/*
validatePubSub makes sure the pubsub adapter is ready to be used properly by a
Blacksmith application.
*/
func validatePubSub(ps PubSub) error {

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "pubsub: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Verify the pubsub ID is not empty.
	if ps.String() == "" {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "PubSub ID must not be empty",
			Path:    []string{"PubSub", "unknown", "String()"},
		})

		return fail
	}

	// We now can add the adapter name to the error message.
	fail.Message = fmt.Sprintf("pubsub/%s: Failed to load adapter", ps.String())

	// It is impossible to deal with nil options.
	if ps.Options() == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "PubSub options must not be nil",
			Path:    []string{"PubSub", ps.String(), "Options()"},
		})

		return fail
	}

	// If the adapter didn't set a topic, use the default one.
	if ps.Options().Topic == "" {
		ps.Options().Topic = Defaults.Topic
	}

	// Avoid cycles.
	ps.Options().Load = nil

	// Return the error if any occurred.
	if len(fail.Validations) > 0 {
		return fail
	}

	return nil
}
