package destination

import (
	"fmt"

	"github.com/nunchistudio/blacksmith/adapter/wanderer"
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

	// Migrations returns the list of all migrations for the destination, regardless
	// their status. This does not contain event-specific migrations but only the
	// ones needed for the destination itself.
	//
	// The adapter can use the package helper/sqlike to easily read migrations files
	// from a directory. See package helper/sqlike for more details.
	//
	// Note: Feature only available in Blacksmith Enterprise.
	Migrations() ([]*wanderer.Migration, error)

	// Migrate is the function called for running a migration for this destination
	// and all of its events. This function is the migration logic for running every
	// migrations of this destination. When being executed, the function has access
	// to a toolkit and the desired migration.
	//
	// It is important to understand that it is up to the adapter to run the migration
	// within a transaction (when applicable).
	//
	// Note: Feature only available in Blacksmith Enterprise.
	Migrate(*wanderer.Toolkit, *wanderer.Migration) error

	// Events returns a list of events the destination can handle. Destinations'
	// events are triggered from sources' events and can also be triggered by
	// destinations' events. When a destination's event is triggered, it is
	// represented as a 'job' across the ecosystem.
	//
	// An event with "*" as its key represents a wildcard event. Such event will
	// be triggered on every other events.
	Events() map[string]Event
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

	// Make sure the destination returns a collection of events. Otherwise we can
	// not continue.
	if d.Events() == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Destination events must not be nil",
			Path:    []string{"Destination", d.String(), "Events()"},
		})

		return fail
	}

	// Validate every events of the destination and add the validation errors.
	for _, e := range d.Events() {
		errs := validateEvent(e)
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
