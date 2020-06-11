package destination

import (
	"context"

	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/internal/adapter"
)

/*
AvailableAdapters is a list of available destination adapters.
*/
var AvailableAdapters = map[string]bool{}

/*
Defaults are the defaults options set for a destination. When not set, these values
will automatically be applied.

We set a hourly interval for 3 days so it give time to teams to be aware of the
failures and debug the destination if needed.
*/
var Defaults = &Options{
	DefaultSchedule: &Schedule{
		Realtime:   false,
		Interval:   "@every 1h",
		MaxRetries: 72,
	},
}

/*
Options is the options a user can pass to create a new destination.
*/
type Options struct {

	// From can be used to download, install, and use an existing adapter. This way
	// the user does not need to develop a custom destination adapter.
	From string `json:"from,omitempty"`

	// Load can be used to load and use a custom destination adapter developed
	// in-house.
	Load Destination `json:"-"`

	// Context is a free key-value dictionary that will be passed to the underlying
	// adapter.
	Context context.Context `json:"-"`

	// DefaultSchedule represents a schedule at which a destination's event should
	// run. This value can be overridden by the underlying adapter if necessary so
	// the user does not make any scheduling mistake. This value can also be overridden
	// by the destination events to benefit a per event basis schedule.
	DefaultSchedule *Schedule `json:"schedule"`
}

/*
Schedule represents a schedule at which a destination's event should run. SaaS APIs
could be used in realtime whereas data warehouses shall be used only a few times
per day.
*/
type Schedule struct {

	// Realtime indicates if the pubsub adapter of the Blacksmith application shall
	// be used to load events to the destination in realtime or not.
	Realtime bool `json:"realtime"`

	// Interval represents an interval or a CRON string at which an event shall be
	// loaded to the destination.
	Interval string `json:"interval"`

	// MaxRetries indicates the maximum number of retries per job the scheduler will
	// attempt to execute for each job. When the limit is reached, the job is marked
	// as "discarded".
	MaxRetries uint16 `json:"max_retries"`
}

/*
ValidateAndLoad validates the destination's options and returns a valid destination
interface.
*/
func (opts *Options) ValidateAndLoad() (Destination, error) {
	var d Destination
	var err error

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "destination: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Set default options needed.
	if opts == nil {
		opts = Defaults
	}

	// Use the custom adapter if the user passed one. Otherwise, make sure the
	// destination adapter is a valid one and load it from the Go plugin.
	if opts.Load != nil {
		d = opts.Load
	} else if opts.From != "" {
		if _, exists := AvailableAdapters[opts.From]; !exists {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Adapter not supported",
				Path:    []string{"Options", "Destination", "From"},
			})

			return nil, fail
		}

		d, err = opts.loadPlugin()
		if err != nil {
			return nil, err
		}
	} else {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "One of 'From' or 'Load' must be set",
			Path:    []string{"Options", "Destination"},
		})

		return nil, fail
	}

	// Validate the destination adapter.
	err = validateDestination(d)
	if err != nil {
		return nil, err
	}

	// We are now sure to be able to use the adapter.
	return d, nil
}

/*
loadPlugin loads a Go plugin using the adapter ID from the destination options.
It returns the destination interface loaded from the Go plugin.
*/
func (opts *Options) loadPlugin() (Destination, error) {

	// Load the Go plugin's symbol from the helper.
	symbol, err := adapter.LoadPlugin(InterfaceDestination, opts.From)
	if err != nil {
		return nil, err
	}

	// Convert the symbol to the desired type.
	converted := symbol.(func(*Options) (Destination, error))

	// Load the Go plugin's destination adapter.
	d, err := converted(opts)
	if err != nil {
		return nil, err
	}

	// Finally, return the destination adapter.
	return d, nil
}
