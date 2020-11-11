package store

import (
	"time"

	"github.com/nunchistudio/blacksmith/helper/rest"
)

/*
Meta includes information about the query's result returned by the store when
looking for entries (events, jobs, or transitions).
*/
type Meta struct {

	// Count is the number of entries found that match the constraints applied to
	// the query (without the limit).
	Count uint16 `json:"count"`

	// Pagination is the pagination details based on the count, offset, and limit.
	Pagination *rest.Pagination `json:"pagination"`

	// Where is the constraints applied to the query to find events, jobs, or
	// transitions. This is included in the meta because the store can set defaults
	// or override some constraints (such as a maximum limit). This allows to be aware
	// of the constraints actually applied to the query.
	Where *WhereEvents `json:"where"`
}

/*
WhereEvents is used to set constraints on the events when looking for entries into
the store.
*/
type WhereEvents struct {

	// SourcesIn makes sure the entries returned by the query have any of the source
	// name present in the slice.
	SourcesIn []string `json:"events.sources_in,omitempty"`

	// SourcesNotIn makes sure the entries returned by the query do not have any
	// of the source name present in the slice.
	SourcesNotIn []string `json:"events.sources_notin,omitempty"`

	// TriggersIn makes sure the entries returned by the query have any of the source's
	// trigger name present in the slice.
	TriggersIn []string `json:"events.triggers_in,omitempty"`

	// TriggersNotIn makes sure the entries returned by the query do not have any
	// of the source's trigger name present in the slice.
	TriggersNotIn []string `json:"events.triggers_notin,omitempty"`

	// VersionsIn makes sure the entries returned by the query have any of the source's
	// version present in the slice.
	VersionsIn []string `json:"events.versions_in,omitempty"`

	// VersionsNotIn makes sure the entries returned by the query do not have any
	// of the source's version present in the slice.
	VersionsNotIn []string `json:"events.versions_notin,omitempty"`

	// CreatedBefore makes sure the entries returned by the query are related to an
	// event created before this instant.
	CreatedBefore *time.Time `json:"events.created_before,omitempty"`

	// CreatedAfter makes sure the entries returned by the query are related to an
	// event created after this instant.
	CreatedAfter *time.Time `json:"events.created_after,omitempty"`

	// AndWhereJobs lets you define additional constraints related to the jobs for
	// the entries you are looking for.
	AndWhereJobs *WhereJobs `json:"jobs,omitempty"`

	// Offset specifies the number of entries to skip before starting to return entries
	// from the query.
	Offset uint16 `json:"offset"`

	// Limit specifies the number of entries to return after the offset clause has
	// been processed.
	Limit uint16 `json:"limit"`
}

/*
WhereJobs is used to set constraints on jobs when looking for entries into the store.
*/
type WhereJobs struct {

	// EventID allows to find every entries related to a specific event ID.
	//
	// Note: When set, other constraints are not applied (except parent offset and
	// limit).
	EventID string `json:"event.id,omitempty"`

	// DestinationsIn makes sure the entries returned by the query have any of the
	// destination name present in the slice.
	DestinationsIn []string `json:"jobs.destinations_in,omitempty"`

	// DestinationsNotIn makes sure the entries returned by the query do not have any
	// of the destination name present in the slice.
	DestinationsNotIn []string `json:"jobs.destinations_notin,omitempty"`

	// ActionsIn makes sure the entries returned by the query have any of the destination's
	// action name present in the slice.
	ActionsIn []string `json:"jobs.actions_in,omitempty"`

	// ActionsNotIn makes sure the entries returned by the query do not have any of
	// the destination's action name present in the slice.
	ActionsNotIn []string `json:"jobs.actions_notin,omitempty"`

	// VersionsIn makes sure the entries returned by the query have any of the
	// destination's version present in the slice.
	VersionsIn []string `json:"jobs.versions_in,omitempty"`

	// VersionsNotIn makes sure the entries returned by the query do not have any of
	// the destination's version present in the slice.
	VersionsNotIn []string `json:"jobs.versions_notin,omitempty"`

	// CreatedBefore makes sure the entries returned by the query are related to a
	// job created before this instant.
	CreatedBefore *time.Time `json:"jobs.created_before,omitempty"`

	// CreatedAfter makes sure the entries returned by the query are related to a
	// job created after this instant.
	CreatedAfter *time.Time `json:"jobs.created_after,omitempty"`

	// AndWhereTransitions lets you define additional constraints related to the
	// transitions for the entries you are looking for.
	AndWhereTransitions *WhereTransitions `json:"transitions,omitempty"`
}

/*
WhereTransitions is used to set constraints on transitions when looking for entries
into the store.
*/
type WhereTransitions struct {

	// JobID allows to find every entries related to a specific job ID.
	//
	// Note: When set, other constraints are not applied (except parent offset and
	// limit).
	JobID string `json:"job.id,omitempty"`

	// StatusIn makes sure the entries returned by the query have any of the status
	// present in the slice.
	StatusIn []string `json:"jobs.status_in,omitempty"`

	// StatusNotIn makes sure the entries returned by the query do not have any of
	// the status present in the slice.
	StatusNotIn []string `json:"jobs.status_notin,omitempty"`

	// MinAttempts makes sure the entries returned by the query have equal to or greater
	// than this number of attempts.
	MinAttempts uint16 `json:"jobs.min_attempts,omitempty"`

	// MaxAttempts makes sure the entries returned by the query have equal to or lesser
	// than this number of attempts.
	MaxAttempts uint16 `json:"jobs.max_attempts,omitempty"`
}
