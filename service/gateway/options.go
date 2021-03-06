package gateway

import (
	"github.com/nunchistudio/blacksmith/helper/rest"
	"github.com/nunchistudio/blacksmith/service"
)

/*
Defaults are the defaults options set for the gateway. When not set, these values
will automatically be applied.
*/
var Defaults = &service.Options{
	Address:    ":9090",
	Middleware: rest.Middleware,
	Admin: &service.Admin{
		Enabled:       false,
		WithDashboard: false,
		Middleware:    rest.Middleware,
	},
}
