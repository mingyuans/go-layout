package docs

import service_v1_users "github.com/mingyuans/go-layout/internal/iam-apiserver/service/v1/users"

// UserResponse User response.
// swagger:response UserResponse
type UserResponse struct {
	// in:body
	Body struct {
		Data service_v1_users.User `json:"data"`
	}
}

// UserRequest PUT or POST users request.
// swagger:parameters UserRequest
type UserRequest struct {
	service_v1_users.User `json:"user"`
}

// swagger:route GET /users/{name} users getUserRequest
//
// Get details for specified user.
//
// Get details for specified user according to input parameters.
//
//     Responses:
//       default: ErrorResponse
//       200: UserResponse

// swagger:route POST /users users UserRequest
//
// Create a user resource.
//
//
//     Responses:
//       default: ErrorResponse
//       200: UserResponse
