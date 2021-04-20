package source

import (
	"github.com/nunchistudio/blacksmith/adapter/pubsub"
)

/*
ModeSubscription is used to indicate the event is triggered from a message received
by a subscription in a Pub / Sub mechanism.
*/
var ModeSubscription = "subscription"

/*
TriggerSubscription is the interface used for triggers using a Pub / Sub topic.

A new subscription trigger can be generated using the Blacksmith CLI:

  $ blacksmith generate trigger --name <name> --mode sub [--path <path>] [--migrations]
*/
type TriggerSubscription interface {

	// Extract in charge of the "E" in the ETL process: it Extracts the data from
	// the source.
	Extract(*Toolkit, *pubsub.Message) (*Event, error)
}

/*
Subscription contains the details about a subscription used by the gateway and
the pubsub adapter.
*/
type Subscription struct {

	// Topic is the topic name the pubsub adapter will use when it is required
	// for working in tandem with the subscription name.
	//
	// Example for Kafka: "<topic>"
	// Example for NATS: "<subject>"
	// Example for RabbitMQ: N/A
	// Example for Amazon Web Services: N/A
	// Example for Google Cloud: N/A
	// Example for Microsoft Azure: "<topic>"
	Topic string `json:"topic,omitempty"`

	// Subscription is the queue or subscription name the pubsub adapter will use
	// to subscribe to messages.
	//
	// Example for Kafka: "<consumer-group>"
	// Example for NATS: "<queue>"
	// Example for RabbitMQ: "<queue>"
	// Example for Amazon Web Services: "arn:aws:sqs:<region>:<id>:<queue>"
	// Example for Google Cloud: "projects/<project>/subscriptions/<subscription>"
	// Example for Microsoft Azure: "<subscription>"
	Subscription string `json:"subscription"`
}
