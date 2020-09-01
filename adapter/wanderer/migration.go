package wanderer

import (
	"time"
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
validDirections is used to make sure the direction passed is valid.
*/
var validDirections = map[string]bool{
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
	// Examples: "source", "source/trigger", "destination", "destination/action"
	InterfaceKind string `json:"interface_kind"`

	// InterfaceString is the string representation of the interface the migration
	// depends on.
	//
	// Examples: "crm", "crm/register", "postgres", "postgres/register"
	InterfaceString string `json:"interface_string"`

	// Directions contains the details about the up and down logic to run. Directions
	// must have two entries, where the key is either "up" or "down".
	Directions map[string]*Direction `json:"directions"`

	// CreatedAt is a timestamp of the migration creation date into the wanderer
	// datastore.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

/*
Direction contains the up or down details for a migration.
*/
type Direction struct {

	// ID is the unique identifier of the direction. It must be a valid KSUID.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	ID string `json:"id"`

	// Direction indicates if it is a "up" or "down" direction.
	Direction string `json:"direction"`

	// Filename is the name of the file containing the direction logic to be run.
	// It must contain the version suffixed by a dot, and the direction prefixed
	// by a dot.
	//
	// Example: "20060102150405.add_rbac.up.sql"
	Filename string `json:"filename"`

	// SHA256 is the SHA256 byte representation of the file's content. Once the
	// direction has already been migrated once, the SHA256 shall never be modified.
	SHA256 []byte `json:"sha256"`

	// Transitions is an array of the direction's transitions. It is used to keep
	// track of successes, failures, and errors so the wanderer is aware of the
	// direction's status.
	//
	// Note: It is up to the adapter to only return the latest direction's transition
	// since this is the only one that really matters in this context.
	Transitions [1]*Transition `json:"transitions"`

	// CreatedAt is a timestamp of the direction creation date into the wanderer
	// datastore.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// MigrationID is the ID of the migration related to this direction.
	MigrationID string `json:"migration_id"`
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

	// StateBefore is the state of the direction before running the migration. This
	// shall be nil when acknowledging the direction.
	StateBefore *string `json:"state_before"`

	// StateAfter is the state of the direction after running the new transition.
	StateAfter string `json:"state_after"`

	// Error keeps track of encountered if any.
	Error error `json:"error"`

	// CreatedAt is a timestamp of the transition creation date into the wanderer
	// datastore.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// MigrationID is the ID of the migration that is being run by the transition.
	MigrationID string `json:"migration_id"`

	// DirectionID is the ID of the direction that is being run by the transition.
	DirectionID string `json:"direction_id"`
}
