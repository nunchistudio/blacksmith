package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*
DefaultLogger is the default logger used in non-production environments when none
where passed when creating the application.
*/
var DefaultLogger = &logrus.Logger{
	Out:   os.Stdout,
	Level: logrus.DebugLevel,
	Hooks: logrus.LevelHooks{},
	Formatter: &logrus.TextFormatter{
		FullTimestamp: true,
	},
	ExitFunc: os.Exit,
}

/*
DefaultLoggerInProduction is the default logger used in production environments
when none where passed when creating the application.
*/
var DefaultLoggerInProduction = &logrus.Logger{
	Out:       os.Stderr,
	Level:     logrus.WarnLevel,
	Hooks:     logrus.LevelHooks{},
	Formatter: &logrus.JSONFormatter{},
	ExitFunc:  os.Exit,
}
