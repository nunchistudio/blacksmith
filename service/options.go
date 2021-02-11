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

	// Admin is the options used to setup the admin REST API and attach it to a
	// service.
	//
	// Reference: https://nunchi.studio/blacksmith/api/introduction/overview
	//
	// Note: Feature only available in Blacksmith Enterprise Edition.
	Admin *Admin `json:"admin"`
}

/*
Admin is the options used to attach the admin REST API and dashboard to a service.

Note: Feature only available in Blacksmith Enterprise Edition.
*/
type Admin struct {

	// Enabled allows to enable the REST API and therefore attach it to a service.
	//
	// API reference: https://nunchi.studio/blacksmith/api
	Enabled bool `json:"enabled"`

	// WithDashboard allows to attach the default Blacksmith dashboard to the
	// service. It can only be set if Enabled is also set to `true`.
	//
	// Note: The default dashboard is an open-source "template" repository. It can
	// be forked and can run on its own server if you do not want to attach it here.
	//
	// Repository: https://github.com/nunchistudio/blacksmith-dashboard
	WithDashboard bool `json:"dashboard"`

	// Middleware is the HTTP middleware chain that will be applied to the admin
	// API.
	Middleware func(http.Handler) http.Handler `json:"-"`
}
