package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*
New returns the appropriate logrus Logger given the environment.
*/
func New() *logrus.Logger {
	var log = DefaultLogger
	if os.Getenv("BLACKSMITH_ENV") == "production" {
		log = DefaultLoggerInProduction
	}

	log.AddHook(&UsingError{})

	return log
}
