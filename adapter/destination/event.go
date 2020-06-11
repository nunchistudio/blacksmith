package destination

import (
	"time"

	"github.com/nunchistudio/blacksmith/adapter/wanderer"
	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
InterfaceEvent is the string representation for the destination's event interface.
*/
var InterfaceEvent = "destination/event"

/*
Event represents an event of a destination adapter. An Event contains all the
logic to handle a specific event for a destination.
*/
type Event interface {

	// String returns the string representation of the destination's event.
	//
	// Examples: "*" (wildcard), "identify"
	String() string

	// Migrations returns the list of all migrations for the event, regardless their
	// status.
	//
	// The adapter can use the package helper/sqlike to easily read migrations files
	// from a directory. See package helper/sqlike for more details.
	//
	// Note: Feature only available in Blacksmith Enterprise.
	Migrations() ([]*wanderer.Migration, error)

	// Schedule represents a schedule at which an event should run. When returning
	// nil, the destination's schedule is applied.
	Schedule() *Schedule

	// Marshal is in charge of the "T" in the ETL process: it needs to Transform the
	// data of the pointer receiver originally passed by other events. It must return
	// a payload including the context and data as JSON marshaled values.
	//
	// Note: If the context in the returned payload is nil, the one from the
	// event will automatically be applied.
	//
	// Note: If the function returns an error, the event can not be considered as
	// transformed. Therefore, no jobs will be created and the event will never run.
	Marshal(*Toolkit) (*Payload, error)

	// Run is in charge of the "L" in the ETL process: it Loads the data to the
	// destination. It is executed either on a schedule basis or in realtime when
	// applicable.
	//
	// The toolkit gives access to the logger and original event. When done, the
	// function can return a list of destinations' events to trigger depending on
	// on the status of the current job. This allows to trigger jobs both in tandem
	// and in cascade. Every events will then be processed by the scheduler, respecting
	// the scheduling options of each one.
	Run(*Toolkit) (*Then, error)
}

/*
Payload represents the fields an event must fill when being loaded into the
destination.
*/
type Payload struct {

	// Context is a dictionary of information that provides useful context about an
	// event. The context should be used inside every events for consistency.
	//
	// It must be a valid JSON since it will be used using json Marshal and Unmarshal
	// functions.
	Context []byte `json:"context"`

	// Data is the byte representation of the data sent by the event.
	//
	// It must be a valid JSON since it will be used using json Marshal and Unmarshal
	// functions.
	Data []byte `json:"data"`

	// SentAt allows you to keep track of the timestamp when the event was originally
	// sent.
	SentAt *time.Time `json:"sent_at,omitempty"`
}

/*
Then allows events to call other events from any destination depending on the
job status.

Example:

  then := &destination.Then{
    OnSucceeded: map[string][]Event{
      "destinationName": []Event{
        EventA{},
        EventB{},
      }
    }
  }
*/
type Then struct {

	// List of destinations events to trigger in case the job has succeeded.
	OnSucceeded map[string][]Event `json:"on_succeeded,omitempty"`

	// List of destinations events to trigger in case the job has failed.
	OnFailed map[string][]Event `json:"on_failed,omitempty"`

	// List of destinations events to trigger in case the job has been discarded.
	OnDiscarded map[string][]Event `json:"on_discarded,omitempty"`
}

/*
validateEvent makes sure an event of a destination adapter is ready to be used
properly by a Blacksmith application.
*/
func validateEvent(e Event) []errors.Validation {
	validations := []errors.Validation{}

	// Verify the event ID is not empty.
	if e.String() == "" {
		validations = append(validations, errors.Validation{
			Message: "Event ID must not be empty",
			Path:    []string{"Event", "unknown", "String()"},
		})
	}

	// Return the validation errors if any occurred.
	if len(validations) > 0 {
		return validations
	}

	return nil
}
