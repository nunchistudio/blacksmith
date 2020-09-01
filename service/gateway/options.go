package gateway

import (
	"context"

	"github.com/nunchistudio/blacksmith/helper/rest"
	"github.com/nunchistudio/blacksmith/service"
)

/*
Defaults are the defaults options set for the gateway. When not set, these values
will automatically be applied.
*/
var Defaults = &service.Options{
	Context:    context.Background(),
	Address:    ":8080",
	Middleware: rest.Middleware,
}
