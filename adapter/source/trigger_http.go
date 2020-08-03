package source

import (
	"net/http"
	"strings"

	"github.com/nunchistudio/blacksmith/helper/errors"
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

/*
validMethods is used to make sure the HTTP method is valid.
*/
var validMethods = map[string]bool{
	"DELETE": true,
	"GET":    true,
	"PATCH":  true,
	"POST":   true,
	"PUT":    true,
}

/*
validate validates a HTTP route.
*/
func (r *Route) validate(source string, trigger string) []errors.Validation {
	validations := []errors.Validation{}

	if r == nil {
		validations = append(validations, errors.Validation{
			Message: "HTTP route must not be nil",
			Path:    []string{"Source", source, "Triggers()", trigger, "Mode()", "UsingHTTP"},
		})

		return validations
	}

	if r.Path == "" {
		validations = append(validations, errors.Validation{
			Message: "HTTP route path must not be empty",
			Path:    []string{"Source", source, "Triggers()", trigger, "Mode()", "UsingHTTP", "Path"},
		})
	} else if r.Path == "/" {
		validations = append(validations, errors.Validation{
			Message: "HTTP route path must not be a wildcard path",
			Path:    []string{"Source", source, "Triggers()", trigger, "Mode()", "UsingHTTP", "Path"},
		})
	} else if strings.HasPrefix(r.Path, "/") == false {
		validations = append(validations, errors.Validation{
			Message: "HTTP route path must start with '/'",
			Path:    []string{"Source", source, "Triggers()", trigger, "Mode()", "UsingHTTP", "Path"},
		})
	} else if strings.HasPrefix(r.Path, "/api") == true {
		validations = append(validations, errors.Validation{
			Message: "HTTP route path must not start with '/api'",
			Path:    []string{"Source", source, "Triggers()", trigger, "Mode()", "UsingHTTP", "Path"},
		})
	}

	if r.Methods == nil {
		validations = append(validations, errors.Validation{
			Message: "HTTP route methods must not be nil",
			Path:    []string{"Source", source, "Triggers()", trigger, "Mode()", "UsingHTTP", "Methods"},
		})

		return validations
	}

	for _, method := range r.Methods {
		method = strings.ToUpper(method)
		if _, valid := validMethods[method]; !valid {
			validations = append(validations, errors.Validation{
				Message: "HTTP route method not valid",
				Path:    []string{"Source", source, "Triggers()", trigger, "Mode()", "UsingHTTP", "Methods", method},
			})
		}
	}

	return validations
}
