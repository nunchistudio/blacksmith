package destination

/*
WithHooks can be implemented by destinations and actions to add custom logic when
the scheduler service is starting and shutting down, or before and after running
migrations.
*/
type WithHooks interface {

	// Init lets you initialize a destination or an action. It is useful when
	// initialization is necessary, such as opening a connection pool with the
	// destination.
	//
	// Init is called when starting the scheduler service or before running migrations.
	// If an error is returned, the running process will try to gracefully shutdown.
	//
	// The Init function of a destination will always be executed before the ones
	// of its actions. Therefore, the Init function of an action will always be
	// executed after the one of its destination.
	//
	// Note: EventID and JobID in Toolkit will always be empty.
	Init(*Toolkit) error

	// Shutdown lets you gracefully shutdown a destination or an action. It is
	// useful when shutting down is necessary, such as closing a connection pool
	// with the destination.
	//
	// Shutdown is called when shutting down the scheduler service or after running
	// migrations. If an error is returned, it will only be logged and the running
	// process will continue its shutdown.
	//
	// The Shutdown function of a destination will always be executed after the ones
	// of its actions. Therefore, the Shutdown function of an action will always be
	// executed before the one of its destination.
	//
	// Note: EventID and JobID in Toolkit will always be empty.
	Shutdown(*Toolkit) error
}
