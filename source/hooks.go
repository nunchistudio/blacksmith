package source

/*
WithHooks can be implemented by sources and triggers to add custom logic when the
gateway service is starting and shutting down, or before and after running
migrations.
*/
type WithHooks interface {

	// Init lets you initialize a source or a trigger. It is useful when initialization
	// is necessary, such as opening a connection pool with the source.
	//
	// Init is called when starting the gateway service or before running migrations.
	// If an error is returned, the running process will try to gracefully shutdown.
	//
	// The Init function of a source will always be executed before the ones of its
	// triggers. Therefore, the Init function of a trigger will always be executed
	// after the one of its source.
	//
	// Note: EventID in Toolkit will always be empty.
	Init(*Toolkit) error

	// Shutdown lets you gracefully shutdown a source or a trigger. It is useful
	// when shutting down is necessary, such as closing a connection pool with the
	// source.
	//
	// Shutdown is called when shutting down the gateway service or after running
	// migrations. If an error is returned, it will only be logged and the running
	// process will continue its shutdown.
	//
	// The Shutdown function of a source will always be executed after the ones of
	// its triggers. Therefore, the Shutdown function of a trigger will always be
	// executed before the one of its source.
	//
	// Note: EventID in Toolkit will always be empty.
	Shutdown(*Toolkit) error
}
