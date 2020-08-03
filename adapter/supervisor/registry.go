package supervisor

/*
Registry holds information about a service registry / mesh.
*/
type Registry struct {

	// Services is a list of the services registered in the supervisor. Only the
	// services registered via Blacksmith are listed.
	Services []*Service `json:"services"`
}

/*
Service is a service registered in the service registry.
*/
type Service struct {

	// ID is the unique identifier of the service.
	ID string `json:"id"`

	// Name is the name of the service.
	//
	// Example: "blacksmith-gateway"
	Name string `json:"name"`

	// Name is the version of the service being run at the moment.
	//
	// Example: "0.10.2"
	Version string `json:"version"`

	// Address is the address of the service.
	Address string `json:"address"`

	// Tags is a slice of tags related to the service.
	// They can be different from the node.
	Tags []string `json:"tags"`

	// Meta is a collection of meta-data related to the service.
	// They can be different from the node.
	Meta map[string]string `json:"meta"`

	// Nodes is a slice of nodes / instances registered for the service.
	Nodes []*Node `json:"nodes"`
}

/*
Node holds information about a node in a distributed environment the service running
is about to join.
*/
type Node struct {

	// ID is the unique identifier of the node to join.
	ID string `json:"id"`

	// Name is the name of the node to join.
	Name string `json:"name"`

	// Address is the address of the node to join.
	Address string `json:"address"`

	// Tags is a slice of tags related to the node.
	// They can be different from the services.
	Tags []string `json:"tags"`

	// Meta is a collection of meta-data related to the node.
	// They can be different from the services.
	Meta map[string]string `json:"meta"`
}

/*
Lock holds information about a lock-key in the distributed system. This allows to
acquire and release access to resources.
*/
type Lock struct {

	// Key is the id / name / key to acquire and release.
	Key string `json:"key"`
}
