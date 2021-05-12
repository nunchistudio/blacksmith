package wanderer

/*
Driver is a custom type allowing the user to only pass supported drivers when
configuring the wanderer adapter.
*/
type Driver string

/*
DriverPostgreSQL is used to leverage PostgreSQL as the wanderer adapter.
*/
var DriverPostgreSQL Driver = "postgres"

/*
Defaults are the defaults options set for the wanderer. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{}

/*
Options is the options a user can pass to use the wanderer adapter.
*/
type Options struct {

	// From is used to set the desired driver for the wanderer adapter.
	From Driver `json:"from,omitempty"`

	// Connection is the connection string to connect to the wanderer.
	Connection string `json:"-"`
}
