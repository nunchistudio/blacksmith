package flow

import (
	"github.com/nunchistudio/blacksmith/flow/destination"
)

/*
Flow is a middleman allowing triggers to run actions.
*/
type Flow interface {

	// Options returns the options originally passed to the Options struct.
	Options() *Options

	// Transform returns a slice of actions to run, grouped by their destination
	// name. It is in charge of the "T" in the ETL process: it is used to Transform
	// data from triggers to actions.
	Transform(*Toolkit) destination.Actions
}
