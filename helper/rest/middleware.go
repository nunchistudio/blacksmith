package rest

import (
	"net/http"
)

/*
Middleware is the default HTTP middleware used in the gateway and scheduler
adapters. It can be replaced by any in-house function to control the
request / response lifecycle of your application.
*/
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(res, req)
	})
}
