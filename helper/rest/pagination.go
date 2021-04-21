package rest

import (
	"math"
)

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

/*
Paginate returns the appropriate pagination details given the count, offset, and
limit of a query.
*/
func Paginate(count, offset, limit uint16) *Pagination {
	p := &Pagination{
		First: 1,
	}

	// Calcul the current page number.
	p.Current = uint16(math.Ceil(float64(offset) / float64(limit)))
	p.Current += p.First

	// Calcul the last page number.
	p.Last = uint16(math.Ceil(float64(count) / float64(limit)))

	// Calcul the previous page number if applicable.
	if (p.Current - 1) >= p.First {
		p.Previous = new(uint16)
		*p.Previous = p.Current - 1
	}

	// Calcul the next page number if applicable.
	if (p.Current + 1) <= p.Last {
		p.Next = new(uint16)
		*p.Next = p.Current + 1
	}

	return p
}
