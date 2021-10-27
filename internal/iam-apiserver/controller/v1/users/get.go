package controller_v1_user

import (
	"github.com/gin-gonic/gin"
	"github.com/mingyuans/go-layout/internal/pkg/server"
	"github.com/mingyuans/go-layout/pkg/log"
)

func (u *UserController) Get(c *gin.Context) {
	log.L(c).Info("Get user function called.")

	user, err := u.srv.Users().Get(c, "sss")
	server.NewRestfulResponseBuilder(c).
		Data(user).
		Error(err).
		SendJSON()
}

func (u *UserController) GetUsers(c *gin.Context) {
	log.L(c).Info("Get users function called.")

	users, err := u.srv.Users().GetUsers(c)
	server.NewRestfulResponseBuilder(c).
		Data(users).
		Error(err).
		SendJSON()
}
