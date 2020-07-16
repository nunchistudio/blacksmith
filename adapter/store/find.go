package store

import (
	"time"
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
	Pagination *Pagination `json:"pagination"`

	// Where is the constraints applied to the query to find events, jobs, or
	// transitions. This is included in the meta because the store can set defaults
	// or override some constraints (such as a maximum limit). This allows to be aware
	// of the constraints actually applied to the query.
	Where *WhereEvents `json:"where"`
}

/*
Pagination holds the pagination details when looking for entries into the store.
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
WhereEvents is used to set constraints on the events when looking for entries into
the store.
*/
type WhereEvents struct {

	// SourcesIn makes sure the entries returned by the query have any of the source
	// name present in the slice.
	SourcesIn []string `json:"sources_in,omitempty"`

	// SourcesNotIn makes sure the entries returned by the query does not have any
	// of the source name present in the slice.
	SourcesNotIn []string `json:"sources_notin,omitempty"`

	// EventsIn makes sure the entries returned by the query have any of the source's
	// event name present in the slice.
	EventsIn []string `json:"events_in,omitempty"`

	// EventsNotIn makes sure the entries returned by the query does not have any
	// of the source's event name present in the slice.
	EventsNotIn []string `json:"events_notin,omitempty"`

	// CreatedBefore makes sure the entries returned by the query are related to an
	// event created before this instant.
	CreatedBefore *time.Time `json:"created_before,omitempty"`

	// CreatedAfter makes sure the entries returned by the query are related to an
	// event created after this instant.
	CreatedAfter *time.Time `json:"created_after,omitempty"`

	// AndWhereJobs lets you define additional constraints related to the jobs for
	// the entries you are looking for.
	AndWhereJobs *WhereJobs `json:"jobs,omitempty"`

	// Offset specifies the number of entries to skip before starting to return entries
	// from the query.
	Offset uint16 `json:"offset,omitempty"`

	// Limit specifies the number of entries to return after the offset clause has
	// been processed.
	Limit uint16 `json:"limit,omitempty"`
}

/*
WhereJobs is used to set constraints on jobs when looking for entries into the store.
*/
type WhereJobs struct {

	// EventID allows to find every entries related to a specific event ID.
	//
	// Note: When set, other constraints are not applied (except parent offset and
	// limit).
	EventID string `json:"event_id,omitempty"`

	// DestinationsIn makes sure the entries returned by the query have any of the
	// destination name present in the slice.
	DestinationsIn []string `json:"destinations_in,omitempty"`

	// DestinationsNotIn makes sure the entries returned by the query does not have any
	// of the destination name present in the slice.
	DestinationsNotIn []string `json:"destinations_notin,omitempty"`

	// EventsIn makes sure the entries returned by the query have any of the destination's
	// event name present in the slice.
	EventsIn []string `json:"events_in,omitempty"`

	// EventsNotIn makes sure the entries returned by the query does not have any of
	// the destination's event name present in the slice.
	EventsNotIn []string `json:"events_notin,omitempty"`

	// CreatedBefore makes sure the entries returned by the query are related to a
	// job created before this instant.
	CreatedBefore *time.Time `json:"created_before,omitempty"`

	// CreatedAfter makes sure the entries returned by the query are related to a
	// job created after this instant.
	CreatedAfter *time.Time `json:"created_after,omitempty"`

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
	JobID string `json:"job_id,omitempty"`

	// StatusIn makes sure the entries returned by the query have any of the status
	// present in the slice.
	StatusIn []string `json:"status_in,omitempty"`

	// StatusNotIn makes sure the entries returned by the query does not have any of
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
