package scheduler

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/internal/adapter"
)

/*
AvailableAdapters is a list of available scheduler adapters.
*/
var AvailableAdapters = map[string]bool{
	"standard": true,
}

/*
Defaults are the defaults options set for the scheduler. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{
	From:    "standard",
	Address: ":8081",
}

/*
Options is the options a user can pass to create a new scheduler.
*/
type Options struct {

	// From can be used to download, install, and use an existing adapter. This way
	// the user does not need to develop a custom scheduler adapter.
	From string

	// Load can be used to load and use a custom scheduler adapter developed in-house.
	Load Scheduler

	// Context is a free key-value dictionary that will be passed to the underlying
	// adapter.
	Context context.Context

	// Address is the HTTP address the scheduler server is listening to.
	//
	// Defaults to ":8081".
	Address string

	// TLS is the TLS settings used to run the TLS server.
	TLS *tls.Config

	// CertFile is the relative path to the certificate file for the TLS server.
	CertFile string

	// KeyFile is the relative path to the key file for the TLS server.
	KeyFile string

	// Middleware is the HTTP middleware chain that will be applied to the HTTP server
	// of the scheduler.
	Middleware func(http.Handler) http.Handler

	// Attach allows you to attach an external HTTP handler to the Blacksmith scheduler.
	// It is useful for adding HTTP routes with custom routing and business logic.
	//
	// If a handler is attached, all routes within this handler will be prefixed with
	// a prefix chosen by the scheduler adapter.
	Attach http.Handler
}

/*
ValidateAndLoad validates the scheduler's options and returns a valid scheduler
interface.
*/
func (opts *Options) ValidateAndLoad() (Scheduler, error) {
	var s Scheduler
	var err error

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "scheduler: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Set default options needed.
	if opts == nil {
		opts = Defaults
	}

	// Use the custom adapter if the user passed one. Otherwise, make sure the
	// scheduler adapter is a valid one and load it from the Go plugin.
	if opts.Load != nil {
		s = opts.Load
	} else {
		if opts.From == "" {
			opts.From = Defaults.From
		}

		if _, exists := AvailableAdapters[opts.From]; !exists {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Adapter not supported",
				Path:    []string{"Options", "Scheduler", "From"},
			})

			return nil, fail
		}

		s, err = opts.loadPlugin()
		if err != nil {
			return nil, err
		}
	}

	// If the user didn't put an address, use the default one.
	if opts.Address == "" {
		opts.Address = Defaults.Address
	}

	// Validate the scheduler adapter.
	err = validateScheduler(s)
	if err != nil {
		return nil, err
	}

	// We are now sure to be able to use the adapter.
	return s, nil
}

/*
loadPlugin loads a Go plugin using the adapter ID from the scheduler options.
It returns the scheduler interface loaded from the Go plugin.
*/
func (opts *Options) loadPlugin() (Scheduler, error) {

	// Load the Go plugin's symbol from the helper.
	symbol, err := adapter.LoadPlugin(opts.Context, "scheduler", opts.From)
	if err != nil {
		return nil, err
	}

	// Convert the symbol to the desired type.
	converted := symbol.(func(*Options) (Scheduler, error))

	// Load the Go plugin's scheduler adapter.
	s, err := converted(opts)
	if err != nil {
		return nil, err
	}

	// Finally, return the scheduler adapter.
	return s, nil
}
