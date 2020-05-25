package adapter

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"runtime"

	"github.com/nunchistudio/blacksmith/helper/errors"
	"github.com/nunchistudio/blacksmith/version"

	"github.com/hashicorp/go-getter"
)

/*
src is the HTTP server where releases can be downloaded from.
*/
var src = "https://github.com/nunchistudio/blacksmith/releases/download/" + version.Blacksmith() + "/"

/*
dst is the destination on the local machine where adapters (Go plugins) will be
installed to.
*/
var dst = ".blacksmith/plugins/" + version.Blacksmith() + "/"

/*
LoadPlugin tries to load a Go plugin given an adapter kind and ID. If the adapter
does not exist on the local machine it will download it.
*/
func LoadPlugin(ctx context.Context, kind string, from string) (plugin.Symbol, error) {
	fail := &errors.Error{
		Message:     fmt.Sprintf("%s/%s: Failed to load Go plugin", kind, from),
		Validations: []errors.Validation{},
	}

	// Find the current working directory.
	wd, err := os.Getwd()
	if err != nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: err.Error(),
		})

		return nil, fail
	}

	// A Go plugin name looks like this: "destination-postgres.so".
	filename := kind + "-" + from + ".so"

	// Save the real path of the plugin.
	file := filepath.Join(wd, dst, filename)

	// If the Go plugin doesn't exist, download it.
	if _, err := os.Stat(file); os.IsNotExist(err) {
		err = downloadPlugin(ctx, wd, kind, from)
		if err != nil {
			return nil, err
		}
	}

	// Try to open the Go plugin.
	adapter, err := plugin.Open(file)
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
func downloadPlugin(ctx context.Context, wd string, kind string, from string) error {

	// Determine the platform.
	platform := runtime.GOOS + "-" + runtime.GOARCH

	// A Go plugin archive looks like this: "destination-postgres_darwin-amd64.zip".
	zipname := kind + "-" + from + "_" + platform + ".zip"

	// Determine the source URL to download.
	link := src + zipname

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
					Message: "Failed to download Go plugin from repository",
					Path:    []string{kind, from, "Options", "From"},
				},
			},
		}
	}

	// No error occurred.
	return nil
}
