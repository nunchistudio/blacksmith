package rest

import (
	"encoding/json"
	"net/http"

	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
ErrorNotFound handles HTTP 404 error responses. When called, the calling function
must return to avoid writing several times on the HTTP response writer.
*/
func ErrorNotFound(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	body := errors.Error{
		StatusCode: 404,
		Message:    "Not Found",
	}

	r, _ := json.Marshal(body)
	res.WriteHeader(body.StatusCode)
	res.Write(r)
}

/*
ErrorMethodNotAllowed handles HTTP 405 error responses. When called, the calling
function must return to avoid writing several times on the HTTP response writer.
*/
func ErrorMethodNotAllowed(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	body := errors.Error{
		StatusCode: 405,
		Message:    "Method Not Allowed",
	}

	r, _ := json.Marshal(body)
	res.WriteHeader(body.StatusCode)
	res.Write(r)
}

/*
ErrorInternal handles HTTP 500 error responses. When called, the calling function
must return to avoid writing several times on the HTTP response writer.
*/
func ErrorInternal(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	body := errors.Error{
		StatusCode: 500,
		Message:    "Internal Server Error",
	}

	r, _ := json.Marshal(body)
	res.WriteHeader(body.StatusCode)
	res.Write(r)
}
