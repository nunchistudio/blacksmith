package source

import (
	"github.com/nunchistudio/blacksmith/adapter/pubsub"
)

/*
ModeSubscriber is used to indicate the event is triggered from a message received
by a subscriber in a Pub / Sub mechanism.
*/
var ModeSubscriber = "subscriber"

/*
TriggerSubscriber is the interface used for triggers using a Pub / Sub topic.
*/
type TriggerSubscriber interface {

	// Extract in charge of the "E" in the ETL process: it Extracts the data from
	// the source.
	Extract(*Toolkit, *pubsub.Message) (*Payload, error)
}

/*
Subscription contains the details about a subscription used by the gateway and
the pubsub adapter.
*/
type Subscription struct {

	// Broker is the middleman's name to use for subscribing to messages. It is not
	// applicable for every adapters. It can be used to group messages per queue.
	// When applicable, the broker is required.
	//
	// Example for Kafka: "<consumer-group>"
	// Example for NATS: "<queue>"
	// Example for RabbitMQ: N/A
	// Example for Amazon Web Services: N/A
	// Example for Google Cloud: N/A
	// Example for Microsoft Azure: "<topic>"
	Broker string `json:"broker,omitempty"`

	// Subscription is the name used by the pubsub adapter to subscribe to messages.
	// Unlike the broker, the ubscription is applicable and required by all adapters.
	//
	// Example for Kafka: "<topic>"
	// Example for NATS: "<subject>"
	// Example for RabbitMQ: "<queue>"
	// Example for Amazon Web Services: "arn:aws:sqs:<region>:<id>:<queue>"
	// Example for Google Cloud: "projects/<project>/subscriptions/<subscription>"
	// Example for Microsoft Azure: "<subscription>"
	Subscription string `json:"subscription"`
}
