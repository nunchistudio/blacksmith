package supervisor

/*
Driver is a custom type allowing the user to only pass supported drivers when
configuring the supervisor adapter.
*/
type Driver string

/*
DriverPostgreSQL is used to leverage PostgreSQL as the supervisor adapter.
*/
var DriverPostgreSQL Driver = "postgres"

/*
DriverConsul is used to leverage Consul as the supervisor adapter.
*/
var DriverConsul Driver = "consul"

/*
Defaults are the defaults options set for the supervisor. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{}

/*
Options is the options a user can pass to use the supervisor adapter.
*/
type Options struct {

	// From is used to set the desired driver for the supervisor adapter.
	From Driver `json:"from,omitempty"`

	// Connection is the connection string to connect to the supervisor.
	Connection string `json:"-"`
}
