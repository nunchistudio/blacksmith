package pubsub

/*
Message holds the information of a message received by the subscriber for the
source triggers.
*/
type Message struct {

	// Body is the marshaled content of the message.
	Body []byte `json:"body"`

	// Metadata can hold some metadata about the message.
	Metadata map[string]string `json:"meta"`
}
