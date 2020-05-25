package errors

/*
Meta is used when dealing with HTTP errors to provide a consistent HTTP response
across adapters.
*/
type Meta struct {

	// Event contains the fields related to an event when dealing with HTTP errors.
	Event *Event `json:"event,omitempty"`
}

/*
Event is the representation of an event inside error.
*/
type Event struct {

	// ID is the generated ID of the event.
	ID string `json:"id"`
}
