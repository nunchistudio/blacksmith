package source

import (
	"context"
	"time"
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

	// Versions is a collection of supported versions for a source. The value of
	// each version is its deprecation date. It must be set to an empty time.Time
	// when the version is still maintained.
	//
	// When nil or empty, versioning is disabled for the source.
	//
	// Note: Feature only available in Blacksmith Enterprise Edition.
	Versions map[string]time.Time `json:"versions,omitempty"`

	// DefaultVersion is the default version to apply if the version is not set by
	// a consumer when executing a trigger. It must be one of the versions supported
	// in Versions.
	//
	// Note: Feature only available in Blacksmith Enterprise Edition.
	DefaultVersion string `json:"default_version,omitempty"`

	// DefaultSchedule represents a schedule at which a source's trigger should run.
	// This value can be overridden by the source triggers to benefit a per trigger
	// basis schedule. This is used for CRON tasks so you can trigger jobs at a
	// given time.
	DefaultSchedule *Schedule `json:"cron"`
}
