package wanderer

import (
	"context"
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
var Defaults = &Options{
	Context: context.Background(),
}

/*
Options is the options a user can pass to use the wanderer adapter.
*/
type Options struct {

	// From is used to set the desired wanderer adapter. It must be one of
	// AvailableAdapters.
	From string `json:"from,omitempty"`

	// Context is a free key-value dictionary that will be passed to the underlying
	// adapter.
	Context context.Context `json:"-"`

	// Connection is the connection string to connect to the wanderer.
	Connection string `json:"-"`
}
