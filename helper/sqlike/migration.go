package sqlike

import (
	"database/sql"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/nunchistudio/blacksmith/adapter/wanderer"
	"github.com/nunchistudio/blacksmith/helper/errors"

	"github.com/segmentio/ksuid"
)

/*
LoadMigrations loads migrations files from a directory.
*/
func LoadMigrations(directory string) ([]*wanderer.Migration, error) {
	fail := &errors.Error{
		Message:     "sqlike: Failed to load migration files",
		Validations: []errors.Validation{},
	}

	// Make sure we can get the working directory.
	// If an error occurred, we can not continue.
	wd, err := os.Getwd()
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return nil, fail
	}

	// Try to open the target directory.
	// If an error occurred, we can not continue.
	f, err := os.Open(filepath.Join(wd, directory))
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
			Path:    strings.Split(directory, "/"),
		})

		return nil, fail
	}

	// Get the file list from the directory.
	// If an error occurred, we can not continue.
	list, err := f.Readdir(-1)
	defer f.Close()
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
			Path:    strings.Split(directory, "/"),
		})

		return nil, fail
	}

	// Sort the files by name.
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name() < list[j].Name()
	})

	// We'll keep track of each file and make sure each migration has its up and
	// down file.
	migrations := []*wanderer.Migration{}
	registered := map[string]*wanderer.Migration{}

	// Go through each migration file.
	for _, file := range list {

		// Make sure we can deal with the file.
		filename := strings.Split(file.Name(), ".")
		if len(filename) != 4 {
			continue
		} else if len(filename[0]) != 14 {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Version number must be formatted like YYYYMMDDHHMISS",
				Path:    append(strings.Split(directory, "/"), file.Name()),
			})
		} else if filename[2] != "up" && filename[2] != "down" {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Migration must either be 'up' or 'down'",
				Path:    append(strings.Split(directory, "/"), file.Name()),
			})
		} else if filename[3] != "sql" {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "File extension not supported (must be '.sql')",
				Path:    append(strings.Split(directory, "/"), file.Name()),
			})
		}

		// Open the desired file so we can then use it.
		f, err := os.Open(filepath.Join(wd, directory, file.Name()))
		defer f.Close()
		if err != nil {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: err.Error(),
				Path:    append(strings.Split(directory, "/"), file.Name()),
			})
		}

		// Retrieve the version name from the file name.
		number := file.Name()[0:14]
		if len(number) != 14 {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Failed to parse version name",
				Path:    append(strings.Split(directory, "/"), file.Name()),
			})
		}

		// Convert the stringified number version to a valid Go time.Time.
		numbert, err := time.Parse("20060102150405", number)
		if err != nil {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Failed to parse version name",
				Path:    append(strings.Split(directory, "/"), file.Name()),
			})
		}

		// Return now if anything bad happened.
		if len(fail.Validations) > 0 {
			return nil, fail
		}

		// Add the migration if it does not already exist. Retrieve the name from
		// the file name.
		if _, exists := registered[number]; !exists {
			registered[number] = &wanderer.Migration{
				ID:      ksuid.New().String(),
				Version: numbert,
				Name:    filename[1],
			}
		}
	}

	// Create a slice of known migrations.
	for _, r := range registered {
		migrations = append(migrations, r)
	}

	// Finally, return the migration files.
	return migrations, nil
}

/*
RunMigration runs a SQL migration within a transaction using the standard
database/sql package.
*/
func RunMigration(db *sql.DB, directory string, migration *wanderer.Migration) error {
	fail := &errors.Error{
		Message:     "sqlike: Failed to run migration file",
		Validations: []errors.Validation{},
	}

	// Make sure we can get the working directory.
	// If an error occurred, we can not continue.
	wd, err := os.Getwd()
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return fail
	}

	// Try to open the file given the migration details.
	filename := migration.Version.Format("20060102150405") + "." + migration.Name + "." + migration.Direction + ".sql"
	file, err := os.Open(filepath.Join(wd, directory, filename))
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
			Path:    strings.Split(directory, "/"),
		})

		return fail
	}

	// Close the file when done.
	defer file.Close()

	// Read the file so we will then be able to run its content.
	buf, err := io.ReadAll(file)
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
			Path:    strings.Split(filepath.Join(wd, directory, filename), "/"),
		})

		return fail
	}

	// Save the content of the file.
	query := string(buf[:])

	// Start the SQL transaction.
	txn, err := db.Begin()
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
			Path:    strings.Split(filepath.Join(wd, directory, filename), "/"),
		})

		return fail
	}

	// Execute the SQL transaction.
	_, err = txn.Exec(query)
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
			Path:    strings.Split(filepath.Join(wd, directory, filename), "/"),
		})

		return fail
	}

	// Finally commit it. If an error occured, rollback it.
	err = txn.Commit()
	if err != nil {
		txn.Rollback()

		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
			Path:    strings.Split(filepath.Join(wd, directory, filename), "/"),
		})

		return fail
	}

	// If we made it here then no error occured.
	return nil
}
