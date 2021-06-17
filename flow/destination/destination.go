package destination

/*
InterfaceDestination is the string representation for the destination interface.
*/
var InterfaceDestination = "destination"

/*
Destination is the interface used to load events to third-party services. Those
can be of any kind, such as APIs or databases.

A new destination can be generated using the Blacksmith CLI:

  $ blacksmith generate destination --name <name> [--path <path>] [--migrations]
*/
type Destination interface {

	// String returns the string representation of the destination.
	//
	// Example: "zendesk"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// Actions returns a list of actions the destination can handle. Destinations'
	// actions are run from sources' triggers and can also be triggered by other
	// destinations' actions. When a destination's action is called, it is
	// represented as a "job" in the platform.
	Actions() map[string]Action
}
