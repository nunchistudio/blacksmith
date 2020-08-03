package source

import (
	"time"

	"github.com/nunchistudio/blacksmith/adapter/destination"
	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
InterfaceTrigger is the string representation for the source's trigger interface.
*/
var InterfaceTrigger = "source/trigger"

/*
Trigger represents a trigger of a source adapter. A Trigger contains all the logic
to handle a specific event for a source.
*/
type Trigger interface {

	// String returns the string representation of the source's trigger.
	//
	// Example: "identify"
	String() string

	// Mode indicates the trigger mode to use along some options to execute the
	// source's trigger. The gateway will trigger the Marshal function based on
	// these options.
	Mode() *Mode
}

/*
Mode indicates how a source's trigger is triggered.
*/
type Mode struct {

	// Mode indicates the trigger mode to trigger the event.
	//
	// - When set to ModeHTTP, the UsingHTTP route is used as the trigger.
	// - When set to ModeCRON, the UsingCRON schedule is used as the trigger.
	// - When set to ModeCDC, no trigger is registered since it is an ongoing
	//   listener. It is up to the Marshal function to include the infinite loop
	//   and return the payload using the channel included in the toolkit.
	Mode string `json:"mode"`

	// UsingHTTP defines the HTTP route the event will react to.
	UsingHTTP *Route `json:"http,omitempty"`

	// UsingCRON represents a schedule at which an event should run. When returning
	// nil, the source's schedule is applied.
	UsingCRON *Schedule `json:"cron,omitempty"`
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

	// Jobs defines the transformation logic for the event from the source trigger
	// to destinations actions. Each destination's action can benefit its own
	// transformation.
	//
	// Example:
	//
	//   Jobs: map[string][]destination.Action{
	//     "crm": []destination.Action{
	//       crm.NewCustomer{
	//         Name:           e.Data.FirstName + " " + e.Data.LastName,
	//         EmailAddresses: []string{ e.Data.Email },
	//       },
	//     },
	//   },
	Jobs map[string][]destination.Action

	// SentAt allows you to keep track of the timestamp when the event was originally
	// sent.
	SentAt *time.Time `json:"sent_at,omitempty"`
}

/*
validateTrigger makes sure a trigger of a source adapter is ready to be used properly
by a Blacksmith application.
*/
func validateTrigger(s string, t Trigger) []errors.Validation {
	validations := []errors.Validation{}

	// Verify the trigger ID is not empty.
	if t.String() == "" {
		validations = append(validations, errors.Validation{
			Message: "Trigger ID must not be empty",
			Path:    []string{"Source", s, "Triggers()", "unknown", "String()"},
		})
	}

	// Make sure the trigger mode is not empty. If so, we can not continue.
	if t.Mode() == nil {
		validations = append(validations, errors.Validation{
			Message: "Trigger mode must not be nil",
			Path:    []string{"Source", s, "Triggers()", t.String(), "Mode()"},
		})

		return validations
	}

	// Depending on the trigger mode, validate the Marshal function.
	switch t.Mode().Mode {
	case ModeCRON:
		_, ok := t.(TriggerCRON)
		if ok == false {
			validations = append(validations, errors.Validation{
				Message: "Marshal function must be of type: func(*source.Toolkit) (*source.Payload, error)",
				Path:    []string{"Source", s, "Triggers()", t.String(), "Marshal()"},
			})
		}

	case ModeCDC:
		_, ok := t.(TriggerCDC)
		if ok == false {
			validations = append(validations, errors.Validation{
				Message: "Marshal function must be of type: func(*source.Toolkit, *source.Notifier)",
				Path:    []string{"Source", s, "Triggers()", t.String(), "Marshal()"},
			})
		}

	case ModeHTTP:
		_, ok := t.(TriggerHTTP)
		if ok == false {
			validations = append(validations, errors.Validation{
				Message: "Marshal function must be of type: func(*source.Toolkit, *http.Request) (*source.Payload, error)",
				Path:    []string{"Source", s, "Triggers()", t.String(), "Marshal()"},
			})
		}

		validations = append(validations, t.Mode().UsingHTTP.validate(s, t.String())...)

	default:
		validations = append(validations, errors.Validation{
			Message: "Trigger mode must be one of 'cron', 'http', 'cdc'",
			Path:    []string{"Source", s, "Triggers()", t.String(), "Mode()", "Mode"},
		})

		return validations
	}

	// Return the validation errors if any occurred.
	if len(validations) > 0 {
		return validations
	}

	return nil
}
