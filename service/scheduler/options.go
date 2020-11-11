package scheduler

import (
	"context"

	"github.com/nunchistudio/blacksmith/helper/rest"
	"github.com/nunchistudio/blacksmith/service"
)

/*
Defaults are the defaults options set for the scheduler. When not set, these values
will automatically be applied.
*/
var Defaults = &service.Options{
	Context:    context.Background(),
	Address:    ":9091",
	Middleware: rest.Middleware,
	Admin: &service.Admin{
		Enabled:    true,
		Middleware: rest.Middleware,
	},
}
