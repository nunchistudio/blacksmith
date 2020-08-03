package supervisor

import (
	"context"

	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/internal/adapter"
	"github.com/nunchistudio/blacksmith/version"
)

/*
AvailableAdapters is a list of available supervisors adapters.
*/
var AvailableAdapters = map[string]bool{
	"consul": true,
}

/*
Defaults are the defaults options set for the supervisor. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{
	Enabled: false,
	Join: &Node{
		Tags: []string{"blacksmith"},
		Meta: map[string]string{
			"go_version":         version.Go(),
			"blacksmith_version": version.Blacksmith(),
		},
	},
}

/*
Options is the options a user can pass to create a new supervisor.
*/
type Options struct {

	// From can be used to download, install, and use an existing adapter.
	From string `json:"from,omitempty"`

	// Load can be used to load and use a custom supervisor adapter developed in-house.
	Load Supervisor `json:"-"`

	// Context is a free key-value dictionary that will be passed to the underlying
	// adapter.
	Context context.Context `json:"-"`

	// Enabled allows the user to enable the Supervisor interface and this way
	// leverage the features of distributed system for high-availability and
	// better fault-tolerance.
	//
	// Note: This option will only apply when using the Enterprise Edition in both
	// the gateway and scheduler adapters.
	Enabled bool `json:"enabled"`

	// Join allows to attach the current instance to a node of the supervisor used.
	// Each instance shall be attached to a different node for distributed lock
	// mechanism.
	Join *Node `json:"node"`
}

/*
ValidateAndLoad validates the supervisor's options and returns a valid supervisor
interface.
*/
func (opts *Options) ValidateAndLoad() (Supervisor, error) {
	var s Supervisor
	var err error

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "supervisor: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Set default options needed.
	if opts == nil {
		opts = Defaults
	}

	// Do not use supervisor if the user does not want to. If this is the case, we can
	// stop here.
	if opts.Enabled == false {
		return nil, nil
	}

	// Use the custom adapter if the user passed one.
	if opts.Load != nil {
		return opts.Load, nil
	}

	// Use the custom adapter if the user passed one. Otherwise, make sure the
	// supervisor adapter is a valid one and load it from the Go plugin.
	// If the user didn't pass an adapter, return with no error since supervisor
	// is optional.
	if opts.Load != nil {
		s = opts.Load
	} else if opts.From != "" {
		if _, exists := AvailableAdapters[opts.From]; !exists {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Adapter not supported",
				Path:    []string{"Options", "Supervisor", "From"},
			})

			return nil, fail
		}

		s, err = opts.loadPlugin()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}

	// If the user didn't attach a node, we can not continue.
	if opts.Join == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Node must not be nil",
			Path:    []string{"Options", "Supervisor", "Join"},
		})

		return nil, fail
	}

	// Validate the supervisor adapter.
	err = validateSupervisor(s)
	if err != nil {
		return nil, err
	}

	// We are now sure to be able to use the adapter.
	return s, nil
}

/*
loadPlugin loads a Go plugin using the adapter ID from the supervisor options.
It returns the supervisor interface loaded from the Go plugin.
*/
func (opts *Options) loadPlugin() (Supervisor, error) {

	// Load the Go plugin's symbol from the helper.
	symbol, err := adapter.LoadPlugin(InterfaceSupervisor, opts.From)
	if err != nil {
		return nil, err
	}

	// Convert the symbol to the desired type.
	converted := symbol.(func(*Options) (Supervisor, error))

	// Load the Go plugin's supervisor adapter.
	s, err := converted(opts)
	if err != nil {
		return nil, err
	}

	// Finally, return the supervisor adapter.
	return s, nil
}
