package blacksmith

import (
	"sync"

	"github.com/nunchistudio/blacksmith/adapter/destination"
	"github.com/nunchistudio/blacksmith/adapter/gateway"
	"github.com/nunchistudio/blacksmith/adapter/pubsub"
	"github.com/nunchistudio/blacksmith/adapter/scheduler"
	"github.com/nunchistudio/blacksmith/adapter/source"
	"github.com/nunchistudio/blacksmith/adapter/store"
	"github.com/nunchistudio/blacksmith/adapter/wanderer"

	"github.com/sirupsen/logrus"
)

/*
Pipeline centrally stores every details and adapters needed to run a Blacksmith
application.
*/
type Pipeline struct {

	// mutex is the local mutex.
	mutex *sync.Mutex

	// WD is the working directory.
	WD string

	// Logger is the logrus Logger for the entire application.
	Logger *logrus.Logger

	// Wanderer is the wanderer adapter used for the entire application.
	// See the adapter/wanderer package for more details.
	Wanderer wanderer.Wanderer

	// Store is the store adapter used for the entire application.
	// See the adapter/store package for more details.
	Store store.Store

	// PubSub is the pubsub adapter used for the entire application.
	// See the adapter/pubsub package for more details.
	PubSub pubsub.PubSub

	// Gateway is the gateway adapter used for the entire application.
	// See the adapter/gateway package for more details.
	Gateway gateway.Gateway

	// Scheduler is the scheduler adapter used for the entire application.
	// See the adapter/scheduler package for more details.
	Scheduler scheduler.Scheduler

	// Sources is a collection of sources used in the application.
	// See the adapter/source package for more details.
	Sources map[string]source.Source

	// Destinations is a collection of destinations used in the application.
	// See the adapter/destination package for more details.
	Destinations map[string]destination.Destination
}

/*
New validates options and load a new data pipeline if no error occurred during the
validation.
*/
func New(opts *Options) (*Pipeline, error) {
	return opts.ValidateAndLoad()
}

/*
Start starts both the Blacksmith gateway and scheduler. It is useful for working
in non-production environments or in application with low traffic. Otherwise,
it is recommended to start the gateway and scheduler in their own process.

TODO: Improve the behavior of this function when the gateway or scheduler exits.
Currently, the standard adapters can throw an "error" when shutting their HTTP
server. This should not happen.
*/
func (p *Pipeline) Start() error {

	// Create channels to keep track of errors and stops.
	failed := make(chan error)
	stopped := make(chan bool)

	// Start the gateway in its own goroutine. If an error occured, write in the
	// error channel.
	go func() {
		err := p.StartGateway()
		if err != nil {
			failed <- err
		}

		stopped <- true
	}()

	// Start the scheduler in its own goroutine. If an error occured, write in the
	// error channel.
	go func() {
		err := p.StartScheduler()
		if err != nil {
			failed <- err
		}

		stopped <- true
	}()

	// Keep the servers running as long as there is no error or interruption.
	for {
		select {
		case err := <-failed:
			return err
		case <-stopped:
			return nil
		}
	}
}

/*
StartGateway starts the Blacksmith gateway. This also make sure other adapters are
properly initialized and shutdown when necessary.
*/
func (p *Pipeline) StartGateway() error {
	p.Logger.Debugf("gateway/%s: Starting HTTP server...", p.Gateway.String())

	// Create the appropriate pubsub toolkit.
	pstk := &pubsub.Toolkit{
		Logger: p.Logger,
	}

	// If the pubsub exists, add the context to the toolkit.
	if p.PubSub != nil {
		pstk.Context = p.PubSub.Options().Context
	}

	// Initialize the pubsub Publisher if needed.
	if p.PubSub != nil {
		if p.PubSub.Publisher() != nil {
			p.Logger.Debugf("pubsub/%s: Initializing publisher...", p.PubSub.String())
			err := p.PubSub.Publisher().Init(pstk)
			if err != nil {
				return err
			}
		}
	}

	// Create the appropriate watcher toolkit for the gateway.
	ltk := &gateway.Toolkit{
		Logger:       p.Logger,
		Sources:      p.Sources,
		Destinations: p.Destinations,
		Store:        p.Store,
		PubSub:       p.PubSub,
	}

	// Create the TLS certificate details.
	cert := &gateway.WithTLS{
		CertFile: p.Gateway.Options().CertFile,
		KeyFile:  p.Gateway.Options().KeyFile,
	}

	// Start the gateway server.
	p.Logger.Debugf("gateway/%s: HTTP server starting up...", p.Gateway.String())
	if err := p.Gateway.ListenAndServe(ltk, cert); err != nil {
		return err
	}

	// Make sure to shutdown the pubsub Publisher properly.
	if p.PubSub != nil {
		if p.PubSub.Publisher() != nil {
			p.Logger.Debugf("pubsub/%s: Shutting down publisher...", p.PubSub.String())
			p.PubSub.Publisher().Shutdown(pstk)
		}
	}

	p.Logger.Infof("gateway/%s: HTTP server shut down", p.Gateway.String())
	return nil
}

/*
StartScheduler starts the Blacksmith scheduler. This also make sure other adapters
are properly initialized and shutdown when necessary.
*/
func (p *Pipeline) StartScheduler() error {
	p.Logger.Debugf("scheduler/%s: Starting HTTP server...", p.Scheduler.String())

	// Create the appropriate pubsub toolkit.
	pstk := &pubsub.Toolkit{
		Logger: p.Logger,
	}

	// If the pubsub exists, add the context to the toolkit.
	if p.PubSub != nil {
		pstk.Context = p.PubSub.Options().Context
	}

	// Initialize the pubsub Subscriber if needed.
	if p.PubSub != nil {
		if p.PubSub.Subscriber() != nil {
			p.Logger.Debugf("pubsub/%s: Initializing subscriber...", p.PubSub.String())
			err := p.PubSub.Subscriber().Init(pstk)
			if err != nil {
				return err
			}
		}
	}

	// Create the appropriate watcher toolkit for the scheduler.
	ltk := &scheduler.Toolkit{
		Logger:       p.Logger,
		Sources:      p.Sources,
		Destinations: p.Destinations,
		Store:        p.Store,
		PubSub:       p.PubSub,
	}

	// Create the TLS certificate details.
	cert := &scheduler.WithTLS{
		CertFile: p.Scheduler.Options().CertFile,
		KeyFile:  p.Scheduler.Options().KeyFile,
	}

	// Start the scheduler server.
	p.Logger.Debugf("scheduler/%s: HTTP server starting up...", p.Scheduler.String())
	if err := p.Scheduler.ListenAndServe(ltk, cert); err != nil {
		return err
	}

	// Make sure to shutdown the pubsub Subscriber properly.
	if p.PubSub != nil {
		if p.PubSub.Subscriber() != nil {
			p.Logger.Debugf("pubsub/%s: Shutting down subscriber...", p.PubSub.String())
			p.PubSub.Subscriber().Shutdown(pstk)
		}
	}

	p.Logger.Infof("scheduler/%s: HTTP server shut down", p.Scheduler.String())
	return nil
}
