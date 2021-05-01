package source

import (
	"time"

	"github.com/nunchistudio/blacksmith/flow"
	"github.com/nunchistudio/blacksmith/flow/destination"
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
	// source's trigger. The gateway will trigger the Extract function based on
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
	// - When set to ModeCDC, no additional options are required.
	// - When set to ModeSubscription, the UsingSubscription options is used as the
	//   trigger. It uses the Pub / Sub adapter configured for the application.
	Mode string `json:"mode"`

	// UsingHTTP defines the HTTP route the event will react to.
	UsingHTTP *Route `json:"http,omitempty"`

	// UsingCRON represents a schedule at which an event should run. When returning
	// nil, the source's schedule is applied.
	UsingCRON *Schedule `json:"cron,omitempty"`

	// UsingSubscription defines the Pub / Sub subscription to use.
	UsingSubscription *Subscription `json:"subscription,omitempty"`
}

/*
Event represents the fields a trigger shall fill to create an event. It will be
used across the application to match the fields between sources and destinations.
*/
type Event struct {

	// Version is the version number of the source used by the event's when triggered.
	//
	// Examples: "v1.0", "2020-10-01"
	Version string `json:"version,omitempty"`

	// Context is a dictionary of information that provides useful context about an
	// event. The context should be used inside every events for consistency.
	//
	// It must be a valid JSON since it will be used using encoding/json Marshal and
	// Unmarshal functions.
	Context []byte `json:"context"`

	// Data is the byte representation of the data sent by the event.
	//
	// It must be a valid JSON since it will be used using encoding/json Marshal and
	// Unmarshal functions.
	Data []byte `json:"data"`

	// Actions allows to create jobs directly from the event when the trigger is
	// executed. Where Flows is used to create jobs from a common data schema, calling
	// Actions directly not from a flow can often be simpler.
	Actions destination.Actions `json:"-"`

	// Flows defines the flows of actions to run when this trigger is executed.
	// See package flow for more details.
	Flows []flow.Flow `json:"-"`

	// SubEvents is a collection of events that need to be created and associated
	// to the event being processed. This is useful for triggers accepting a batch
	// of events in a single request.
	SubEvents []*SubEvent `json:"sub_events"`

	// SentAt allows you to keep track of the timestamp when the event was originally
	// sent.
	SentAt *time.Time `json:"sent_at,omitempty"`
}

/*
SubEvent represents an event created by and associated to a parent event. A SubEvent
can be returned by a Trigger when extracting data, allowing to process a batch of
events in a single request.
*/
type SubEvent struct {

	// Trigger is the trigger's name of the sub-event created.
	Trigger string `json:"trigger"`

	// Context is a dictionary of information that provides useful context about a
	// sub-event.
	//
	// It must be a valid JSON since it will be used using encoding/json Marshal and
	// Unmarshal functions.
	Context []byte `json:"context"`

	// Data is the byte representation of the data sent by the sub-event.
	//
	// It must be a valid JSON since it will be used using encoding/json Marshal and
	// Unmarshal functions.
	Data []byte `json:"data"`

	// Actions allows to create jobs directly from the event when the trigger is
	// executed. Where Flows is used to create jobs from a common data schema, calling
	// Actions directly not from a flow can often be simpler.
	//
	// The jobs created by the actions are related to the sub-event, not the parent
	// event.
	Actions destination.Actions `json:"-"`

	// Flows defines the flows of actions to run when this trigger is executed.
	// See package flow for more details.
	//
	// The jobs created by the actions returned are related to the sub-event, not
	// the parent event.
	Flows []flow.Flow `json:"-"`
}
