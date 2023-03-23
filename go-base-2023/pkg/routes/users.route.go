package routes

import (
	"basego/pkg/auth"
	"basego/pkg/modules/users"

	"github.com/gin-gonic/gin"
)

func userRoute(routes *gin.Engine, nameGroup string) {
	usersController := users.UsersController{}
	usersRoute := routes.Group(nameGroup)
	{
		usersRoute.GET("/me", auth.Auth, usersController.GetMyInfo)
		usersRoute.GET("/", auth.Auth, usersController.FindUserByEmail)
	}
}
