package pubsub

import (
	"context"

	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/internal/adapter"
)

/*
AvailableAdapters is a list of available pubsub adapters.
*/
var AvailableAdapters = map[string]bool{
	"kafka":    true,
	"nats":     true,
	"rabbitmq": true,
}

/*
Defaults are the defaults options set for the pubsub. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{
	Enabled: false,
	Topic:   "blacksmith",
}

/*
Options is the options a user can pass to create a new pubsub.
*/
type Options struct {

	// From can be used to download, install, and use an existing adapter. This way
	// the user does not need to develop a custom pubsub adapter.
	From string

	// Load can be used to load and use a custom pubsub adapter developed in-house.
	Load PubSub

	// Context is a free key-value dictionary that will be passed to the underlying
	// adapter.
	Context context.Context

	// Enabled allows the user to enable the PubSub interface and this way distribute
	// jobs to destinations in realtime. If disabled, the scheduler will load jobs
	// to destinations given the schedule of each destination and event.
	Enabled bool

	// Connection is the connection string to connect to the pubsub.
	Connection string

	// Topic is the topic name the pubsub adapter will use to publish and subscribe
	// messages to.
	Topic string
}

/*
ValidateAndLoad validates the pubsub's options and returns a valid pubsub interface.
*/
func (opts *Options) ValidateAndLoad() (PubSub, error) {
	var ps PubSub
	var err error

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "pubsub: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Set default options needed.
	if opts == nil {
		opts = Defaults
	}

	// Do not use pubsub if the user does not want to. If this is the case, we can
	// stop here.
	if opts.Enabled == false {
		return nil, nil
	}

	// Use the custom adapter if the user passed one.
	if opts.Load != nil {
		return opts.Load, nil
	}

	// Use the custom adapter if the user passed one. Otherwise, make sure the
	// pubsub adapter is a valid one and load it from the Go plugin.
	// If the user didn't pass an adapter, return with no error since pubsub
	// is optional.
	if opts.Load != nil {
		ps = opts.Load
	} else if opts.From != "" {
		if _, exists := AvailableAdapters[opts.From]; !exists {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Adapter not supported",
				Path:    []string{"Options", "PubSub", "From"},
			})

			return nil, fail
		}

		ps, err = opts.loadPlugin()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}

	// If the user didn't put a topic, use the default one.
	if opts.Topic == "" {
		opts.Topic = Defaults.Topic
	}

	// Validate the pubsub adapter.
	err = validatePubSub(ps)
	if err != nil {
		return nil, err
	}

	// We are now sure to be able to use the adapter.
	return ps, nil
}

/*
loadPlugin loads a Go plugin using the adapter ID from the pubsub options.
It returns the pubsub interface loaded from the Go plugin.
*/
func (opts *Options) loadPlugin() (PubSub, error) {

	// Load the Go plugin's symbol from the helper.
	symbol, err := adapter.LoadPlugin(opts.Context, "pubsub", opts.From)
	if err != nil {
		return nil, err
	}

	// Convert the symbol to the desired type.
	converted := symbol.(func(*Options) (PubSub, error))

	// Load the Go plugin's pubsub adapter.
	ps, err := converted(opts)
	if err != nil {
		return nil, err
	}

	// Finally, return the pubsub adapter.
	return ps, nil
}
