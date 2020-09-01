package wanderer

/*
InterfaceWanderer is the string representation for the wanderer interface.
*/
var InterfaceWanderer = "wanderer"

/*
Wanderer is the interface used to persist the migrations in a datastore to keep
track of migrations states.
*/
type Wanderer interface {

	// String returns the string representation of the adapter.
	//
	// Example: "postgres"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// AddMigrations inserts a list of migrations into the wanderer given the data
	// passed in params. It returns an error if any occurred.
	AddMigrations(*Toolkit, []*Migration) error

	// FindMigration returns a migration given the migration ID passed in params.
	FindMigration(*Toolkit, string) (*Migration, error)

	// FindMigrations returns a list of migrations matching the constraints passed
	// in params. It also returns meta information about the query, such as pagination
	// and the constraints really applied to it.
	FindMigrations(*Toolkit, *WhereMigrations) ([]*Migration, *Meta, error)

	// AddDirections inserts a list of directions into the datastore.
	AddDirections(*Toolkit, []*Direction) error

	// FindDirection returns a direction given the direction ID passed in params.
	FindDirection(*Toolkit, string) (*Direction, error)

	// FindDirections returns a list of directions matching the constraints passed
	// in params. It also returns meta information about the query, such as pagination
	// and the constraints really applied to it.
	FindDirections(*Toolkit, *WhereMigrations) ([]*Direction, *Meta, error)

	// AddTransitions inserts a list of transitions into the datastore to update
	// their related direction status. We insert new transitions instead of updating
	// the direction itself to keep track of the migration direction's history.
	AddTransitions(*Toolkit, []*Transition) error

	// FindTransition returns a transition given the transition ID passed in params.
	FindTransition(*Toolkit, string) (*Transition, error)

	// FindTransitions returns a list of transitions matching the constraints passed
	// in params. It also returns meta information about the query, such as pagination
	// and the constraints really applied to it.
	FindTransitions(*Toolkit, *WhereMigrations) ([]*Transition, *Meta, error)
}

/*
WithMigrate can be implemented by sources and destinations to benefit custom data
and database schema migrations.

Note: Feature only available in Blacksmith Enterprise Edition.
*/
type WithMigrate interface {

	// Migrate is the migration logic for running every migrations for a source or
	// a destination. The function gives access only to the migrations that need to
	// run with the appropriate direction "up" or "down".
	Migrate(*Toolkit, []*Migration) error
}

/*
WithMigrations must be implemented by sources (and / or by their triggers) and
destinations (and / or by their actions) already implementing the WithMigrate
interface.

Note: Feature only available in Blacksmith Enterprise Edition.
*/
type WithMigrations interface {

	// Migrations returns a slice of migrations regardless their status. The wanderer
	// will then be able to process and keep track of each and every one of them.
	//
	// Note: The adapter can use the package helper/sqlike to easily read migrations
	// files from a directory. See package helper/sqlike for more details.
	Migrations() ([]*Migration, error)
}
