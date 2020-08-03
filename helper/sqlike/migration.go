package sqlike

import (
	"crypto/sha256"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/nunchistudio/blacksmith/adapter/wanderer"
	"github.com/nunchistudio/blacksmith/helper/errors"

	"github.com/segmentio/ksuid"
)

/*
LoadMigrationFiles loads migrations files from a directory.
*/
func LoadMigrationFiles(directory string) ([]*wanderer.Migration, error) {
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
			Path:    []string{"Directory", directory},
		})

		return nil, fail
	}

	// Try to open the target directory.
	// If an error occurred, we can not continue.
	f, err := os.Open(filepath.Join(wd, directory))
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
			Path:    []string{"Directory", directory},
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
			Path:    []string{"Directory", directory},
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
		var direction string

		// Make sure the migration file is valid. If not, continue to the next one.
		if strings.Contains(file.Name(), ".up.") {
			direction = "up"
		} else if strings.Contains(file.Name(), ".down.") {
			direction = "down"
		} else {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Not a proper migration file name",
				Path:    []string{"File", directory, file.Name()},
			})

			continue
		}

		// Open the desired file so we can then use it.
		f, err := os.Open(wd + directory + "/" + file.Name())
		defer f.Close()
		if err != nil {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: err.Error(),
				Path:    []string{"File", directory, file.Name()},
			})
		}

		// Store the SHA256. This might be useful for the wanderer adapter to keep
		// track of file changes.
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: err.Error(),
				Path:    []string{"File", directory, file.Name()},
			})
		}

		// Retrieve the version name from the file name.
		version := file.Name()[0:14]
		_, err = strconv.ParseUint(version, 10, 32)
		if err != nil || len(version) != 14 {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Failed to parse version name",
				Path:    []string{"File", directory, file.Name()},
			})
		}

		// Add the migration if it does not already exist. A migration must have a
		// up and down files.
		if _, exists := registered[version]; !exists {
			registered[version] = &wanderer.Migration{
				ID:      ksuid.New().String(),
				Version: version,
				Run:     map[string]*wanderer.Direction{},
			}
		}

		// Add the direction to the migration.
		registered[version].Run[direction] = &wanderer.Direction{
			Filename: file.Name(),
			SHA256:   h.Sum(nil),
		}
	}

	// Create a slice of known migrations.
	for _, r := range registered {
		migrations = append(migrations, r)
	}

	// No need to keep track of validation errors if it is empty.
	if len(fail.Validations) == 0 {
		fail = nil
	}

	// Finally, return the migration files and the error if any occurred.
	return migrations, fail
}
