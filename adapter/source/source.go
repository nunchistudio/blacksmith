package source

import (
	"fmt"

	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
InterfaceSource is the string representation for the source interface.
*/
var InterfaceSource = "source"

/*
Source is the interface used to load events from cloud services, databases, or
any kind of application able to send data or push notifications.
*/
type Source interface {

	// String returns the string representation of the adapter.
	//
	// Example: "stripe"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// Triggers returns a list of triggers the source is able to take care of.
	// Events will then be forwared to related destinations' actions.
	Triggers() map[string]Trigger
}

/*
validateSource makes sure a source adapter is ready to be used properly by a
Blacksmith application.
*/
func validateSource(s Source) error {

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "source: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Verify the source ID is not empty.
	if s.String() == "" {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Source ID must not be empty",
			Path:    []string{"Source", "unknown", "String()"},
		})

		return fail
	}

	// We now can add the adapter name to the error message.
	fail.Message = fmt.Sprintf("source/%s: Failed to load adapter", s.String())

	// It is impossible to deal with nil options.
	if s.Options() == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Source options must not be nil",
			Path:    []string{"Source", s.String(), "Options()"},
		})

		return fail
	}

	// Verify the source default scheduling options.
	if s.Options().DefaultSchedule != nil {
		if s.Options().DefaultSchedule.Interval == "" {
			s.Options().DefaultSchedule.Interval = Defaults.DefaultSchedule.Interval
		}
	} else {
		s.Options().DefaultSchedule = Defaults.DefaultSchedule
	}

	// Make sure the source returns a collection of events. Otherwise we can not
	// continue.
	if s.Triggers() == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Source events must not be nil",
			Path:    []string{"Source", s.String(), "Triggers()"},
		})

		return fail
	}

	// Validate every events of the source and add the validation errors.
	for _, t := range s.Triggers() {
		errs := validateTrigger(s.String(), t)
		if errs != nil {
			fail.Validations = append(fail.Validations, errs...)
		}
	}

	// Avoid cycles.
	s.Options().Load = nil

	// Return the error if any occurred.
	if len(fail.Validations) > 0 {
		return fail
	}

	return nil
}
