package source

import (
	"context"
)

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
Options is the options a user can pass to use a source.
*/
type Options struct {

	// Load is used to load and use a source.
	Load Source `json:"-"`

	// Context is a free key-value dictionary that will be passed to the source.
	Context context.Context `json:"-"`

	// DefaultSchedule represents a schedule at which a source's trigger should run.
	// This value can be overridden by the source triggers to benefit a per trigger
	// basis schedule. This is used for CRON tasks so you can trigger jobs at a
	// given time.
	DefaultSchedule *Schedule `json:"cron"`
}
