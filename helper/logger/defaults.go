package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*
Default is the default logger used by Blacksmith applications.
*/
var Default = &logrus.Logger{
	Out:   os.Stdout,
	Level: logrus.DebugLevel,
	Hooks: logrus.LevelHooks{},
	Formatter: &logrus.TextFormatter{
		FullTimestamp: true,
	},
	ExitFunc: os.Exit,
}

/*
CLI is the logger used by the Blacksmith CLI for managing logs in a non-running
application.
*/
var CLI = &logrus.Logger{
	Out:   os.Stdout,
	Level: logrus.InfoLevel,
	Hooks: logrus.LevelHooks{},
	Formatter: &logrus.TextFormatter{
		DisableTimestamp: true,
	},
	ExitFunc: os.Exit,
}
