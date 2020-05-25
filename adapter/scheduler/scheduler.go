package scheduler

import (
	"fmt"
	"net/http"

	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
InterfaceScheduler is the string representation for the scheduler interface.
*/
var InterfaceScheduler = "scheduler"

/*
Scheduler is the interface used to handle jobs in an asynchronous way, using a
HTTP server.
*/
type Scheduler interface {

	// String returns the string representation of the adapter.
	//
	// Example: "standard"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// Handler returns a net/http Handler allowing the use of the scheduler as a
	// standard HTTP handler in an external Go application.
	Handler(*Toolkit) (http.Handler, error)

	// ListenAndServe starts the HTTP server. This is the equivalent of the net/http
	// ListenAndServe and function except it is wrapped so you can subscribe to events
	// asynchronously using the pubsub Subscriber and handle them across the different
	// adapters such as the store and destinations.
	ListenAndServe(*Toolkit, *WithTLS) error

	// Shutdown gracefully shuts down the server without interrupting any active
	// connections. It is the equivalent of the net/http Shutdown function.
	Shutdown(*Toolkit) error
}

/*
WithTLS allows you to attach TLS certificate files when creating the HTTP server.
*/
type WithTLS struct {

	// CertFile is the relative path to the certificate file.
	CertFile string

	// KeyFile is the relative path to the key file.
	KeyFile string
}

/*
validateScheduler makes sure the scheduler adapter is ready to be used properly
by a Blacksmith application.
*/
func validateScheduler(s Scheduler) error {

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "scheduler: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Verify the scheduler ID is not empty.
	if s.String() == "" {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Scheduler ID must not be empty",
			Path:    []string{"Scheduler", "unknown", "String()"},
		})

		return fail
	}

	// We now can add the adapter name to the error message.
	fail.Message = fmt.Sprintf("scheduler/%s: Failed to load adapter", s.String())

	// It is impossible to deal with nil options.
	if s.Options() == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Scheduler options must not be nil",
			Path:    []string{"Scheduler", s.String(), "Options()"},
		})

		return fail
	}

	// If the adapter didn't set an address, use the default one.
	if s.Options().Address == "" {
		s.Options().Address = Defaults.Address
	}

	// Avoid cycles.
	s.Options().Load = nil

	// Return the error if any occurred.
	if len(fail.Validations) > 0 {
		return fail
	}

	return nil
}
