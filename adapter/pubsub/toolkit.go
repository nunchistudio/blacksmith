package pubsub

import (
	"context"

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

	// Context is the context originally passed when creating the pubsub.
	Context context.Context
}
