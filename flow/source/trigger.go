package source

import (
	"time"

	"github.com/nunchistudio/blacksmith/flow"
)

/*
InterfaceTrigger is the string representation for the source's trigger interface.
*/
var InterfaceTrigger = "source/trigger"

/*
Trigger represents a trigger of a source adapter. A Trigger contains all the logic
to handle a specific event for a source.
*/
type Trigger interface {

	// String returns the string representation of the source's trigger.
	//
	// Example: "identify"
	String() string

	// Mode indicates the trigger mode to use along some options to execute the
	// source's trigger. The gateway will trigger the Marshal function based on
	// these options.
	Mode() *Mode
}

/*
Mode indicates how a source's trigger is triggered.
*/
type Mode struct {

	// Mode indicates the trigger mode to trigger the event.
	//
	// - When set to ModeHTTP, the UsingHTTP route is used as the trigger.
	// - When set to ModeCRON, the UsingCRON schedule is used as the trigger.
	// - When set to ModeCDC, no trigger is registered since it is an ongoing
	//   listener. It is up to the Marshal function to include the infinite loop
	//   and return the payload using the channel passed in params.
	Mode string `json:"mode"`

	// UsingHTTP defines the HTTP route the event will react to.
	UsingHTTP *Route `json:"http,omitempty"`

	// UsingCRON represents a schedule at which an event should run. When returning
	// nil, the source's schedule is applied.
	UsingCRON *Schedule `json:"cron,omitempty"`
}

/*
Payload represents the fields a trigger must fill. It will be used across the
application to match the fields between sources and destinations.
*/
type Payload struct {

	// Context is a dictionary of information that provides useful context about an
	// event. The context should be used inside every events for consistency.
	//
	// It must be a valid JSON since it will be used using json Marshal and Unmarshal
	// functions.
	Context []byte `json:"context"`

	// Data is the byte representation of the data sent by the event.
	//
	// It must be a valid JSON since it will be used using json Marshal and Unmarshal
	// functions.
	Data []byte `json:"data"`

	// Flows defines the flows of actions to run when this trigger is executed.
	// See package flow for more details.
	Flows []flow.Flow

	// SentAt allows you to keep track of the timestamp when the event was originally
	// sent.
	SentAt *time.Time `json:"sent_at,omitempty"`
}
