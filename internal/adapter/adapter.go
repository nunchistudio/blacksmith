package adapter

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"runtime"
	"strings"

	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/version"

	"github.com/hashicorp/go-getter"
	"github.com/mitchellh/go-homedir"
)

/*
dst is the relative path on the local machine where adapters (Go plugins) will
be installed to.
*/
var dst = filepath.Join(".blacksmith", "plugins", version.Blacksmith())

/*
LoadPlugin tries to load a Go plugin given an adapter kind and ID. If the adapter
does not exist on the local machine it will download it.
*/
func LoadPlugin(kind string, from string) (plugin.Symbol, error) {
	fail := &errors.Error{
		Message:     fmt.Sprintf("%s/%s: Failed to load Go plugin", kind, from),
		Validations: []errors.Validation{},
	}

	// Find the working directory.
	wd, err := os.Getwd()
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return nil, fail
	}

	// Find the user's home directory.
	ud, err := homedir.Dir()
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return nil, fail
	}

	// We need to know if the plugin has been found. If so, where is it?
	var isPluginInstalled bool
	var foundAt string

	// A Go plugin name looks like this "store-postgres_darwin-amd64.so" or like
	// this "store-postgres_docker-alpine.so" if an environment has been specified.
	filename := kind + "-" + from

	// Add the appropriate suffix to the file name.
	var suffix string = runtime.GOOS + "-" + runtime.GOARCH
	if os.Getenv("BLACKSMITH_ENV") != "" {
		suffix = os.Getenv("BLACKSMITH_ENV")
	}

	// Only deal with ".so" files.
	filename += "_" + suffix + ".so"

	// Possible directories to look at. We first look at the root directory, then
	// at the user's home directory, and finally at the working directory.
	locations := []string{
		filepath.Join(dst, filename),
		filepath.Join(ud, dst, filename),
		filepath.Join(wd, dst, filename),
	}

	// Look for the plugin file. If it has been found, save its location and break
	// the loop.
	for _, location := range locations {
		f, _ := os.Stat(location)
		if f != nil {
			foundAt = location
			isPluginInstalled = true
			break
		}
	}

	// If the plugin has not been found we need to download it. It will be saved
	// in the working directory.
	if isPluginInstalled == false {
		foundAt = filepath.Join(wd, dst, filename)
		err = downloadPlugin(wd, kind, from)
		if err != nil {
			return nil, err
		}
	}

	// We can now try to open the Go plugin.
	adapter, err := plugin.Open(foundAt)
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Failed to open Go plugin: check version compatibility",
			Path:    []string{kind, from, "Options", "From"},
		})

		return nil, fail
	}

	// Lookup the Go plugin's symbol.
	symbol, err := adapter.Lookup("New")
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Failed to find the adapter function `New`",
			Path:    []string{kind, from, "Options", "From"},
		})

		return nil, fail
	}

	// If we made it here we can return the plugin's symbol.
	return symbol, nil
}

/*
downloadPlugin downloads an adapter (a Go plugin) from the GitHub repository given
a kind and an adapter ID.
*/
func downloadPlugin(wd string, kind string, from string) error {

	// Start to create the file name.
	filename := kind + "-" + from

	// The defaults Go plugins can be downloaded for supported environments and
	// architectures.
	var suffix string = runtime.GOOS + "-" + runtime.GOARCH
	var src = "https://github.com/nunchistudio/blacksmith/releases/download/" + version.Blacksmith() + "/"

	// If an environment is specified, download from its own repository and use the
	// appropriate suffix.
	if os.Getenv("BLACKSMITH_ENV") != "" {
		suffix = os.Getenv("BLACKSMITH_ENV")

		platform := strings.Split(suffix, "-")
		src = "https://github.com/nunchistudio/blacksmith-" + platform[0] + "/releases/download/" + version.Blacksmith() + "/"
	}

	// Add the suffix with the ".zip".
	filename += "_" + suffix + ".zip"

	// Determine the source URL to download.
	link := src + filename

	// Get the adapter's archive and extract it at the right destination.
	client := &getter.Client{
		Src:  link,
		Dst:  dst,
		Pwd:  wd,
		Mode: getter.ClientModeAny,
	}

	// Return if any error occurred.
	if err := client.Get(); err != nil {
		return &errors.Error{
			Message: fmt.Sprintf("%s/%s: Failed to load Go plugin", kind, from),
			Validations: []errors.Validation{
				{
					Message: "Failed to download Go plugin from repository: is the environment supported?",
					Path:    []string{kind, from, "Options", "From"},
				},
			},
		}
	}

	// No error occurred.
	return nil
}
