package source

/*
InterfaceSource is the string representation for the source interface.
*/
var InterfaceSource = "source"

/*
Source is the interface used to load events from cloud services, databases, or
any kind of application able to send data or push notifications.
*/
type Source interface {

	// String returns the string representation of the source.
	//
	// Example: "stripe"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// Triggers returns a list of triggers the source is able to take care of.
	// Their respective Extract function will automatically be triggered by the
	// gateway given their Mode.
	Triggers() map[string]Trigger
}
