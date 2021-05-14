package supervisor

/*
InterfaceSupervisor is the string representation for the supervisor interface.
*/
var InterfaceSupervisor = "supervisor"

/*
Supervisor is the interface used to properly run Blacksmith applications in
distributed environments. This allows strong data consistency and better
infrastructure reliability.
*/
type Supervisor interface {

	// String returns the string representation of the adapter.
	//
	// Example: "consul"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// Init lets you initialize the Supervisor. It is useful to create a session
	// across nodes and register a service instance in the service registry if
	// applicable.
	Init(*Toolkit) error

	// Shutdown lets you gracefully shutdown a service in the Supervisor. It is
	// useful to destroy a session and deregister a service instance from the
	// service registry if applicable.
	Shutdown(*Toolkit) error

	// Lock allows to acquire a key in the semaphore. It returns true if the key
	// is successfully acquired.
	Lock(*Toolkit, *Lock) (bool, error)

	// Unlock allows to release a key from the semaphore. It returns true if the
	// key is successfully released.
	Unlock(*Toolkit, *Lock) (bool, error)

	// Status returns the semaphore status for a given key.
	Status(*Toolkit, *Lock) (*Semaphore, error)
}

/*
Service is a service registered in the service registry.
*/
type Service struct {

	// ID is the unique identifier of the service.
	//
	// Example: "1UYc8EebLqCAFMOSkbYZdJwNLAJ"
	ID string `json:"id"`

	// Name is the name of the service.
	//
	// Example: "blacksmith-gateway"
	Name string `json:"name"`

	// Version is the Blacksmith version being run at the moment by the service.
	//
	// Example: "0.17.0"
	Version string `json:"version"`

	// Address is the address of the service.
	//
	// Example: ":9090"
	Address string `json:"address"`

	// Tags is a slice of tags related to the service.
	Tags []string `json:"tags"`

	// Meta is a collection of meta-data related to the service.
	Meta map[string]string `json:"meta"`
}
