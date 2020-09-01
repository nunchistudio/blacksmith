package flow

import (
	"context"
)

/*
Defaults are the defaults options set for a flow. When not set, these values will
automatically be applied.
*/
var Defaults = &Options{
	Context: context.Background(),
	Enabled: false,
}

/*
Options is the options a user can pass to use a flow.
*/
type Options struct {

	// Context is a free key-value dictionary that will be passed to the flow.
	Context context.Context `json:"-"`

	// Enabled lets the user enable or disable a flow.
	Enabled bool `json:"enabled"`
}
