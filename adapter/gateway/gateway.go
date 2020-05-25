package gateway

import (
	"fmt"
	"net/http"

	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
InterfaceGateway is the string representation for the gateway interface.
*/
var InterfaceGateway = "gateway"

/*
Gateway is the interface used to receive and handle events from sources using a
HTTP server.
*/
type Gateway interface {

	// String returns the string representation of the adapter.
	//
	// Example: "standard"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// Handler returns a net/http Handler allowing the use of the gateway as a
	// standard HTTP handler in an external Go application.
	Handler(*Toolkit) (http.Handler, error)

	// ListenAndServe starts the HTTP server. This is the equivalent of the net/http
	// ListenAndServe and function except it is wrapped so you can watch for sources'
	// events and handle them across the different adapters such as the store and
	// destinations.
	ListenAndServe(*Toolkit, *WithTLS) error

	// Shutdown gracefully shuts down the server without interrupting any active
	// connections. It is the equivalent of the net/http Shutdown function.
	Shutdown(*Toolkit) error
}

/*
WithTLS allows you to attach TLS certificate files when creating the HTTP server.
*/
type WithTLS struct {

	// CertFile is the relative path to the certificate file.
	CertFile string

	// KeyFile is the relative path to the key file.
	KeyFile string
}

/*
validateGateway makes sure the gateway adapter is ready to be used properly by a
Blacksmith application.
*/
func validateGateway(g Gateway) error {

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "gateway: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Verify the gateway ID is not empty.
	if g.String() == "" {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Gateway ID must not be empty",
			Path:    []string{"Gateway", "unknown", "String()"},
		})

		return fail
	}

	// We now can add the adapter name to the error message.
	fail.Message = fmt.Sprintf("gateway/%s: Failed to load adapter", g.String())

	// It is impossible to deal with nil options.
	if g.Options() == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Gateway options must not be nil",
			Path:    []string{"Gateway", g.String(), "Options()"},
		})

		return fail
	}

	// If the adapter didn't set an address, use the default one.
	if g.Options().Address == "" {
		g.Options().Address = Defaults.Address
	}

	// Avoid cycles.
	g.Options().Load = nil

	// Return the error if any occurred.
	if len(fail.Validations) > 0 {
		return fail
	}

	return nil
}
