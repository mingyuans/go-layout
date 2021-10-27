package controller_v1_user

import (
	"github.com/gin-gonic/gin"
	"github.com/mingyuans/go-layout/internal/pkg/server"
	"github.com/mingyuans/go-layout/pkg/log"
)

func (u *UserController) CreateUser(c *gin.Context) {
	log.L(c).Info("Get users function called.")

	users, err := u.srv.Users().Create(c, "sss")
	server.NewRestfulResponseBuilder(c).
		Data(users).
		Error(err).
		SendJSON()
}
