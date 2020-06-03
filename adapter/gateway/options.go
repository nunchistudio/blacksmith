package gateway

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/internal/adapter"
)

/*
AvailableAdapters is a list of available gateway adapters.
*/
var AvailableAdapters = map[string]bool{
	"standard": true,
}

/*
Defaults are the defaults options set for the gateway. When not set, these values
will automatically be applied.
*/
var Defaults = &Options{
	From:    "standard",
	Address: ":8080",
}

/*
Options is the options a user can pass to create a new gateway.
*/
type Options struct {

	// From can be used to download, install, and use an existing adapter. This way
	// the user does not need to develop a custom gateway adapter.
	From string

	// Load can be used to load and use a custom gateway adapter developed in-house.
	Load Gateway

	// Context is a free key-value dictionary that will be passed to the underlying
	// adapter.
	Context context.Context

	// Address is the HTTP address the gateway server is listening to.
	//
	// Defaults to ":8080".
	Address string

	// TLS is the TLS settings used to run the TLS server.
	TLS *tls.Config

	// CertFile is the relative path to the certificate file for the TLS server.
	CertFile string

	// KeyFile is the relative path to the key file for the TLS server.
	KeyFile string

	// Middleware is the HTTP middleware chain that will be applied to the HTTP server
	// of the gateway.
	Middleware func(http.Handler) http.Handler

	// Attach allows you to attach an external HTTP handler to the Blacksmith gateway.
	// It is useful for adding HTTP routes with custom routing and business logic.
	//
	// If a handler is attached, all routes within this handler will be prefixed with
	// a prefix chosen by the gateway adapter.
	Attach http.Handler
}

/*
ValidateAndLoad validates the gateway's options and returns a valid gateway
interface.
*/
func (opts *Options) ValidateAndLoad() (Gateway, error) {
	var g Gateway
	var err error

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "gateway: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Set default options needed.
	if opts == nil {
		opts = Defaults
	}

	// Use the custom adapter if the user passed one. Otherwise, make sure the
	// gateway adapter is a valid one and load it from the Go plugin.
	if opts.Load != nil {
		g = opts.Load
	} else {
		if opts.From == "" {
			opts.From = Defaults.From
		}

		if _, exists := AvailableAdapters[opts.From]; !exists {
			fail.Validations = append(fail.Validations, errors.Validation{
				Message: "Adapter not supported",
				Path:    []string{"Options", "Gateway", "From"},
			})

			return nil, fail
		}

		g, err = opts.loadPlugin()
		if err != nil {
			return nil, err
		}
	}

	// If the user didn't put an address, use the default one.
	if opts.Address == "" {
		opts.Address = Defaults.Address
	}

	// Validate the gateway adapter.
	err = validateGateway(g)
	if err != nil {
		return nil, err
	}

	// We are now sure to be able to use the adapter.
	return g, nil
}

/*
loadPlugin loads a Go plugin using the adapter ID from the gateway options.
It returns the gateway interface loaded from the Go plugin.
*/
func (opts *Options) loadPlugin() (Gateway, error) {

	// Load the Go plugin's symbol from the helper.
	symbol, err := adapter.LoadPlugin(InterfaceGateway, opts.From)
	if err != nil {
		return nil, err
	}

	// Convert the symbol to the desired type.
	converted := symbol.(func(*Options) (Gateway, error))

	// Load the Go plugin's gateway adapter.
	h, err := converted(opts)
	if err != nil {
		return nil, err
	}

	// Finally, return the gateway adapter.
	return h, nil
}
