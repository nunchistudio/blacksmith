package source

/*
ValidMethods is used to make sure the HTTP method is valid.
*/
var ValidMethods = map[string]bool{
	"DELETE": true,
	"GET":    true,
	"PATCH":  true,
	"POST":   true,
	"PUT":    true,
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
	// Example: "/api/user"
	Path string `json:"path"`
}
