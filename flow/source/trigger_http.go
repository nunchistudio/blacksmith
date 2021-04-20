package source

import (
	"net/http"
)

/*
ModeHTTP is used to indicate the event is triggered from a HTTP request.
*/
var ModeHTTP = "http"

/*
TriggerHTTP is the interface used for triggers using a HTTP route.

A new HTTP trigger can be generated using the Blacksmith CLI:

  $ blacksmith generate trigger --name <name> --mode http [--path <path>] [--migrations]
*/
type TriggerHTTP interface {

	// Extract in charge of the "E" in the ETL process: it Extracts the data from
	// the source.
	Extract(*Toolkit, *http.Request) (*Event, error)
}

/*
Route contains the details about a HTTP route used by the gateway.
*/
type Route struct {

	// Methods is a list of HTTP methods allowed for the route and the given path.
	//
	// Example: []string{"POST"}
	Methods []string `json:"methods"`

	// Path is the HTTP path of the route.
	//
	// Example: "/webhooks/crm/user"
	Path string `json:"path"`

	// ShowMeta is used to display (or not) the metadata in the HTTP response,
	// such as the event's context and jobs details.
	ShowMeta bool `json:"show_meta"`

	// ShowData is used to display (or not) the data in the HTTP response. It should
	// be disabled if any sensitive data can be returned, such as private tokens.
	ShowData bool `json:"show_data"`
}
