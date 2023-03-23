package utils

import "github.com/gin-gonic/gin"

type CustomResponse struct {
	Status   int
	Message  string
	Metadata any
	Option   any
}

func (res CustomResponse) Send(ctx *gin.Context) {
	ctx.JSON(res.Status, res)
}
