package store

import (
	"fmt"

	"github.com/nunchistudio/blacksmith/adapter/wanderer"
	"github.com/nunchistudio/blacksmith/helper/errors"
)

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

	// Migrations returns the list of all migrations for the store, regardless
	// their status.
	//
	// The adapter can use the package helper/sqlike to easily read migrations files
	// from a directory. See package helper/sqlike for more details.
	//
	// Note: Feature only available in Blacksmith Enterprise.
	Migrations() ([]*wanderer.Migration, error)

	// Migrate is the function called for running a migration for the store. This
	// function is the migration logic for running every migrations of the store.
	// When being executed, the function has access to a toolkit and the desired
	// migration.
	//
	// It is important to understand that it is up to the adapter to run the migration
	// within a transaction (when applicable).
	//
	// Note: Feature only available in Blacksmith Enterprise.
	Migrate(*wanderer.Toolkit, *wanderer.Migration) error

	// InsertQueue inserts a queue of events into the datastore given the data passed
	// in params. It returns an error if any occurred. This method shall be called
	// by the gateway when a new event happens.
	InsertQueue(*Toolkit, *Queue) error

	// AddJobs inserts a list of jobs into the datastore.
	AddJobs(*Toolkit, []*Job) error

	// AddTransitions inserts a list of transitions into the datastore to update
	// their related job status. We insert new transitions instead of updating the
	// job itself to keep track of job history. This method shall be called by the
	// scheduler for registering jobs transitions.
	AddTransitions(*Toolkit, []*Transition) error

	// Find returns a list of acknowledged jobs' events given some properties passed
	// in params. Returned jobs are grouped by events.
	Find(*Toolkit, *WhereJob) ([]*Event, error)

	// Delete deletes a list of given events from the store. It is useful to clear
	// unused events so the store can be more performant.
	Delete(*Toolkit, []*Event) error
}

/*
validateStore makes sure the store adapter is ready to be used properly by a
Blacksmith application.
*/
func validateStore(s Store) error {

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "store: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Verify the store ID is not empty.
	if s.String() == "" {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Store ID must not be empty",
			Path:    []string{"Store", "unknown", "String()"},
		})

		return fail
	}

	// We now can add the adapter name to the error message.
	fail.Message = fmt.Sprintf("store/%s: Failed to load adapter", s.String())

	// It is impossible to deal with nil options.
	if s.Options() == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Store options must not be nil",
			Path:    []string{"Store", s.String(), "Options()"},
		})

		return fail
	}

	// Avoid cycles.
	s.Options().Load = nil

	// Return the error if any occurred.
	if len(fail.Validations) > 0 {
		return fail
	}

	return nil
}
