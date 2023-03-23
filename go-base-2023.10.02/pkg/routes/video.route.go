package routes

import (
	"gobase/pkg/controllers"
	"gobase/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func video(routes *gin.Engine) {
	controller := controllers.Video
	route := routes.Group("videos")
	{
		route.POST("", controller.Create)
		route.GET("", middlewares.AuthMiddleware(), controller.Get)
		route.GET("/:id", middlewares.AuthMiddleware(), controller.GetById)
	}
}
