package wanderer

import (
	"time"
)

/*
Meta includes information about the query's result returned by the wanderer when
looking for entries (migrations, directions, or transitions).
*/
type Meta struct {

	// Count is the number of entries found that match the constraints applied to
	// the query (without the limit).
	Count uint16 `json:"count"`

	// Pagination is the pagination details based on the count, offset, and limit.
	Pagination *Pagination `json:"pagination"`

	// Where is the constraints applied to the query to find migrations, directions,
	// or transitions. This is included in the meta because the wanderer can set
	// defaults or override some constraints (such as a maximum limit). This allows
	// to be aware of the constraints actually applied to the query.
	Where *WhereMigrations `json:"where"`
}

/*
Pagination holds the pagination details when looking for entries into the wanderer.
*/
type Pagination struct {

	// Current is the current page.
	Current uint16 `json:"current"`

	// Previous is the previous page. It will be nil if there is no previous page.
	Previous *uint16 `json:"previous"`

	// Next is the next page. It will be nil if there is no next page.
	Next *uint16 `json:"next"`

	// First is the first page. It will always be 1.
	First uint16 `json:"first"`

	// Last is the last page.
	Last uint16 `json:"last"`
}

/*
WhereMigrations is used to set constraints on the migrations when looking for
entries into the wanderer.
*/
type WhereMigrations struct {

	// InterfaceKindsIn makes sure the entries returned by the query is related to
	// any of the interface kind present in the slice.
	InterfaceKindsIn []string `json:"interface_kinds_in,omitempty"`

	// InterfaceKindsNotIn makes sure the entries returned by the query is not related
	// to any of the interface kind present in the slice.
	InterfaceKindsNotIn []string `json:"interface_kinds_notin,omitempty"`

	// InterfaceStringsIn makes sure the entries returned by the query is related to
	// any of the interface string present in the slice.
	InterfaceStringsIn []string `json:"interface_strings_in,omitempty"`

	// InterfaceStringsNotIn makes sure the entries returned by the query is not related
	// to any of the interface string present in the slice.
	InterfaceStringsNotIn []string `json:"interface_strings_notin,omitempty"`

	// CreatedBefore makes sure the entries returned by the query are related to a
	// migration created before this instant.
	CreatedBefore *time.Time `json:"created_before,omitempty"`

	// CreatedAfter makes sure the entries returned by the query are related to a
	// migration created after this instant.
	CreatedAfter *time.Time `json:"created_after,omitempty"`

	// AndWhereDirections lets you define additional constraints related to the
	// directions for the entries you are looking for.
	AndWhereDirections *WhereDirections `json:"directions,omitempty"`

	// Offset specifies the number of entries to skip before starting to return entries
	// from the query.
	Offset uint16 `json:"offset,omitempty"`

	// Limit specifies the number of entries to return after the offset clause has
	// been processed.
	Limit uint16 `json:"limit,omitempty"`
}

/*
WhereDirections is used to set constraints on directions when looking for entries
into the wanderer.
*/
type WhereDirections struct {

	// MigrationID allows to find every entries related to a specific migration ID.
	//
	// Note: When set, other constraints are not applied (except parent offset and
	// limit).
	MigrationID string `json:"migration_id,omitempty"`

	// DirectionsIn makes sure the entries returned by the query have any of the
	// direction present in the slice.
	DirectionsIn []string `json:"directions_in,omitempty"`

	// DirectionsNotIn makes sure the entries returned by the query do not have any
	// of the direction present in the slice.
	DirectionsNotIn []string `json:"directions_notin,omitempty"`

	// CreatedBefore makes sure the entries returned by the query are related to a
	// direction created before this instant.
	CreatedBefore *time.Time `json:"created_before,omitempty"`

	// CreatedAfter makes sure the entries returned by the query are related to a
	// direction created after this instant.
	CreatedAfter *time.Time `json:"created_after,omitempty"`

	// AndWhereTransitions lets you define additional constraints related to the
	// transitions for the entries you are looking for.
	AndWhereTransitions *WhereTransitions `json:"transitions,omitempty"`
}

/*
WhereTransitions is used to set constraints on transitions when looking for entries
into the wanderer.
*/
type WhereTransitions struct {

	// DirectionID allows to find every entries related to a specific direction ID.
	//
	// Note: When set, other constraints are not applied (except parent offset and
	// limit).
	DirectionID string `json:"direction_id,omitempty"`

	// StatusIn makes sure the entries returned by the query have any of the status
	// present in the slice.
	StatusIn []string `json:"status_in,omitempty"`

	// StatusNotIn makes sure the entries returned by the query do not have any of
	// the status present in the slice.
	StatusNotIn []string `json:"status_notin,omitempty"`

	// MinAttempts makes sure the entries returned by the query have equal to or greater
	// than this number of attempts.
	MinAttempts uint16 `json:"min_attempts,omitempty"`

	// MaxAttempts makes sure the entries returned by the query have equal to or lesser
	// than this number of attempts.
	MaxAttempts uint16 `json:"max_attempts,omitempty"`

	// CreatedBefore makes sure the entries returned by the query are related to a
	// transition created before this instant.
	CreatedBefore *time.Time `json:"created_before,omitempty"`

	// CreatedAfter makes sure the entries returned by the query are related to a
	// transition created after this instant.
	CreatedAfter *time.Time `json:"created_after,omitempty"`
}
