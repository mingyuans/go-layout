package service_v1

import users "github.com/mingyuans/go-layout/internal/apiserver/service/v1/users"

type Service interface {
	Users() users.UserService
}

type service struct {
}

func (s service) Users() users.UserService {
	return users.NewUserService()
}

func NewService() Service {
	return &service{}
}
