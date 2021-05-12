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
	// Format for AWS SNS / SQS: N/A
	// Format for Azure Service Bus: "<topic>"
	// Format for Google Pub / Sub: N/A
	// Format for Apache Kafka: "<topic>"
	// Format for NATS: "<subject>"
	// Format for RabbitMQ: N/A
	Topic string `json:"topic,omitempty"`

	// Subscription is the queue or subscription name the pubsub adapter will use
	// to subscribe to messages.
	//
	// Format for AWS SNS / SQS: "arn:aws:sqs:<region>:<id>:<queue>"
	// Format for Azure Service Bus: "<subscription>"
	// Format for Google Pub / Sub: "projects/<project>/subscriptions/<subscription>"
	// Format for Apache Kafka: "<consumer-group>"
	// Format for NATS: "<queue>"
	// Format for RabbitMQ: "<queue>"
	Subscription string `json:"subscription"`
}
