package logger

import (
	"encoding/json"

	"github.com/nunchistudio/blacksmith/helper/errors"

	"github.com/sirupsen/logrus"
)

/*
UsingError respect the logrus Hook interface and allows the Blacksmith logger
to format errors (using the helper/errors package) across adapters and in the
application.
*/
type UsingError struct{}

/*
Levels return the level used by the hook. Use for error level only.
*/
func (hook *UsingError) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
	}
}

/*
Fire format the given message to an appropriate error.
*/
func (hook *UsingError) Fire(entry *logrus.Entry) error {

	// Make sure the data is not nil.
	if entry.Data == nil {
		entry.Data = logrus.Fields{}
	}

	// Unmarshal the message. Do not continue if the underlying error type is not
	// known.
	var fail errors.Error
	err := json.Unmarshal([]byte(entry.Message), &fail)
	if err != nil {
		return nil
	}

	// Use the given message as the error message.
	entry.Message = fail.Message

	// Add a status code if needed.
	if fail.StatusCode > 0 {
		entry.Data["statusCode"] = fail.StatusCode
	}

	// Add validations if needed.
	if fail.Validations != nil && len(fail.Validations) > 0 {
		entry.Data["validations"] = fail.Validations
	}

	// Add meta info for HTTP response if needed.
	if fail.Meta != nil {
		entry.Data["meta"] = fail.Meta
	}

	return nil
}
