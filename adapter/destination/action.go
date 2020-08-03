package destination

import (
	"time"

	"github.com/nunchistudio/blacksmith/adapter/store"
	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
InterfaceAction is the string representation for the destination's action interface.
*/
var InterfaceAction = "destination/action"

/*
Action represents a specific action to run against a destination.
*/
type Action interface {

	// String returns the string representation of the destination's action.
	//
	// Example: "identify"
	String() string

	// Schedule represents a schedule at which an action should run. When returning
	// nil, the parent destination's schedule is applied.
	Schedule() *Schedule

	// Marshal is in charge of the "T" in the ETL process: it needs to Transform the
	// data of the pointer receiver originally passed by sources' triggers or
	// destinations' actions. It must return a payload including the context and
	// data as JSON marshaled values.
	//
	// Note: If the context in the returned payload is nil, the one from the event
	// will automatically be applied.
	//
	// Note: If the function returns an error, the event can not be considered as
	// transformed. Therefore, no jobs will be created and the action will never run.
	Marshal(*Toolkit) (*Payload, error)

	// Run is in charge of the "L" in the ETL process: it Loads the data to the
	// destination's endpoint. It is executed either on a schedule basis or in realtime
	// when applicable.
	//
	// The queue only includes received events triggering the action. The jobs inside
	// each event are therefore specific to this action only. The jobs not related
	// to this destination's action are not included.
	//
	// When done, the function can return a list of destinations' actions to run
	// depending on on the status of the current job. This allows to trigger jobs
	// both in tandem and in cascade. Every events will then be processed by the
	// scheduler, respecting the scheduling options of each one.
	Run(*Toolkit, *store.Queue) (*Then, error)
}

/*
Payload represents the fields an action must fill when being loaded into the
destination.
*/
type Payload struct {

	// Context is a dictionary of information that provides useful context about an
	// event. The context should be used inside every events for consistency.
	//
	// Note: It must be a valid JSON since it will be used using encoding/json Marshal
	// and Unmarshal functions.
	Context []byte `json:"context"`

	// Data is the byte representation of the data sent by the event.
	//
	// Note: It must be a valid JSON since it will be used using encoding/json Marshal
	// and Unmarshal functions.
	Data []byte `json:"data"`

	// SentAt allows you to keep track of the timestamp when the event was originally
	// sent.
	SentAt *time.Time `json:"sent_at,omitempty"`
}

/*
Then allows actions to call other actions from any destination depending on the
job status.

Example:

  then := &destination.Then{
    OnSucceeded: map[string][]Action{
      "destinationName": []Action{
        ActionA{},
        ActionB{},
      }
    }
  }
*/
type Then struct {

	// List of destinations actions to run in case the job has succeeded.
	OnSucceeded map[string][]Action `json:"on_succeeded,omitempty"`

	// List of destinations actions to run in case the job has failed.
	OnFailed map[string][]Action `json:"on_failed,omitempty"`

	// List of destinations actions to run in case the job has been discarded.
	OnDiscarded map[string][]Action `json:"on_discarded,omitempty"`
}

/*
validateAction makes sure an action of a destination adapter is ready to be used
properly by a Blacksmith application. We do not need to validate the scheduling
options since the ones of the parent destination are used in case of nil.
*/
func validateAction(d string, a Action) []errors.Validation {
	validations := []errors.Validation{}

	// Verify the action ID is not empty.
	if a.String() == "" {
		validations = append(validations, errors.Validation{
			Message: "Action ID must not be empty",
			Path:    []string{"Destination", d, "Actions()", "unknown", "String()"},
		})
	}

	// Return the validation errors if any occurred.
	if len(validations) > 0 {
		return validations
	}

	return nil
}
