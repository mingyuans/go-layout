package controller_v1_user

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/log"
	"github.com/mingyuans/go-layout/internal/pkg/server"
)

func (u *UserController) Get(c *gin.Context) {
	log.Info("Get user function called.")

	user, err := u.srv.Users().Get(c, "sss")
	server.WriteResponse(c, user, err)
}
