package logger

import (
	"github.com/sirupsen/logrus"
)

/*
New returns the appropriate logrus Logger given the environment.
*/
func New() *logrus.Logger {
	var log = DefaultLogger
	log.AddHook(&UsingError{})

	return log
}
