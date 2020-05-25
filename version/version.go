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
	return "v0.10.0"
}

/*
Go is the Go runtime version number that is being run at the moment, formatted
for semantic versioning.
*/
func Go() string {
	v := runtime.Version()
	v = strings.Replace(v, "go", "v", 1)

	return v
}
