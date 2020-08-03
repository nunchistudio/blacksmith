package blacksmith

import (
	"sync"

	"github.com/nunchistudio/blacksmith/adapter/destination"
	"github.com/nunchistudio/blacksmith/adapter/gateway"
	"github.com/nunchistudio/blacksmith/adapter/pubsub"
	"github.com/nunchistudio/blacksmith/adapter/scheduler"
	"github.com/nunchistudio/blacksmith/adapter/source"
	"github.com/nunchistudio/blacksmith/adapter/store"
	"github.com/nunchistudio/blacksmith/adapter/supervisor"
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

	// Supervisor is the supervisor adapter used for the entire application.
	// See the adapter/supervisor package for more details.
	Supervisor supervisor.Supervisor

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
*/
func (p *Pipeline) Start() error {
	var err error

	// Create channels to keep track of errors and stops.
	failed := make(chan error)

	// Create a waiting group that will wait for both the gateway and scheduler to
	// gracefully shutdown.
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Start the gateway in its own goroutine. If an error occured, write in the
	// error channel.
	go func() {
		err = p.StartGateway()
		if err != nil {
			failed <- err
		}

		wg.Done()
	}()

	// Start the scheduler in its own goroutine. If an error occured, write in the
	// error channel.
	go func() {
		err = p.StartScheduler()
		if err != nil {
			failed <- err
		}

		wg.Done()
	}()

	// Keep the servers running as long as there is no error or interruption.
	go func() {
		for {
			select {
			case <-failed:
				err = <-failed
			}
		}
	}()

	// Wait for the server to shutdown properly and return an error if any occured.
	wg.Wait()
	return err
}

/*
StartGateway starts the Blacksmith gateway.
*/
func (p *Pipeline) StartGateway() error {

	// Create the TLS certificate details.
	cert := &gateway.WithTLS{
		CertFile: p.Gateway.Options().CertFile,
		KeyFile:  p.Gateway.Options().KeyFile,
	}

	// Create the appropriate watcher toolkit for the gateway.
	gtk := &gateway.Toolkit{
		Logger:       p.Logger,
		Sources:      p.Sources,
		Destinations: p.Destinations,
		Store:        p.Store,
		PubSub:       p.PubSub,
		Supervisor:   p.Supervisor,
	}

	// Start the gateway server.
	p.Logger.Debugf("%s/%s: Server starting up...", gateway.InterfaceGateway, p.Gateway.String())
	err := p.Gateway.ListenAndServe(gtk, cert)
	if err != nil {
		return err
	}

	return nil
}

/*
StartScheduler starts the Blacksmith scheduler.
*/
func (p *Pipeline) StartScheduler() error {

	// Create the TLS certificate details.
	cert := &scheduler.WithTLS{
		CertFile: p.Scheduler.Options().CertFile,
		KeyFile:  p.Scheduler.Options().KeyFile,
	}

	// Create the appropriate watcher toolkit for the scheduler.
	stk := &scheduler.Toolkit{
		Logger:       p.Logger,
		Sources:      p.Sources,
		Destinations: p.Destinations,
		Store:        p.Store,
		PubSub:       p.PubSub,
		Supervisor:   p.Supervisor,
	}

	// Start the scheduler server.
	p.Logger.Debugf("%s/%s: Server starting up...", scheduler.InterfaceScheduler, p.Scheduler.String())
	err := p.Scheduler.ListenAndServe(stk, cert)
	if err != nil {
		return err
	}

	return nil
}
