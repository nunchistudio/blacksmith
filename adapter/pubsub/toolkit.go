package pubsub

import (
	"github.com/sirupsen/logrus"
)

/*
Toolkit gives you access to a set of usefull tools when dealing with Publisher and
Subscriber.
*/
type Toolkit struct {

	// Logger gives access to the logrus Logger passed in options when creating the
	// Blacksmith application.
	Logger *logrus.Logger
}
