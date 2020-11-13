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

	// Unmarshal the message.
	var err errors.Error
	json.Unmarshal([]byte(entry.Message), &err)

	// Use the given message as the error message.
	entry.Message = err.Message

	// Add a status code if needed.
	if err.StatusCode > 0 {
		entry.Data["statusCode"] = err.StatusCode
	}

	// Add validations if needed.
	if err.Validations != nil && len(err.Validations) > 0 {
		entry.Data["validations"] = err.Validations
	}

	// Add meta info for HTTP response if needed.
	if err.Meta != nil {
		entry.Data["meta"] = err.Meta
	}

	return nil
}
