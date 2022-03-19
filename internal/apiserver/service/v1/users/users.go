package service_v1_users

import (
	"context"
	"github.com/marmotedu/errors"
	"github.com/mingyuans/go-layout/internal/pkg/code"
	"github.com/mingyuans/go-layout/pkg/log"
)

type UserService interface {
	Get(ctc context.Context, username string) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
	Create(ctx context.Context, username string) (*User, error)
}

type userServiceImpl struct {
}

func NewUserService() UserService {
	return &userServiceImpl{}
}

func (u *userServiceImpl) GetUsers(ctx context.Context) ([]User, error) {
	var users = []User{
		{
			Nickname: "abc",
			Password: "111",
			Email:    "abc@sd.com",
			Phone:    "123",
		},
		{
			Nickname: "abcdef",
			Password: "111222",
			Email:    "abc@sd22.com",
			Phone:    "123111",
		},
	}

	return users, nil
}

func (u *userServiceImpl) Get(ctc context.Context, username string) (*User, error) {
	log.L(ctc).Info("call userServiceImpl Get method.")
	return nil, errors.WithCode(code.ErrUserNotFound, "Can't find the user.")
}

func (u *userServiceImpl) Create(ctx context.Context, username string) (*User, error) {
	return nil, errors.WithCode(code.ErrUserAlreadyExist, "User is existed.")
}
