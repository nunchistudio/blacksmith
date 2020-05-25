package destination

import (
	"github.com/sirupsen/logrus"

	"github.com/nunchistudio/blacksmith/adapter/store"
)

/*
Toolkit contains a suite of utilities and data to help the user successfully run
the event functions.
*/
type Toolkit struct {

	// Logger gives access to the logrus Logger passed in options when creating the
	// Blacksmith application.
	Logger *logrus.Logger

	// Event is the event retrieved from the store, including all the elements
	// from the original event received by the gateway or the parent job.
	Event *store.Event

	// Job is the job being run at the moment. It is outside the Event key because
	// an event can have several jobs to execute.
	Job *store.Job
}
