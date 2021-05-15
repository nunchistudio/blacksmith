package version

import (
	"runtime"
	"strings"
)

/*
Blacksmith is the Blacksmith version number that is being run at the moment,
formatted for semantic versioning.
*/
func Blacksmith() string {
	return "0.17.1"
}

/*
Go is the Go runtime version number that is being run at the moment, formatted
for semantic versioning.
*/
func Go() string {
	return strings.Replace(runtime.Version(), "go", "", 1)
}
