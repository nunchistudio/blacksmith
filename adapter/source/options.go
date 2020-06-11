package source

import (
	"context"

	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/internal/adapter"
)

/*
AvailableAdapters is a list of available source adapters.
*/
var AvailableAdapters = map[string]bool{}

/*
Defaults are the defaults options set for a destination. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{
	DefaultSchedule: &Schedule{
		Interval: "@every 1h",
	},
}

/*
Options is the options a user can pass to create a new source.
*/
type Options struct {

	// From can be used to download, install, and use an existing adapter. This way
	// the user does not need to develop a custom source adapter.
	From string `json:"from,omitempty"`

	// Load can be used to load and use a custom source adapter developed in-house.
	Load Source `json:"-"`

	// Context is a free key-value dictionary that will be passed to the underlying
	// adapter.
	Context context.Context `json:"-"`

	// DefaultSchedule represents a schedule at which a source's event should run.
	// This value can be overridden by the source events to benefit a per event basis
	// schedule. This is used for change-data-capture so you can listen for non HTTP
	// events such as database notifications.
	DefaultSchedule *Schedule `json:"cron"`
}

/*
Schedule represents a schedule at which a source's event should run. This is used
for change-data-capture so you can listen for non HTTP events such as database
notifications.
*/
type Schedule struct {

	// Interval represents an interval or a CRON string at which an event shall be
	// loaded from the source.
	Interval string `json:"interval"`
}

/*
ValidateAndLoad validates the source's options and returns a valid source interface.
*/
func (opts *Options) ValidateAndLoad() (Source, error) {
	var s Source
	var err error

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "source: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Set default options needed.
	if opts == nil {
		opts = Defaults
	}

	// Use the custom adapter if the user passed one. Otherwise, make sure the
	// source adapter is a valid one and load it from the Go plugin.
	if opts.Load != nil {
		s = opts.Load
	} else if opts.From != "" {
		if _, exists := AvailableAdapters[opts.From]; !exists {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Adapter not supported",
				Path:    []string{"Options", "Source", "From"},
			})

			return nil, fail
		}

		s, err = opts.loadPlugin()
		if err != nil {
			return nil, err
		}
	} else {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "One of 'From' or 'Load' must be set",
			Path:    []string{"Options", "Source"},
		})

		return nil, fail
	}

	// Validate the source adapter.
	err = validateSource(s)
	if err != nil {
		return nil, err
	}

	// We are now sure to be able to use the adapter.
	return s, nil
}

/*
loadPlugin loads a Go plugin using the adapter ID from the source options.
It returns the source interface loaded from the Go plugin.
*/
func (opts *Options) loadPlugin() (Source, error) {

	// Load the Go plugin's symbol from the helper.
	symbol, err := adapter.LoadPlugin(InterfaceSource, opts.From)
	if err != nil {
		return nil, err
	}

	// Convert the symbol to the desired type.
	converted := symbol.(func(*Options) (Source, error))

	// Load the Go plugin's source adapter.
	s, err := converted(opts)
	if err != nil {
		return nil, err
	}

	// Finally, return the source adapter.
	return s, nil
}
