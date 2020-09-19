package destination

import (
	"time"

	"github.com/nunchistudio/blacksmith/adapter/store"
)

/*
InterfaceAction is the string representation for the destination's action interface.
*/
var InterfaceAction = "destination/action"

/*
Actions is used to return a slice of Action grouped by their destination name.
This is used by the package flow when creating flow to distribute data from
triggers to actions.
*/
type Actions map[string][]Action

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

	// Marshal is in charge of marshalling the received data for the action. It
	// can be in charge of the "T" in the ETL process if needed: it can Transform
	// the data of the pointer receiver originally passed by sources' triggers or
	// destinations' actions. It must return a payload including the context and
	// data as JSON marshaled values.
	//
	// If the context in the returned payload is nil, the one from the event will
	// automatically be applied.
	//
	// If the function returns an error, the event can not be considered as transformed.
	// Therefore, no jobs will be created and the action will never run.
	Marshal(*Toolkit) (*Payload, error)

	// Load is in charge of the "L" in the ETL process: it Loads the data to the
	// destination's endpoint. It is executed either on a schedule basis or in realtime
	// when applicable.
	//
	// The queue only includes received events triggering the action. The jobs inside
	// each event are therefore specific to this action only.
	//
	// When desired, the function can return a list of destinations' actions (in
	// Then) to run depending on on the status of the current job. Every jobs will
	// then be processed by the scheduler, respecting the scheduling options of
	// each one.
	Load(*Toolkit, *store.Queue, chan<- Then)
}

/*
Payload represents the fields an action must fill when being loaded into the
destination.
*/
type Payload struct {

	// Version is the version number of the destination used by a flow when executed.
	//
	// Examples: "v1.0", "2020-10-01"
	Version string `json:"version,omitempty"`

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
Then allows to inform the scheduler of job status and to execute other actions from
the same destination depending on the job status.
*/
type Then struct {

	// Jobs is the list of the job IDs being processed. It informs the scheduler
	// of the status of the desired jobs.
	//
	// When nil or empty, all jobs from the queue will be affected by the result.
	// This allows to either load the data entry-per-entry or in batch if the
	// destination allows it. If a job ID is not returned, the scheduler will not
	// be aware of its status and will mark it as "unknown".
	Jobs []string `json:"jobs"`

	// Error is the error encountered when loading the data into the destination's
	// action.
	//
	// When not nil the related jobs will be either marked as "failed" or "discarded"
	// given the max retries of the action. OnFailed or OnDiscarded will automatically
	// be applied by the scheduler as jobs to be executed. When Error is nil the jobs
	// are marked as "succeeded" and OnSucceeded will be applied.
	Error error `json:"error"`

	// ForceDiscard manually marks a job as discarded. It is useful if you know it
	// is impossible for the job to succeed even after multiple retries.
	//
	// When set, Error must not be nil.
	ForceDiscard bool `json:"is_force_discarded"`

	// List of destinations actions to run in case the job has succeeded.
	OnSucceeded []Action `json:"on_succeeded,omitempty"`

	// List of destinations actions to run in case the job has failed.
	OnFailed []Action `json:"on_failed,omitempty"`

	// List of destinations actions to run in case the job has been discarded.
	OnDiscarded []Action `json:"on_discarded,omitempty"`
}
