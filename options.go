package blacksmith

import (
	"github.com/nunchistudio/blacksmith/adapter/pubsub"
	"github.com/nunchistudio/blacksmith/adapter/store"
	"github.com/nunchistudio/blacksmith/adapter/supervisor"
	"github.com/nunchistudio/blacksmith/adapter/wanderer"
	"github.com/nunchistudio/blacksmith/flow/destination"
	"github.com/nunchistudio/blacksmith/flow/source"
	"github.com/nunchistudio/blacksmith/service"

	"github.com/sirupsen/logrus"
)

/*
Options is the options a user can pass to create a new application.
*/
type Options struct {

	// Logger allows you to use a logrus Logger across all Blacksmith adapters and
	// the application built on top of it.
	Logger *logrus.Logger `json:"-"`

	// Supervisor is the options passed to use the supervisor adapter.
	// The supervisor is optional.
	Supervisor *supervisor.Options `json:"supervisor"`

	// Wanderer is the options passed to use the wanderer adapter.
	// The wanderer is optional.
	Wanderer *wanderer.Options `json:"wanderer"`

	// Store is the options passed to use the store adapter.
	Store *store.Options `json:"store"`

	// PubSub is the options passed to use the pubsub adapter.
	// The pusub is optional.
	PubSub *pubsub.Options `json:"pubsub"`

	// Gateway is the options passed to use the gateway service.
	Gateway *service.Options `json:"gateway"`

	// Scheduler is the options passed to use the scheduler service.
	Scheduler *service.Options `json:"scheduler"`

	// Sources is a slice of options passed to create sources.
	Sources []*source.Options `json:"-"`

	// Destinations is a slice of options passed to create destinations.
	Destinations []*destination.Options `json:"-"`

	// License holds the license details for using Blacksmith Enterprise Edition.
	// This is not necessary for using the Standard Edition.
	License *License `json:"-"`
}

/*
License holds the license details for using Blacksmith Enterprise Edition.
*/
type License struct {

	// Key is the subscription license key. This information can also be set using
	// the environment variable `BLACKSMITH_LICENSE_KEY`
	Key string `json:"-"`

	// Token is the subscription license token. This information can also be set
	// using the environment variable `BLACKSMITH_LICENSE_TOKEN`.
	Token string `json:"-"`
}
