package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/mingyuans/go-layout/internal/apiserver/controller/v1/users"
)

//Install APIs
func installControllers(g *gin.Engine) {
	v1Group := g.Group("/v1")
	installUsersController(v1Group)
}

//不同 resources 的 API 单独一个 function 和 module
func installUsersController(g *gin.RouterGroup) {
	v1Users := g.Group("/users")
	userController := controller_v1_user.NewUserController()

	v1Users.GET("", userController.GetUsers)
	v1Users.GET(":userid", userController.Get)

	v1Users.POST("", userController.CreateUser)
}
