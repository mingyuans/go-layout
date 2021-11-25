package controller_v1_user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/mingyuans/go-layout/internal/pkg/code"
	"github.com/mingyuans/go-layout/internal/pkg/server"
	"github.com/mingyuans/go-layout/pkg/log"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

// Get GET users/userid
func (u *UserController) Get(c *gin.Context) {
	userid := c.Param("userid")
	log.L(c).Info("Get user.", zap.String("userid", userid))

	errs := validator.New().Var(userid,"required,len=32")
	if errs != nil {
		fmt.Println(errs)
		server.NewRestfulResponseBuilder(c).
			Error(errors.WithCode(code.ErrValidation,"The user id {%s} is invalid.",userid)).
			SendJSON()
		return
	}

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
