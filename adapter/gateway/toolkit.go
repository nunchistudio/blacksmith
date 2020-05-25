package gateway

import (
	"github.com/sirupsen/logrus"

	"github.com/nunchistudio/blacksmith/adapter/destination"
	"github.com/nunchistudio/blacksmith/adapter/pubsub"
	"github.com/nunchistudio/blacksmith/adapter/source"
	"github.com/nunchistudio/blacksmith/adapter/store"
)

/*
Toolkit contains a suite of utilities to help the user successfully run the
ListenAndServe function.
*/
type Toolkit struct {

	// Logger gives access to the logrus Logger passed in options when creating the
	// Blacksmith application.
	Logger *logrus.Logger

	// Sources is the collection of sources registered in the Blacksmith application.
	// It is up to the gateway to know which sources' events need to be watched.
	Sources map[string]source.Source

	// Destinations is the collection of destinations registered in the Blacksmith
	// application. It is up to the gateway to know which destinations' events need
	// to be resolved.
	Destinations map[string]destination.Destination

	// Store is the store adapter registered in the Blacksmith application. You can
	// use it to insert new transitions the datastore.
	Store store.Store

	// PubSub is the pubsub adapter registered in the Blacksmith application. You
	// can use it to publish events in realtime using the pubsub Publisher.
	PubSub pubsub.PubSub
}
