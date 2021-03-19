package supervisor

import (
	"github.com/sirupsen/logrus"
)

/*
Toolkit contains a suite of utilities and data to help the adapter successfully
run the supervisor functions.
*/
type Toolkit struct {

	// Logger gives access to the logrus Logger passed in options when creating the
	// Blacksmith application.
	Logger *logrus.Logger

	// Service holds details about the service accessing the semaphore. It shall be
	// used by the adapter to allow (or not) the lock and unlock of a key. A running
	// service can not lock or unlock resources already used by an other service.
	Service *Service
}
