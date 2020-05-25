package store

import (
	"context"

	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/internal/adapter"
)

/*
AvailableAdapters is a list of available store adapters.
*/
var AvailableAdapters = map[string]bool{
	"postgres": true,
}

/*
Defaults are the defaults options set for the store. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{}

/*
Options is the options a user can pass to create a new store.
*/
type Options struct {

	// From can be used to download, install, and use an existing adapter. This way
	// the user does not need to develop a custom store adapter.
	From string

	// Load can be used to load and use a custom store adapter developed in-house.
	Load Store

	// Context is a free key-value dictionary that will be passed to the underlying
	// adapter.
	Context context.Context

	// Connection is the connection string to connect to the store.
	Connection string
}

/*
ValidateAndLoad validates the store's options and returns a valid store interface.
*/
func (opts *Options) ValidateAndLoad() (Store, error) {
	var s Store
	var err error

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "store: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Set default options needed.
	if opts == nil {
		opts = Defaults
	}

	// Use the custom adapter if the user passed one. Otherwise, make sure the
	// store adapter is a valid one and load it from the Go plugin.
	if opts.Load != nil {
		s = opts.Load
	} else if opts.From != "" {
		if _, exists := AvailableAdapters[opts.From]; !exists {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Adapter not supported",
				Path:    []string{"Options", "Store", "From"},
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
			Path:    []string{"Options", "Store"},
		})

		return nil, fail
	}

	// Validate the store adapter.
	err = validateStore(s)
	if err != nil {
		return nil, err
	}

	// We are now sure to be able to use the adapter.
	return s, nil
}

/*
loadPlugin loads a Go plugin using the adapter ID from the store options.
It returns the store interface loaded from the Go plugin.
*/
func (opts *Options) loadPlugin() (Store, error) {

	// Load the Go plugin's symbol from the helper.
	symbol, err := adapter.LoadPlugin(opts.Context, "store", opts.From)
	if err != nil {
		return nil, err
	}

	// Convert the symbol to the desired type.
	converted := symbol.(func(*Options) (Store, error))

	// Load the Go plugin's store adapter.
	s, err := converted(opts)
	if err != nil {
		return nil, err
	}

	// Finally, return the store adapter.
	return s, nil
}
