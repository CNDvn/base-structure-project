package routes

import (
	"gobase/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func auth(routes *gin.Engine) {
	controller := controllers.Auth
	route := routes.Group("auth")
	{
		route.POST("/sign-up/username", controller.SignUpWithUsername)
		route.POST("/sign-in/username", controller.SignInWithUsername)
	}
}
