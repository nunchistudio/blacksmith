package destination

import (
	"fmt"

	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
InterfaceDestination is the string representation for the destination interface.
*/
var InterfaceDestination = "destination"

/*
Destination is the interface used to load events to third-party services. Those
can be of any kind, such as APIs or databases.
*/
type Destination interface {

	// String returns the string representation of the adapter.
	//
	// Example: "zendesk"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// Actions returns a list of actions the destination can handle. Destinations'
	// actions are run from sources' triggers and can also be triggered by other
	// destinations' actions. When a destination's action is called, it is
	// represented as a "job" across the ecosystem.
	Actions() map[string]Action
}

/*
validateDestination makes sure a destination adapter is ready to be used properly
by a Blacksmith application.
*/
func validateDestination(d Destination) error {

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "destination: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Verify the destination ID is not empty.
	if d.String() == "" {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Destination ID must not be empty",
			Path:    []string{"Destination", "unknown", "String()"},
		})

		return fail
	}

	// We now can add the adapter name to the error message.
	fail.Message = fmt.Sprintf("destination/%s: Failed to load adapter", d.String())

	// It is impossible to deal with nil options.
	if d.Options() == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Destination options must not be nil",
			Path:    []string{"Destination", d.String(), "Options()"},
		})

		return fail
	}

	// Verify the destination default scheduling options.
	if d.Options().DefaultSchedule != nil {
		if d.Options().DefaultSchedule.Interval == "" {
			d.Options().DefaultSchedule.Interval = Defaults.DefaultSchedule.Interval
		}

		if d.Options().DefaultSchedule.MaxRetries == 0 {
			d.Options().DefaultSchedule.MaxRetries = Defaults.DefaultSchedule.MaxRetries
		}
	} else {
		d.Options().DefaultSchedule = Defaults.DefaultSchedule
	}

	// Make sure the destination returns a collection of actions. Otherwise we can
	// not continue.
	if d.Actions() == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Destination actions must not be nil",
			Path:    []string{"Destination", d.String(), "Actions()"},
		})

		return fail
	}

	// Validate every actions of the destination and add the validation errors.
	for _, a := range d.Actions() {
		errs := validateAction(d.String(), a)
		if errs != nil {
			fail.Validations = append(fail.Validations, errs...)
		}
	}

	// Avoid cycles.
	d.Options().Load = nil

	// Return the error if any occurred.
	if len(fail.Validations) > 0 {
		return fail
	}

	return nil
}
