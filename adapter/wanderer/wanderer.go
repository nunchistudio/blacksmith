package wanderer

import (
	"fmt"

	"github.com/nunchistudio/blacksmith/helper/errors"
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

	// Ack acknowledges every migrations into the wanderer datastore regardless their
	// status.
	Ack(*Toolkit, []*Migration) error

	// Find returns a list of acknowledged migrations given some properties passed
	// in params.
	Find(*Toolkit, *WhereMigration) ([]*Migration, error)

	// Update updates migrations status into the wanderer datastore.
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
