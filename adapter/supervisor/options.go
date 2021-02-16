package supervisor

import (
	"context"

	"github.com/nunchistudio/blacksmith/version"
)

/*
AvailableAdapters is a list of available supervisors adapters.
*/
var AvailableAdapters = map[string]bool{
	"consul":   true,
	"postgres": true,
}

/*
Defaults are the defaults options set for the supervisor. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{
	Context: context.Background(),
	Join: &Node{
		Tags: []string{"blacksmith"},
		Meta: map[string]string{
			"go_version":         version.Go(),
			"blacksmith_version": version.Blacksmith(),
		},
	},
}

/*
Options is the options a user can pass to use the supervisor adapter.
*/
type Options struct {

	// From is used to set the desired supervisor adapter. It must be one of
	// AvailableAdapters.
	From string `json:"from,omitempty"`

	// Context is a free key-value dictionary that will be passed to the underlying
	// adapter.
	Context context.Context `json:"-"`

	// Join allows to attach the current instance to a node of the supervisor used.
	// Each instance shall be attached to a different node for distributed lock
	// mechanism.
	Join *Node `json:"node"`
}
