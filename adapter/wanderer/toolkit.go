package wanderer

import (
	"github.com/sirupsen/logrus"
)

/*
Toolkit contains a suite of utilities and data to help the adapter successfully
run the wanderer functions.
*/
type Toolkit struct {

	// Logger gives access to the logrus Logger passed in options when creating the
	// Blacksmith application.
	Logger *logrus.Logger

	// WD is the rooted path name corresponding to the current directory. It can be
	// used to read a migration file in a directory.
	WD string
}
