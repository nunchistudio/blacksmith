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
}
