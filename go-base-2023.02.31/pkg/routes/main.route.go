package routes

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	routes := gin.Default()

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

	authRoute(routes, "auth")
	userRoute(routes, "users")

	return routes
}
