package store

/*
Driver is a custom type allowing the user to only pass supported drivers when
configuring the store adapter.
*/
type Driver string

/*
DriverPostgreSQL is used to leverage PostgreSQL as the store adapter.
*/
var DriverPostgreSQL Driver = "postgres"

/*
Defaults are the defaults options set for the store. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{
	PurgePolicies: []*PurgePolicy{},
}

/*
Options is the options a user can pass to use the store adapter.
*/
type Options struct {

	// From is used to set the desired driver for the store adapter.
	From Driver `json:"from,omitempty"`

	// Connection is the connection string to connect to the store.
	Connection string `json:"-"`

	// PurgePolicies allows to define several intervals to purge entries in the
	// store given advanced constraints.
	PurgePolicies []*PurgePolicy `json:"purge"`
}

/*
PurgePolicy is a policy to purge the store. It will run given the interval and
will be applied given the constraints.
*/
type PurgePolicy struct {

	// Define the constraints to retrieve the entries to purge from store.
	//
	// Note: Offset, limit, and pagination will not be applied.
	WhereEvents *WhereEvents `json:"where"`

	// Interval represents an interval or a CRON string at which events, jobs, and
	// transitions shall be purged from the store.
	Interval string `json:"interval"`
}
