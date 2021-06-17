package warehouse

import (
	"database/sql"
)

/*
Options is the options a source or destination can pass to be leveraged as a data
warehouse.
*/
type Options struct {

	// Name indicates the identifier of the SQL database which will be used as name
	// for the warehouse.
	//
	// Example: "sqlike(mypostgres)"
	// Required.
	Name string

	// DB is the database connection created using the package database/sql of the
	// standard library.
	//
	// Required.
	DB *sql.DB
}
