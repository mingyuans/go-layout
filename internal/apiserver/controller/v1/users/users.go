package controller_v1_user

import "github.com/mingyuans/go-layout/internal/apiserver/service/v1"

// UserController create a user handler used to handle request for user resource.
type UserController struct {
	srv service_v1.Service
}

// NewUserController creates a user handler.
func NewUserController() *UserController {
	return &UserController{
		srv: service_v1.NewService(),
	}
}
