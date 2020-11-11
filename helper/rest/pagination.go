package rest

/*
Pagination holds the pagination details when looking for entries in the REST API.
*/
type Pagination struct {

	// Current is the current page.
	Current uint16 `json:"current"`

	// Previous is the previous page. It will be nil if there is no previous page.
	Previous *uint16 `json:"previous"`

	// Next is the next page. It will be nil if there is no next page.
	Next *uint16 `json:"next"`

	// First is the first page. It will always be 1.
	First uint16 `json:"first"`

	// Last is the last page.
	Last uint16 `json:"last"`
}
