package routes

import (
	"fmt"
	"gobase/pkg/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	routes := gin.Default()

	routes.Use(middlewares.CORSMiddleware())

	// setup format logger
	routes.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format(time.RFC1123),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	video(routes)
	auth(routes)
	image(routes)
	return routes
}
