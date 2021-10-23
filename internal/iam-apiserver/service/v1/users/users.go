package service_v1_users

import (
	"context"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/log"
)

type UserService interface {
	Get(ctc context.Context, username string) (*User, error)
}

type userServiceImpl struct {
}

func NewUserService() UserService {
	return &userServiceImpl{}
}

func (u *userServiceImpl) Get(ctc context.Context, username string) (*User, error) {
	log.Info("call userServiceImpl Get method.")
	return nil, errors.New("Testing error message.")
}
