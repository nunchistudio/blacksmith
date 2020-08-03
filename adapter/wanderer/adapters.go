package wanderer

/*
WithMigrate can be implemented by adapters — such as sources and destinations —
to benefit custom data and database schema migrations.

Note: Feature only available in Blacksmith Enterprise Edition.
*/
type WithMigrate interface {

	// Migrate is the migration logic for running every migrations of an adapter.
	// The function gives access to the migration being run at the moment.
	Migrate(*Toolkit, *Migration) error
}

/*
WithMigrations must be implemented by adapters — such as sources and destinations —
already implementing the WithMigrate interface. It allows to have migrations
for each adapter.

It is separated from WithMigrate because each sources' trigger and destinations'
action can have its specific migrations isolated from its parent adapter.

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
