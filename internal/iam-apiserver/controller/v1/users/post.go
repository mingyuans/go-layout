package controller_v1_user

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/log"
	"github.com/mingyuans/go-layout/internal/pkg/server"
)

func (u *UserController) CreateUser(c *gin.Context) {
	log.Info("Get users function called.")

	users, err := u.srv.Users().Create(c, "sss")
	server.NewRestfulResponseBuilder(c).
		Data(users).
		Error(err).
		SendJSON()
}
