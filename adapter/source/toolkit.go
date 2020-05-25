package source

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

/*
Toolkit contains a suite of utilities and data to help the user successfully run
the Marshal function.
*/
type Toolkit struct {

	// Logger gives access to the logrus Logger passed in options when creating the
	// Blacksmith application.
	Logger *logrus.Logger

	// Request is the original net/http Request. It shall not be modified by any
	// means before you access it.
	//
	// Note: Since this is the HTTP request, this will always be nil on events using
	// the "schedule" mode.
	Request *http.Request

	// Payload allows you to write the payload and send it to the gateway using this
	// channel instead of returning it from the function. It is only used when dealing
	// with infinite loop inside "forever" events.
	//
	// Note: Payload is only used for events using the "forever" mode. Otherwise, this
	// is nil. This is the equivalent of returning the payload using the return keyword
	// on regular function.
	Payload chan<- *Payload

	// Error allows you to write an error (if any occurred) and send it to the gateway
	// using this channel instead of returning it from the function. It is only used
	// when dealing with infinite loop inside "forever" events.
	//
	// Note: Error is only used for events using the "forever" mode. Otherwise, this
	// is nil. This is the equivalent of returning an error using the return keyword
	// on regular function.
	Error chan<- error
}
