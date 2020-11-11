package wanderer

import (
	"time"

	"github.com/nunchistudio/blacksmith/helper/rest"
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
	Pagination *rest.Pagination `json:"pagination"`

	// Where is the constraints applied to the query to find migrations or transitions.
	// This is included in the meta because the wanderer can set defaults or override
	// some constraints (such as a maximum limit). This allows to be aware of the
	// constraints actually applied to the query.
	Where *WhereMigrations `json:"where"`
}

/*
WhereMigrations is used to set constraints on the migrations when looking for
entries into the wanderer.
*/
type WhereMigrations struct {

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
	Offset uint16 `json:"offset"`

	// Limit specifies the number of entries to return after the offset clause has
	// been processed.
	Limit uint16 `json:"limit"`
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
	MigrationID string `json:"migration.id,omitempty"`

	// StatusIn makes sure the entries returned by the query have any of the status
	// present in the slice.
	StatusIn []string `json:"status_in,omitempty"`

	// StatusNotIn makes sure the entries returned by the query do not have any of
	// the status present in the slice.
	StatusNotIn []string `json:"status_notin,omitempty"`
}
