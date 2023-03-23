package routes

import (
	"gobase/pkg/controllers"
	"gobase/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func image(routes *gin.Engine) {
	controller := controllers.Image
	route := routes.Group("images")
	{
		route.POST("", middlewares.AuthMiddleware(), controller.Create)
		route.GET("", middlewares.AuthMiddleware(), controller.Get)
		route.GET("/:id", middlewares.AuthMiddleware(), controller.GetById)
	}
}
