package wanderer

import (
	"context"

	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/internal/adapter"
)

/*
AvailableAdapters is a list of available wanderer adapters.
*/
var AvailableAdapters = map[string]bool{
	"postgres": true,
}

/*
Defaults are the defaults options set for the wanderer. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{}

/*
Options is the options a user can pass to create a new wanderer.
*/
type Options struct {

	// From can be used to download, install, and use an existing adapter. This way
	// the user does not need to develop a custom wanderer adapter.
	From string

	// Load can be used to load and use a custom wanderer adapter developed in-house.
	Load Wanderer

	// Context is a free key-value dictionary that will be passed to the underlying
	// adapter.
	Context context.Context

	// Connection is the connection string to connect to the wanderer.
	Connection string
}

/*
ValidateAndLoad validates the wanderer's options and returns a valid wanderer
interface.
*/
func (opts *Options) ValidateAndLoad() (Wanderer, error) {
	var w Wanderer
	var err error

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "wanderer: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Set default options needed.
	if opts == nil {
		opts = Defaults
	}

	// Use the custom adapter if the user passed one. Otherwise, make sure the
	// wanderer adapter is a valid one and load it from the Go plugin.
	// If the user didn't pass an adapter, return with no error since the wanderer
	// is optional.
	if opts.Load != nil {
		w = opts.Load
	} else if opts.From != "" {
		if _, exists := AvailableAdapters[opts.From]; !exists {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Adapter not supported",
				Path:    []string{"Options", "Wanderer", "From"},
			})

			return nil, fail
		}

		w, err = opts.loadPlugin()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}

	// Validate the wanderer adapter.
	err = validateWanderer(w)
	if err != nil {
		return nil, err
	}

	// We are now sure to be able to use the adapter.
	return w, nil
}

/*
loadPlugin loads a Go plugin using the adapter ID from the wanderer options.
It returns the wanderer interface loaded from the Go plugin.
*/
func (opts *Options) loadPlugin() (Wanderer, error) {

	// Load the Go plugin's symbol from the helper.
	symbol, err := adapter.LoadPlugin(InterfaceWanderer, opts.From)
	if err != nil {
		return nil, err
	}

	// Convert the symbol to the desired type.
	converted := symbol.(func(*Options) (Wanderer, error))

	// Load the Go plugin's wanderer adapter.
	w, err := converted(opts)
	if err != nil {
		return nil, err
	}

	// Finally, return the wanderer adapter.
	return w, nil
}
