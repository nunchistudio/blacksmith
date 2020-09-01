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

	// Run returns a slice of actions to run, grouped by their destination name.
	Run(*Toolkit) destination.Actions
}
