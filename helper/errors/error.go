package errors

/*
Error is used across every Blacksmith packages to normalize error handling. Error
implements the standard error interface.
*/
type Error struct {

	// StatusCode is the HTTP status code returned by the caller.
	StatusCode int `json:"statusCode,omitempty"`

	// Message is the error message.
	Message string `json:"message"`

	// Validations is a list of validation errors.
	Validations []Validation `json:"validations,omitempty"`

	// Meta includes meta details about the error. It is only used when dealing
	// with HTTP errors to provide a consistent HTTP response across adapters.
	Meta *Meta `json:"meta,omitempty"`
}

/*
Validation gives additional info about the parent error. An error can have multiple
error validations.
*/
type Validation struct {

	// Message is the error message.
	Message string `json:"message"`

	// Path gives more details about where the error happened.
	Path []string `json:"path"`
}

/*
Error returns a stringified representation of the marshalled error. This allows
to use Error as a standard error across Blacksmith packages and can be unmarshall
by the logrus Logger with custom hooks.
*/
func (err *Error) Error() string {
	msg := err.Message

	if err.Validations != nil && len(err.Validations) > 0 {
		msg += "\n"
		for _, validation := range err.Validations {
			msg += "  - " + validation.Message + "\n"
		}
	}

	return msg
}
