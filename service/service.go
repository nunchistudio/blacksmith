package service

import (
	"net/http"
)

/*
Service is the interface used to create the gateway and scheduler services.
*/
type Service interface {

	// String returns the string representation of the service.
	//
	// Example: "enterprise"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// Handler returns a net/http Handler allowing the use of the service as a
	// standard HTTP handler in an external Go application.
	Handler(*Toolkit) (http.Handler, error)

	// ListenAndServe starts the HTTP server. This is the equivalent of the net/http
	// ListenAndServe.
	ListenAndServe(*Toolkit, *WithTLS) error

	// Shutdown gracefully shuts down the server without interrupting any active
	// connections such as CRON tasks. It is the equivalent of the net/http Shutdown
	// function.
	Shutdown(*Toolkit) error
}

/*
WithTLS allows you to attach TLS certificate files when creating the HTTP server.
*/
type WithTLS struct {

	// CertFile is the relative path to the certificate file.
	CertFile string `json:"-"`

	// KeyFile is the relative path to the key file.
	KeyFile string `json:"-"`
}
