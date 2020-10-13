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
StatusExecutingUp is used to inform a migration is executing its "up" logic.
*/
var StatusExecutingUp = "executing(up)"

/*
StatusExecutingDown is used to inform a migration is executing its "down" logic.
*/
var StatusExecutingDown = "executing(down)"

/*
StatusSucceededUp is used to inform a migration has succeeded its "up" logic.
*/
var StatusSucceededUp = "succeeded(up)"

/*
StatusSucceededDown is used to inform a migration has succeeded its "down" logic.
In other words, the migration has successfully been rollbacked.
*/
var StatusSucceededDown = "succeeded(down)"

/*
StatusFailedUp is used to inform a migration has failed to run its "up" logic.
*/
var StatusFailedUp = "failed(up)"

/*
StatusFailedDown is used to inform a migration has failed to run its "up" logic.
*/
var StatusFailedDown = "failed(down)"

/*
StatusDiscarded is used to mark a migration as discarded. If a rollbacked migration
is marked as discarded it can not be run again.
*/
var StatusDiscarded = "discarded"

/*
Migration contains the details about a specific migration, including its up and
down details.
*/
type Migration struct {

	// ID is the unique identifier of the migration. It must be a valid KSUID.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	ID string `json:"id"`

	// Version is the 14 character timestamp when the migration has been created.
	// The name of the version is of the form YYYYMMDDHHMISS, which is a UTC
	// timestamp identifying the migration.
	//
	// Example: "20060102150405"
	Version time.Time `json:"version"`

	// Scope is the string representation of the scope of the migration.
	//
	// Examples: "source:crm", "source/trigger:crm/register"
	Scope string `json:"scope"`

	// Name is the name of the migration to run or rollback. This is a part of the
	// migration's filename.
	//
	// Example: "add_rbac"
	Name string `json:"name"`

	// Direction indicates which direction is being run at the moment when running
	// or rolling back migrations. It shall be one of "up" or "down".
	Direction string `json:"-"`

	// Transitions is an array of the migration's transitions. It is used to keep
	// track of successes, failures, and errors so the wanderer is aware of the
	// migration's status.
	//
	// Note: It is up to the adapter to only return the latest migration's transition
	// since this is the only one that really matters in this context.
	Transitions [1]*Transition `json:"transitions"`

	// CreatedAt is a timestamp of the migration creation date into the wanderer
	// datastore.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

/*
Transition is used to keep track of the status of migrations.
*/
type Transition struct {

	// ID is the unique identifier of the transition. It must be a valid KSUID.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	ID string `json:"id"`

	// StateBefore is the state of the migration before running it. This shall be
	// nil when acknowledging the migration.
	StateBefore *string `json:"state_before"`

	// StateAfter is the state of the migration after running the new transition.
	StateAfter string `json:"state_after"`

	// Error keeps track of encountered if any.
	Error error `json:"error"`

	// CreatedAt is a timestamp of the transition creation date into the wanderer
	// datastore.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// MigrationID is the ID of the migration that is being run by the transition.
	MigrationID string `json:"migration_id"`
}
