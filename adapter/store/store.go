package store

/*
InterfaceStore is the string representation for the store interface.
*/
var InterfaceStore = "store"

/*
Store is the interface used to persist the jobs queue in a datastore to keep track
of jobs states.
*/
type Store interface {

	// String returns the string representation of the adapter.
	//
	// Example: "postgres"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// AddEvents inserts a queue of events into the datastore given the data passed
	// in params. It returns an error if any occurred.
	AddEvents(*Toolkit, []*Event) error

	// FindEvent returns a event given the event ID passed in params.
	FindEvent(*Toolkit, string) (*Event, error)

	// FindEvents returns a list of events matching the constraints passed in params.
	// It also returns meta information about the query, such as pagination and the
	// constraints actually applied to it.
	FindEvents(*Toolkit, *WhereEvents) ([]*Event, *Meta, error)

	// AddJobs inserts a list of jobs into the datastore.
	AddJobs(*Toolkit, []*Job) error

	// FindJob returns a job given the job ID passed in params.
	FindJob(*Toolkit, string) (*Job, error)

	// FindJobs returns a list of jobs matching the constraints passed in params.
	// It also returns meta information about the query, such as pagination and the
	// constraints actually applied to it.
	FindJobs(*Toolkit, *WhereEvents) ([]*Job, *Meta, error)

	// AddTransitions inserts a list of transitions into the datastore to update
	// their related job status. We insert new transitions instead of updating the
	// job itself to keep track of the job's history.
	AddTransitions(*Toolkit, []*Transition) error

	// FindTransition returns a transition given the transition ID passed in params.
	FindTransition(*Toolkit, string) (*Transition, error)

	// FindTransitions returns a list of transitions matching the constraints passed
	// in params. It also returns meta information about the query, such as pagination
	// and the constraints really applied to it.
	FindTransitions(*Toolkit, *WhereEvents) ([]*Transition, *Meta, error)

	// Purge purges every events, jobs, and transitions from the store. It is run
	// for each purge policies defined in the store's options, at the defined
	// intervals.
	Purge(*Toolkit, *WhereEvents) error
}
