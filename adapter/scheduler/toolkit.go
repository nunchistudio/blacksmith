package scheduler

import (
	"github.com/nunchistudio/blacksmith/adapter/destination"
	"github.com/nunchistudio/blacksmith/adapter/pubsub"
	"github.com/nunchistudio/blacksmith/adapter/source"
	"github.com/nunchistudio/blacksmith/adapter/store"
	"github.com/nunchistudio/blacksmith/adapter/supervisor"

	"github.com/sirupsen/logrus"
)

/*
Toolkit contains a suite of utilities to help the adapter successfully run the
scheduler.
*/
type Toolkit struct {

	// Logger gives access to the logrus Logger passed in options when creating the
	// Blacksmith application.
	Logger *logrus.Logger

	// Sources is the collection of sources registered in the Blacksmith application.
	Sources map[string]source.Source

	// Destinations is the collection of destinations registered in the Blacksmith
	// application.
	Destinations map[string]destination.Destination

	// Store is the store adapter registered in the Blacksmith application.
	Store store.Store

	// PubSub is the pubsub adapter registered in the Blacksmith application.
	PubSub pubsub.PubSub

	// Supervisor is the supervisor adapter registered in the Blacksmith application.
	Supervisor supervisor.Supervisor
}
