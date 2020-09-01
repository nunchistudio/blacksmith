package service

import (
	"context"
	"crypto/tls"
	"net/http"
)

/*
Options is the options a user can pass to configure the gateway or the scheduler.
*/
type Options struct {

	// Context is a free key-value dictionary that will be passed to the underlying
	// service.
	Context context.Context `json:"-"`

	// Address is the HTTP address the server is listening to.
	Address string `json:"address"`

	// TLS is the TLS settings used to run the TLS server.
	TLS *tls.Config `json:"-"`

	// CertFile is the relative path to the certificate file for the TLS server.
	CertFile string `json:"-"`

	// KeyFile is the relative path to the key file for the TLS server.
	KeyFile string `json:"-"`

	// Middleware is the HTTP middleware chain that will be applied to the HTTP server.
	Middleware func(http.Handler) http.Handler `json:"-"`

	// Attach allows you to attach an external HTTP handler to the server. It is
	// useful for adding HTTP routes with custom routing and business logic.
	//
	// If a handler is attached, all routes within this handler will be prefixed with
	// "/api".
	Attach http.Handler `json:"-"`
}
