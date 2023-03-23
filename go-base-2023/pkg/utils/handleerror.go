package utils

import "github.com/gin-gonic/gin"

type CustomError struct {
	Status  int
	Message string
}

func (err CustomError) Send(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(err.Status, err)
	ctx.Abort()
}
