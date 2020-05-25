package wanderer

import (
	"github.com/nunchistudio/blacksmith/helper/mutex"

	"github.com/sirupsen/logrus"
)

/*
Toolkit contains a suite of utilities and data to help the user successfully run
the functions against the wanderer.
*/
type Toolkit struct {

	// Logger gives access to the logrus Logger passed in options when creating the
	// Blacksmith application.
	Logger *logrus.Logger

	// Lock let you use a remote access lock using the mutex package. It is not
	// mandatory but highly recommended to avoid data access conflicts.
	Lock *mutex.Lock

	// WD is the rooted path name corresponding to the current directory. It can be
	// used to read a migration file in a directory.
	WD string
}
