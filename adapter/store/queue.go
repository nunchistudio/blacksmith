package store

import (
	"time"
)

/*
StatusAcknowledged is used to mark a job as acknowledged. This is used when
registering new jobs into the store.
*/
var StatusAcknowledged = "acknowledged"

/*
StatusAwaiting is used to mark a job as awaiting. This is used when a job is
awaiting to be run.
*/
var StatusAwaiting = "awaiting"

/*
StatusExecuting is used to mark a job as executing. This is used when a job is
being executed.
*/
var StatusExecuting = "executing"

/*
StatusSucceeded is used to mark a job as succeeded.
*/
var StatusSucceeded = "succeeded"

/*
StatusFailed is used to mark a job as failed.
*/
var StatusFailed = "failed"

/*
StatusDiscarded is used to mark a job as discarded.
*/
var StatusDiscarded = "discarded"

/*
Queue keeps track of incoming events, their jobs, and their jobs' transitions.
*/
type Queue struct {

	// Events is the collection of incoming or awaiting events.
	Events []*Event `json:"events"`
}

/*
Event define the fields stored in the datastore about an event.
*/
type Event struct {

	// ID is the unique identifier of the event. It must be a valid KSUID.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	ID string `json:"id"`

	// Source is the string representation of the incoming event's source.
	Source string `json:"source"`

	// Name is the string representation of the incoming or awaiting event.
	Name string `json:"name"`

	// Context is the marshaled representation of the "context" key presents in the
	// event's payload.
	Context []byte `json:"context,omitempty"`

	// Data is the marshaled representation of the "data" key presents in the event's
	// payload.
	Data []byte `json:"data,omitempty"`

	// Jobs is a list of jobs to execute related to the event. A destination should
	// have at most 2 jobs per event: a wildcard and the specific event.
	Jobs []*Job `json:"jobs"`

	// SentAt is the timestamp of when the event is originally sent by the source.
	// It can be nil if none was provided.
	SentAt *time.Time `json:"sent_at,omitempty"`

	// ReceivedAt is the timestamp of when the event is received by the gateway.
	// This shall always be overridden by the gateway.
	ReceivedAt time.Time `json:"received_at"`

	// IngestedAt is a timestamp of the event creation date into the store.
	// It can be nil if none was provided.
	IngestedAt *time.Time `json:"ingested_at,omitempty"`
}

/*
Job is the definition of a job that needs to run for a given event and a specific
destination.
*/
type Job struct {

	// ID is the unique identifier of the job. It must be a valid KSUID.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	ID string `json:"id"`

	// Destination is the string representation of the destination the job needs to
	// run to.
	Destination string `json:"destination"`

	// Event is the string representation of the incoming or awaiting event. It is
	// also present in the job so we can make a distinction between specific events
	// and wildcard events.
	Event string `json:"event"`

	// Context is the marshaled representation of the "context" key presents in the
	// event's payload when the destination's event has been marshaled.
	Context []byte `json:"context,omitempty"`

	// Data is the marshaled representation of the "data" key presents in the event's
	// payload when the destination's event has been marshaled.
	Data []byte `json:"data,omitempty"`

	// Transitions is an array of the job's transitions. It is used to keep track of
	// successes, failures, and errors so the store is aware of the job's status.
	//
	// Note: It is up to the adapter to only return the latest job's transition since
	// this is the only one that really matters in this context.
	Transitions [1]*Transition `json:"transitions"`

	// CreatedAt is a timestamp of the job creation date into the store.
	CreatedAt time.Time `json:"created_at"`

	// EventID is the ID of the event related to this job.
	EventID string `json:"event_id,omitempty"`

	// ParentJobID is the ID of the parent job ID.
	ParentJobID *string `json:"parent_job_id,omitempty"`
}

/*
Transition represents a job's transition to keep track of its states.
*/
type Transition struct {

	// ID is the unique identifier of the transition. It must be a valid KSUID.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	ID string `json:"id"`

	// Attempt represents the number of tentatives that the job has run before
	// succeeded.
	Attempt uint16 `json:"attempt"`

	// StateBefore is the state of the job before running the new transition. See
	// status details for more info. This shall be nil when receiving the job from
	// the gateway.
	StateBefore *string `json:"state_before"`

	// StateAfter is the state of the job after running the new transition. See
	// status details for more info.
	StateAfter string `json:"state_after"`

	// Error keeps track of encountered error if any.
	Error error `json:"error,omitempty"`

	// CreatedAt is a timestamp of the transition creation date into the store.
	CreatedAt time.Time `json:"created_at"`

	// EventID is the ID of the event related to this job's transition.
	EventID string `json:"event_id,omitempty"`

	// JobID is the ID of the job related to this transition.
	JobID string `json:"job_id,omitempty"`
}
