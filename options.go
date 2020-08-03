package blacksmith

import (
	"os"
	"sync"

	"github.com/nunchistudio/blacksmith/adapter/destination"
	"github.com/nunchistudio/blacksmith/adapter/gateway"
	"github.com/nunchistudio/blacksmith/adapter/pubsub"
	"github.com/nunchistudio/blacksmith/adapter/scheduler"
	"github.com/nunchistudio/blacksmith/adapter/source"
	"github.com/nunchistudio/blacksmith/adapter/store"
	"github.com/nunchistudio/blacksmith/adapter/supervisor"
	"github.com/nunchistudio/blacksmith/adapter/wanderer"
	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/helper/logger"

	"github.com/sirupsen/logrus"
)

/*
Options is the options a user can pass to create a new application.
*/
type Options struct {

	// Logger allows you to use a logrus Logger across all Blacksmith adapters and
	// the application built on top of it.
	Logger *logrus.Logger

	// Supervisor is the options passed to create a new supervisor adapter.
	// The supervisor is optional.
	Supervisor *supervisor.Options

	// Wanderer is the options passed to create a new wanderer adapter.
	// The wanderer is optional.
	Wanderer *wanderer.Options

	// Store is the options passed to create a new store adapter.
	Store *store.Options

	// PubSub is the options passed to create a new pubsub adapter.
	// The pusub is optional.
	PubSub *pubsub.Options

	// Gateway is the options passed to create a new gateway adapter.
	Gateway *gateway.Options

	// Scheduler is the options passed to create a new scheduler adapter.
	Scheduler *scheduler.Options

	// Sources is a slice of options passed to create source adapters.
	Sources []*source.Options

	// Destinations is a slice of options passed to create destination adapters.
	Destinations []*destination.Options
}

/*
ValidateAndLoad validates the application's options and returns a valid Blacksmith
application.
*/
func (opts *Options) ValidateAndLoad() (*Pipeline, error) {
	fail := &errors.Error{
		Message:     "options: Failed to validate and load",
		Validations: []errors.Validation{},
	}

	// Find the current working directory.
	wd, err := os.Getwd()
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return nil, fail
	}

	// Use the given logger. If none where passed use the default one.
	var log *logrus.Logger = opts.Logger
	if opts.Logger == nil {
		log = logger.New()
	}

	// Initialize a new data pipeline.
	p := &Pipeline{
		mutex:        &sync.Mutex{},
		WD:           wd,
		Logger:       log,
		Sources:      map[string]source.Source{},
		Destinations: map[string]destination.Destination{},
	}

	// Make sure lock and unlock the mutex.
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Validate and load the supervisor. Since the supervisor is optional, only
	// validate it if provided.
	if opts.Supervisor != nil {
		p.Logger.Debug("supervisor: Validating and loading adapter...")
		p.Supervisor, err = opts.Supervisor.ValidateAndLoad()
		if p.Supervisor == nil || err != nil {
			p.Logger.Error(err)
			return nil, err
		}
	}

	// Validate and load the wanderer. Since the wanderer is optional, only validate
	// it if provided.
	if opts.Wanderer != nil {
		p.Logger.Debug("wanderer: Validating and loading adapter...")
		p.Wanderer, err = opts.Wanderer.ValidateAndLoad()
		if p.Wanderer == nil || err != nil {
			p.Logger.Error(err)
			return nil, err
		}
	}

	// Validate and load the store.
	p.Logger.Debug("store: Validating and loading adapter...")
	p.Store, err = opts.Store.ValidateAndLoad()
	if p.Store == nil || err != nil {
		p.Logger.Error(err)
		return nil, err
	}

	// Validate and load the pusub. Since the pusub is optional, only validate
	// it if provided.
	if opts.PubSub != nil {
		p.Logger.Debug("pubsub: Validating and loading adapter...")
		p.PubSub, err = opts.PubSub.ValidateAndLoad()
		if err != nil {
			p.Logger.Error(err)
			return nil, err
		}
	}

	// Validate and load the sources.
	if opts.Sources != nil {
		for _, options := range opts.Sources {
			p.Logger.Debug("source: Validating and loading adapter...")
			s, err := options.ValidateAndLoad()
			if s == nil || err != nil {
				p.Logger.Error(err)
				return nil, err
			}

			p.Sources[s.String()] = s
			p.Logger.Debugf("source/%s: Validated and added", s.String())
		}
	}

	// Validate and load the destinations.
	if opts.Destinations != nil {
		for _, options := range opts.Destinations {
			p.Logger.Debug("destination: Validating and loading adapter...")
			d, err := options.ValidateAndLoad()
			if d == nil || err != nil {
				p.Logger.Error(err)
				return nil, err
			}

			p.Destinations[d.String()] = d
			p.Logger.Debugf("destination/%s: Validated and added", d.String())
		}
	}

	// Validate and load the gateway.
	p.Logger.Debug("gateway: Validating and loading adapter...")
	p.Gateway, err = opts.Gateway.ValidateAndLoad()
	if p.Gateway == nil || err != nil {
		p.Logger.Error(err)
		return nil, err
	}

	// Validate and load the scheduler.
	p.Logger.Debug("scheduler: Validating and loading adapter...")
	p.Scheduler, err = opts.Scheduler.ValidateAndLoad()
	if p.Scheduler == nil || err != nil {
		p.Logger.Error(err)
		return nil, err
	}

	// Return the pipeline.
	return p, nil
}
