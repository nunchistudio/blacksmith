package wanderer

import (
	"time"

	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
StatusAcknowledged is used to mark a migration as acknowledged. This is used
when registering new migrations into the wanderer.
*/
var StatusAcknowledged = "acknowledged"

/*
StatusAwaiting is used to mark a migration as awaiting. This is used when a
migration is awaiting to be run.
*/
var StatusAwaiting = "awaiting"

/*
StatusExecuting is used to mark a migration as executing. This is used when a
migration is being executed.
*/
var StatusExecuting = "executing"

/*
StatusSucceeded is used to mark a migration as succeeded.
*/
var StatusSucceeded = "succeeded"

/*
StatusFailed is used to mark a migration as failed.
*/
var StatusFailed = "failed"

/*
StatusRollbacked is used to mark a migration as rollbacked.
*/
var StatusRollbacked = "rollbacked"

/*
ValidDirections is used to make sure the direction passed is valid.
*/
var ValidDirections = map[string]bool{
	"up":   true,
	"down": true,
}

/*
Migration contains the details about a specific migration, including its up and
down details.
*/
type Migration struct {

	// ID is the unique identifier of the migration. It must be a valid KSUID.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	ID string `json:"id"`

	// Version is a 14 character string representation of the timestamp when the
	// migration has been created. The name of the version is of the form
	// YYYYMMDDHHMMSS, which is UTC timestamp identifying the migration.
	//
	// Example: "20060102150405"
	Version string `json:"version"`

	// InterfaceKind is the string representation of the interface kind the migration
	// depends on.
	//
	// Example: "destination/event"
	InterfaceKind string `json:"interface_kind"`

	// InterfaceString is the string representation of the interface string the
	// migration depends on.
	//
	// Example: "postgres/identify"
	InterfaceString string `json:"interface_string"`

	// Run contains the details about the up and down logic to run. Run must have
	// at most 2 entries, corresponding to the up and down directions.
	Run map[string]*Direction `json:"run"`

	// CreatedAt is a timestamp of the migration creation date into the wanderer
	// datastore.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

/*
Direction contains the up and down details of a migration.
*/
type Direction struct {

	// ID is the unique identifier of the direction. It must be a valid KSUID.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	ID string `json:"id"`

	// Filename is the name of the file containing the direction logic to be run.
	// It must contain the version suffixed by a dot, and the direction prefixed
	// by a dot.
	//
	// Example: "20060102150405.add_rbac.up.sql"
	Filename string `json:"filename"`

	// SHA256 is the SHA256 byte representation of the file's content. Once the
	// direction has already been migrated once, the SHA256 shall never be modified.
	SHA256 []byte `json:"sha256"`

	// CreatedAt is a timestamp of the direction creation date into the wanderer
	// datastore.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Transitions is an array of the direction's transitions. It is used to keep
	// track of successes, failures, and errors so the wanderer is aware of the
	// migration's status. It is up to the adapter to only return the latest
	// migration's transition since this is the only one that really matters.
	Transitions [1]*Transition `json:"transitions"`
}

/*
Transition is used to keep track of the status of migration's direction.
*/
type Transition struct {

	// ID is the unique identifier of the transition. It must be a valid KSUID.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	ID string `json:"id"`

	// Attempt represents the number of tentatives that the direction has run
	// before succeeded.
	Attempt uint16 `json:"attempt"`

	// StateBefore is the state of the direction before running the new transition.
	// See status details for more info.
	StateBefore *string `json:"state_before"`

	// StateAfter is the state of the direction after running the new transition.
	// See status details for more info.
	StateAfter string `json:"state_after"`

	// Error keeps track of encountered if any.
	Error error `json:"error"`

	// CreatedAt is a timestamp of the transition creation date into the wanderer
	// datastore.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// LockID is the ID of the lock used to run the migration. This is here for
	// convenience and should not be included in results if used in an API.
	LockID string `json:"-"`

	// MigrationID is the ID of the migration that is being run by the transition.
	// This is here for convenience and should not be included in results if used
	// in an API.
	MigrationID string `json:"-"`

	// DirectionID is the ID of the direction that is being run by the transition.
	// This is here for convenience and should not be included in results if used
	// in an API.
	DirectionID string `json:"-"`
}

/*
WhereMigration is used to find migrations in the wanderer.
*/
type WhereMigration struct {

	// Direction is the string representation of the direction we are looking for.
	// Must be one of "up" or "down".
	Direction string `json:"direction"`

	// StatusIn contains the status where the migrations status match at least one
	// element in the list.
	StatusIn []string `json:"status_in"`

	// StatusNotIn contains the status where the migrations status must not match
	// at least one element in the list.
	StatusNotIn []string `json:"status_notin"`

	// Version is the minimum or maximum value for the version for a down or a up
	// direction.
	Version string `json:"version"`
}

/*
Validate validates the filters passed as params to find specific migrations.
*/
func (where *WhereMigration) Validate() error {
	fail := &errors.Error{
		Message:     "wanderer: Failed to validate 'where' filters",
		Validations: []errors.Validation{},
	}

	if where == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "where must not be nil",
			Path:    []string{"where"},
		})

		return fail
	}

	if where.Direction == "" {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Direction must not be empty",
			Path:    []string{"where", "Direction"},
		})
	}

	if _, exists := ValidDirections[where.Direction]; !exists {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Direction must be one of 'up', 'down'",
			Path:    []string{"where", "Direction"},
		})
	}

	if where.StatusIn == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "StatusIn must not be nil",
			Path:    []string{"where", "StatusIn"},
		})
	}

	if where.StatusNotIn == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "StatusNotIn must not be nil",
			Path:    []string{"where", "StatusNotIn"},
		})
	}

	if where.Version == "" {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Version must not be empty",
			Path:    []string{"where", "Version"},
		})
	} else if len(where.Version) != 14 {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Version must be 14 character long",
			Path:    []string{"where", "Version"},
		})
	}

	if len(fail.Validations) > 0 {
		return fail
	}

	return nil
}
