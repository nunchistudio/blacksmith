package warehouse

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nunchistudio/blacksmith/helper/errors"

	"github.com/flosch/pongo2/v4"
)

/*
Warehouse represents a data warehouse. End-users can run SQL template syntax which
compile to SQL, on top of the native dialect of the database.
*/
type Warehouse struct {

	// options are the options originally passed to the Options struct.
	options *Options
}

/*
AsWarehouse can be implemented by sources and destinations to benefit a template
SQL syntax on top of a SQL database for TLT.

Note: It is implemented by the third-party package sqlike to easily leverage the
standard database/sql and run operations / queries on top of the SQL database. See
Go module at https://pkg.go.dev/github.com/nunchistudio/blacksmith-modules/sqlike
for more details.
*/
type AsWarehouse interface {

	// AsWarehouse returns a data warehouse so end-user can run SQL operations and
	// queries on top the native dialect of the database.
	AsWarehouse() (*Warehouse, error)
}

/*
New returns a new data warehouse.
*/
func New(opts *Options) (*Warehouse, error) {
	wh := &Warehouse{
		options: opts,
	}

	return wh, nil
}

/*
Compile compiles a SQL template to a SQL query string.
*/
func (wh *Warehouse) Compile(filename string, data map[string]interface{}) (string, error) {
	fail := &errors.Error{
		Message:     fmt.Sprintf("warehouse/%s: Failed to compile SQL file", wh.options.Name),
		Validations: []errors.Validation{},
	}

	// Make sure we can get the working directory.
	// If an error occurred, we can not continue.
	wd, err := os.Getwd()
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return "", fail
	}

	// Create a template from the file.
	tmpl, err := pongo2.FromFile(filepath.Join(wd, filename))
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return "", fail
	}

	// Execute the template to get the compiled SQL query.
	out, err := tmpl.Execute(data)
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return "", fail
	}

	return out, nil
}

/*
Exec executes a SQL query string against the warehouse within a transaction.
*/
func (wh *Warehouse) Exec(query string) error {
	fail := &errors.Error{
		Message:     fmt.Sprintf("warehouse/%s: Failed to run operation", wh.options.Name),
		Validations: []errors.Validation{},
	}

	tx, err := wh.options.DB.Begin()
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return fail
	}

	defer tx.Rollback()

	_, err = tx.Exec(query)
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return fail
	}

	err = tx.Commit()
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return fail
	}

	return nil
}

/*
Query executes a SQL query against the warehouse, returning rows.
*/
func (wh *Warehouse) Query(query string) (*sql.Rows, error) {
	fail := &errors.Error{
		Message:     fmt.Sprintf("warehouse/%s: Failed to run query", wh.options.Name),
		Validations: []errors.Validation{},
	}

	rows, err := wh.options.DB.Query(query)
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return nil, fail
	}

	return rows, nil
}
