package source

import (
	"strings"
	"time"

	"github.com/nunchistudio/blacksmith/adapter/destination"
	"github.com/nunchistudio/blacksmith/adapter/wanderer"
	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
InterfaceEvent is the string representation for the source's event interface.
*/
var InterfaceEvent = "source/event"

/*
ModeHTTP is used to indicate the event is trigeered from a HTTP request.
*/
var ModeHTTP = "http"

/*
ModeCRON is used to indicate the event is trigeered from a CRON task.
*/
var ModeCRON = "cron"

/*
ModeForever is used to indicate the event is a forever loop. It is used for
ongoing listeners like database notifications.
*/
var ModeForever = "forever"

/*
Event represents an event of a source adapter. An Event contains all the logic to
handle a specific event for a source.
*/
type Event interface {

	// String returns the string representation of the source's event.
	//
	// Example: "identify"
	String() string

	// Migrations returns the list of all migrations for the event, regardless their
	// status.
	//
	// The adapter can use the package helper/sqlike to easily read migrations files
	// from a directory. See package helper/sqlike for more details.
	//
	// Note: Feature only available in Blacksmith Enterprise.
	Migrations() ([]*wanderer.Migration, error)

	// Trigger indicates the trigger mode to use along some options to execute the
	// source's event. The gateway will trigger the Marshal function based on these
	// options.
	Trigger() *Trigger

	// Marshal in charge of the "E" in the ETL process: it Extracts the data from
	// the source given the Trigger declared.
	//
	// The function allows to return data to destinations events as jobs. It is part
	// of the "T" in the ETL process: it transforms the payload from the source's
	// event to destinations events.
	//
	// When done the function can either return a payload or write to the channel
	// using the toolkit if it is event using the "forever" mode.
	//
	// Note: If the context in the returned payload is nil, the one from the
	// event will automatically be applied.
	//
	// Note: If the function returns an error, the event can not be considered as
	// extracted. Therefore, no events and no jobs will be created and the event
	// will never run.
	Marshal(*Toolkit) (*Payload, error)
}

/*
Trigger indicates how to run an event. It can executed from a HTTP request, from
a CRON task, or using an ongoing listener.
*/
type Trigger struct {

	// Mode indicates the trigger mode to execute the event.
	//
	// - When set to ModeHTTP, the OnHTTP route is used as the trigger.
	// - When set to ModeCRON, the OnCRON schedule is used as the trigger.
	// - When set to ModeForever, no trigger is registered since it is an ongoing
	//   listener. It is up to the Marshal function to include the infinite loop
	//   and return the payload using the channel included in the toolkit.
	Mode string

	// OnHTTP defines the HTTP route the event will react to. This must be nil if
	// the source is specifically designed for change-data-capture using an appropriate
	// schedule.
	OnHTTP *Route

	// OnCRON represents a schedule at which an event should run. When returning
	// nil, the source's schedule is applied. This is used for change-data-capture
	// so you can listen for non-HTTP events, like scheduled tasks.
	OnCRON *Schedule
}

/*
Payload represents the fields an event must fill. It will be used across all adapters
to match the fields between sources and destinations.
*/
type Payload struct {

	// Context is a dictionary of information that provides useful context about an
	// event. The context should be used inside every events for consistency.
	//
	// It must be a valid JSON since it will be used using json Marshal and Unmarshal
	// functions.
	Context []byte `json:"context"`

	// Data is the byte representation of the data sent by the event.
	//
	// It must be a valid JSON since it will be used using json Marshal and Unmarshal
	// functions.
	Data []byte `json:"data"`

	// Jobs defines the transformation logic for the event from source to destinations.
	// Each destination's event can benefit its own transformation.
	//
	// Example:
	//
	//   Jobs: map[string][]destination.Event{
	//     "crm": []destination.Event{
	//       crm.NewCustomer{
	//         Name:           e.Data.FirstName + " " + e.Data.LastName,
	//         EmailAddresses: []string{ e.Data.Email },
	//       },
	//     },
	//   },
	Jobs map[string][]destination.Event

	// SentAt allows you to keep track of the timestamp when the event was originally
	// sent.
	SentAt *time.Time `json:"sent_at,omitempty"`
}

/*
validateEvent makes sure an event of a source adapter is ready to be used properly
by a Blacksmith application.
*/
func validateEvent(e Event) []errors.Validation {
	validations := []errors.Validation{}

	// Verify the event ID is not empty.
	if e.String() == "" {
		validations = append(validations, errors.Validation{
			Message: "Event ID must not be empty",
			Path:    []string{"Event", "unknown", "String()"},
		})
	}

	// Make sure the trigger mode is not empty. If so, we can not continue.
	if e.Trigger() == nil {
		validations = append(validations, errors.Validation{
			Message: "Event trigegr must not be nil",
			Path:    []string{"Event", e.String(), "Trigger()"},
		})

		return validations
	}

	// Depending on the event mode, we validate either the schedule or the route
	// option. Note that, for now, we do not validate any details in the schedule
	// function because, when not provided, the source's schedule details will be
	// applied.
	switch e.Trigger().Mode {
	case ModeCRON:
	case ModeForever:

	case ModeHTTP:
		if e.Trigger().OnHTTP == nil {
			validations = append(validations, errors.Validation{
				Message: "Event route must not be nil",
				Path:    []string{"Event", e.String(), "Trigger()", "OnHTTP"},
			})

			return validations
		}

		if e.Trigger().OnHTTP.Path == "" {
			validations = append(validations, errors.Validation{
				Message: "HTTP route path must not be empty",
				Path:    []string{"Event", e.String(), "Trigger()", "OnHTTP", "Path"},
			})
		} else if e.Trigger().OnHTTP.Path == "/" {
			validations = append(validations, errors.Validation{
				Message: "HTTP route path must not be a wildcard path",
				Path:    []string{"Event", e.String(), "Trigger()", "OnHTTP", "Path"},
			})
		} else if strings.HasPrefix(e.Trigger().OnHTTP.Path, "/") == false {
			validations = append(validations, errors.Validation{
				Message: "HTTP route path must start with '/'",
				Path:    []string{"Event", e.String(), "Trigger()", "OnHTTP", "Path"},
			})
		} else if strings.HasPrefix(e.Trigger().OnHTTP.Path, "/api") == true {
			validations = append(validations, errors.Validation{
				Message: "HTTP route path must not start with '/api'",
				Path:    []string{"Event", e.String(), "Trigger()", "OnHTTP", "Path"},
			})
		}

		if e.Trigger().OnHTTP.Methods == nil {
			validations = append(validations, errors.Validation{
				Message: "HTTP route methods must not be nil",
				Path:    []string{"Event", e.String(), "Trigger()", "OnHTTP", "Methods"},
			})

			return validations
		}

		for _, method := range e.Trigger().OnHTTP.Methods {
			method = strings.ToUpper(method)
			if _, valid := ValidMethods[method]; !valid {
				validations = append(validations, errors.Validation{
					Message: "HTTP route method not valid",
					Path:    []string{"Event", e.String(), "Trigger()", "OnHTTP", "Methods", method},
				})
			}
		}

	default:
		validations = append(validations, errors.Validation{
			Message: "Trigger mode must be one of 'cron', 'http', 'forever'",
			Path:    []string{"Event", e.String(), "Trigger()", "Mode"},
		})

		return validations
	}

	// Return the validation errors if any occurred.
	if len(validations) > 0 {
		return validations
	}

	return nil
}
