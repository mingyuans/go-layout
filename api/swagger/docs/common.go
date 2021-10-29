package docs

import "github.com/mingyuans/go-layout/internal/pkg/server"

// ErrorResponse The response means there is an error.
// swagger:response ErrorResponse
type ErrorResponse struct {
	// in:body
	Body struct {
		// Meta summary information
		//
		// Required: true
		Meta server.Meta `json:"meta"`
	}
}
