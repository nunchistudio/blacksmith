package source

import (
	"net/http"
)

/*
ModeHTTP is used to indicate the event is trigeered from a HTTP request.
*/
var ModeHTTP = "http"

/*
TriggerHTTP is the interface used for triggers using a HTTP route.
*/
type TriggerHTTP interface {

	// Marshal in charge of the "E" in the ETL process: it Extracts the data from
	// the source.
	Marshal(*Toolkit, *http.Request) (*Payload, error)
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
}
