package wanderer

import (
	"time"
)

/*
Meta includes information about the query's result returned by the wanderer when
looking for entries (migrations or transitions).
*/
type Meta struct {

	// Count is the number of entries found that match the constraints applied to
	// the query (without the limit).
	Count uint16 `json:"count"`

	// Pagination is the pagination details based on the count, offset, and limit.
	Pagination *Pagination `json:"pagination"`

	// Where is the constraints applied to the query to find migrations or transitions.
	// This is included in the meta because the wanderer can set defaults or override
	// some constraints (such as a maximum limit). This allows to be aware of the
	// constraints actually applied to the query.
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

	// Name allows to search for a migration by its name.
	Name string `json:"name,omitempty"`

	// VersionBefore makes sure the entries returned by the query are related to a
	// migration versioned before this instant.
	VersionedBefore *time.Time `json:"versioned_before,omitempty"`

	// VersionAfter makes sure the entries returned by the query are related to a
	// migration versioned after this instant.
	VersionedAfter *time.Time `json:"versioned_after,omitempty"`

	// ScopeIn makes sure the entries returned by the query is related to any of
	// the scope kind present in the slice.
	ScopeIn []string `json:"scope_in,omitempty"`

	// ScopeNotIn makes sure the entries returned by the query is not related to
	// any of the scope present in the slice.
	ScopeNotIn []string `json:"scope_notin,omitempty"`

	// AndWhereTransitions lets you define additional constraints related to the
	// transitions for the migrations you are looking for.
	AndWhereTransitions *WhereTransitions `json:"transitions,omitempty"`

	// Offset specifies the number of entries to skip before starting to return entries
	// from the query.
	Offset uint16 `json:"offset,omitempty"`

	// Limit specifies the number of entries to return after the offset clause has
	// been processed.
	Limit uint16 `json:"limit,omitempty"`
}

/*
WhereTransitions is used to set constraints on transitions when looking for entries
into the wanderer.
*/
type WhereTransitions struct {

	// MigrationID allows to find every entries related to a specific migration ID.
	//
	// Note: When set, other constraints are not applied (except parent offset and
	// limit).
	MigrationID string `json:"migration_id,omitempty"`

	// StatusIn makes sure the entries returned by the query have any of the status
	// present in the slice.
	StatusIn []string `json:"status_in,omitempty"`

	// StatusNotIn makes sure the entries returned by the query do not have any of
	// the status present in the slice.
	StatusNotIn []string `json:"status_notin,omitempty"`

	// CreatedBefore makes sure the entries returned by the query are related to a
	// transition created before this instant.
	CreatedBefore *time.Time `json:"created_before,omitempty"`

	// CreatedAfter makes sure the entries returned by the query are related to a
	// transition created after this instant.
	CreatedAfter *time.Time `json:"created_after,omitempty"`
}
