package controller

import "prospect/mobile-physician-api/pkg/dao"

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response paginationResponse
type ResponsePaginator struct {
	page         int
	ItemsPerPage int

	// in: body
	Members      [] dao.Entity
}
