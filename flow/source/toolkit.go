package source

import (
	"github.com/nunchistudio/blacksmith/adapter/supervisor"

	"github.com/sirupsen/logrus"
)

/*
Toolkit contains a suite of utilities and data to help the user successfully run
the source functions.
*/
type Toolkit struct {

	// Logger gives access to the logrus Logger passed in options when creating the
	// Blacksmith application.
	Logger *logrus.Logger

	// Service represents the instance of the gateway service registered in the
	// supervisor and currently executing the event.
	//
	// Note: This is nil when there is no supervisor adapter configured.
	Service *supervisor.Service

	// EventID is the unique identifier of the event generated by the gateway and
	// that is being processed.
	//
	// Note: This is not applicable for triggers using the CDC mode.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	EventID string
}
