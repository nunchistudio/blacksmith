package destination

import (
	"context"
	"time"
)

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
Options is the options a user can pass to use a destination.
*/
type Options struct {

	// Load is used to load and use a destination.
	Load Destination `json:"-"`

	// Context is a free key-value dictionary that will be passed to the destination.
	Context context.Context `json:"-"`

	// Versions is a collection of supported versions for a destination. The value
	// of each version is its deprecation date. It must be set to an empty time.Time
	// when the version is still maintained.
	//
	// When nil or empty, versioning is disabled for the destination.
	//
	// Note: Feature only available in Blacksmith Enterprise Edition.
	Versions map[string]time.Time `json:"versions,omitempty"`

	// DefaultVersion is the default version to apply if the version is not set by
	// a flow when executing an action. It must be one of the versions supported in
	// Versions.
	//
	// Note: Feature only available in Blacksmith Enterprise Edition.
	DefaultVersion string `json:"default_version,omitempty"`

	// DefaultSchedule represents a schedule at which a destination's action should
	// run. This value can be overridden by the underlying destination if necessary
	// so the user does not make any scheduling mistake. This value can also be
	// overridden by each destination action to benefit a per action basis schedule.
	DefaultSchedule *Schedule `json:"schedule"`
}

/*
Schedule represents a schedule at which a destination's action should run. SaaS
APIs could be used in realtime whereas data warehouses shall be used only a few
times per day.
*/
type Schedule struct {

	// Realtime indicates if the pubsub adapter of the Blacksmith application shall
	// be used to load events to the destination in realtime or not. When false, the
	// Interval will be used.
	Realtime bool `json:"realtime"`

	// Interval represents an interval or a CRON string at which an event shall be
	// loaded to the destination. It is used as the time-lapse between retries in
	// case of a job failure.
	Interval string `json:"interval"`

	// MaxRetries indicates the maximum number of retries per job the scheduler will
	// attempt to execute for each job. When the limit is reached, the job is marked
	// as "discarded".
	MaxRetries uint16 `json:"max_retries"`
}
