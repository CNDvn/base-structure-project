package routes

import (
	"basego/pkg/auth"

	"github.com/gin-gonic/gin"
)

func authRoute(routes *gin.Engine, nameGroup string) {
	authRoute := routes.Group(nameGroup)
	{
		authRoute.POST("/", auth.SignUp)
	}
}
