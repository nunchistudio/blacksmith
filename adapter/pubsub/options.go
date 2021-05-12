package pubsub

/*
Driver is a custom type allowing the user to only pass supported drivers when
configuring the Pub / Sub adapter.
*/
type Driver string

/*
DriverAWSSNSSQS is used to leverage AWS SNS / SQS as the Pub / Sub adapter.
*/
var DriverAWSSNSSQS Driver = "aws/snssqs"

/*
DriverAzureServiceBus is used to leverage Azure Service Bus as the Pub / Sub
adapter.
*/
var DriverAzureServiceBus Driver = "azure/servicebus"

/*
DriverGooglePubSub is used to leverage Google Pub / Sub as the Pub / Sub adapter.
*/
var DriverGooglePubSub Driver = "google/pubsub"

/*
DriverKafka is used to leverage Apache Kafka as the Pub / Sub adapter.
*/
var DriverKafka Driver = "kafka"

/*
DriverNATS is used to leverage NATS as the Pub / Sub adapter.
*/
var DriverNATS Driver = "nats"

/*
DriverRabbitMQ is used to leverage RabbitMQ as the Pub / Sub adapter.
*/
var DriverRabbitMQ Driver = "rabbitmq"

/*
Defaults are the defaults options set for the pubsub. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{
	Topic:        "blacksmith",
	Subscription: "blacksmith",
}

/*
Options is the options a user can pass to use the pubsub adapter.
*/
type Options struct {

	// From is used to set the desired driver for the Pub / Sub adapter.
	From Driver `json:"from,omitempty"`

	// Connection is the connection string to connect to the pubsub.
	Connection string `json:"-"`

	// Topic is the topic name the pubsub adapter will use to publish messages.
	//
	// Format for AWS SNS / SQS: "arn:aws:sns:<region>:<id>:<topic>"
	// Format for Azure Service Bus: "<topic>"
	// Format for Google Pub / Sub: "projects/<project>/topics/<topic>"
	// Format for Apache Kafka: "<topic>"
	// Format for NATS: "<subject>"
	// Format for RabbitMQ: "<exchange>"
	Topic string `json:"topic"`

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
