package controllers

import (
	"gobase/pkg/reqdto"
	"gobase/pkg/services"
	"gobase/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TImage struct{}

func (t *TImage) Create(ctx *gin.Context) {
	var dto *reqdto.TCreateImageReqDto
	ctx.Bind(&dto)

	user, err := utils.GetUser(ctx)
	if err != nil {
		err.Send(ctx)
		return
	}

	if resObj, err := services.Image.Create(dto, user); err != nil {
		err.Send(ctx)

	} else {
		utils.CustomResponse{
			Status:   http.StatusCreated,
			Metadata: resObj,
		}.Send(ctx)
	}

}

func (t *TImage) Get(ctx *gin.Context) {
	var pagination utils.TPagination
	ctx.BindQuery(&pagination)

	user, err := utils.GetUser(ctx)
	if err != nil {
		err.Send(ctx)
		return
	}

	if resObj, err := services.Image.Get(user, &pagination); err != nil {
		err.Send(ctx)
	} else {
		utils.CustomResponse{
			Status:   http.StatusCreated,
			Metadata: resObj,
		}.Send(ctx)
	}
}

func (t *TImage) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := utils.GetUser(ctx)
	if err != nil {
		err.Send(ctx)
		return
	}

	if resObj, err := services.Image.GetById(id, user); err != nil {
		err.Send(ctx)
	} else {
		utils.CustomResponse{
			Status:   http.StatusCreated,
			Metadata: resObj,
		}.Send(ctx)
	}
}
