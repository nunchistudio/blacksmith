package wanderer

import (
	"fmt"

	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/helper/mutex"
)

/*
InterfaceWanderer is the string representation for the wanderer interface.
*/
var InterfaceWanderer = "wanderer"

/*
Wanderer is the interface used to create a new migration wanderer. This acts with
a remote mutex allowing teams to safely run migrations with no conflicts.
*/
type Wanderer interface {

	// String returns the string representation of the adapter.
	//
	// Example: "postgres"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// Migrations returns the list of all migrations for the wanderer, regardless
	// their status.
	//
	// The adapter can use the sqlike package to easily read migrations files from
	// a directory. See sqlike package for more details.
	Migrations() ([]*Migration, error)

	// Migrate is the function called for running a migration for the wanderer. This
	// function is the migration logic for running every migrations of the wanderer.
	// When being executed, the function has access to a toolkit and the desired
	// migration.
	//
	// The adapter can use the sqlike package to easily run a SQL migration within
	// a transaction. See sqlike package for more details.
	Migrate(*Toolkit, *Migration) error

	// Mutex returns a remote mutex so the wanderer can safely act across a team or
	// several teams to lock migrations state.
	Mutex() mutex.Mutex

	// Ack acknowledges every migrations into the wanderer datastore regardless their
	// status. It needs a mutex lock to ensure migrations state.
	Ack(*Toolkit, []*Migration) error

	// Find returns a list of acknowledged migrations given some properties passed
	// in params. It needs a mutex lock to ensure migrations state.
	Find(*Toolkit, *WhereMigration) ([]*Migration, error)

	// Update updates migrations status into the wanderer datastore. It needs a mutex
	// lock to ensure migrations state.
	Update(*Toolkit, []*Migration) error
}

/*
validateWanderer makes sure the wanderer adapter is ready to be used properly by
a Blacksmith application.
*/
func validateWanderer(w Wanderer) error {

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "wanderer: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Verify the wanderer ID is not empty.
	if w.String() == "" {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Wanderer ID must not be empty",
			Path:    []string{"Wanderer", "unknown", "String()"},
		})

		return fail
	}

	// We now can add the adapter name to the error message.
	fail.Message = fmt.Sprintf("wanderer/%s: Failed to load adapter", w.String())

	// It is impossible to deal with nil options.
	if w.Options() == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Wanderer options must not be nil",
			Path:    []string{"Wanderer", w.String(), "Options()"},
		})

		return fail
	}

	// Avoid cycles.
	w.Options().Load = nil

	// Return the error if any occurred.
	if len(fail.Validations) > 0 {
		return fail
	}

	return nil
}
