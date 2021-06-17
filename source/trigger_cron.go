package source

/*
ModeCRON is used to indicate the event is triggered from a CRON task.
*/
var ModeCRON = "cron"

/*
TriggerCRON is the interface used for triggers using a CRON logic.

A new CRON trigger can be generated using the Blacksmith CLI:

  $ blacksmith generate trigger --name <name> --mode cron [--path <path>] [--migrations]
*/
type TriggerCRON interface {

	// Extract in charge of the "E" in the ETL process: it Extracts the data from
	// the source.
	Extract(*Toolkit) (*Event, error)
}

/*
Schedule represents a schedule at which a source's trigger should run.
*/
type Schedule struct {

	// Interval represents an interval or a CRON string at which a trigger shall be
	// triggered.
	Interval string `json:"interval"`
}
