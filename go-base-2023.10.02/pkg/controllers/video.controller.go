package controllers

import (
	"gobase/pkg/reqdto"
	"gobase/pkg/services"
	"gobase/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TVideo struct{}

func (t *TVideo) Create(ctx *gin.Context) {
	var dto *reqdto.TCreateVideoReqDto
	// ctx.BindJSON(&dto)
	ctx.Bind(&dto)

	user, err := utils.GetUser(ctx)
	if err != nil {
		err.Send(ctx)
		return
	}

	if resObj, err := services.Video.Create(dto, user); err != nil {
		err.Send(ctx)

	} else {
		utils.CustomResponse{
			Status:   http.StatusCreated,
			Metadata: resObj,
		}.Send(ctx)
	}

}

func (t *TVideo) Get(ctx *gin.Context) {
	var pagination utils.TPagination
	ctx.BindQuery(&pagination)

	user, err := utils.GetUser(ctx)
	if err != nil {
		err.Send(ctx)
		return
	}

	if resObj, err := services.Video.Get(user, &pagination); err != nil {
		err.Send(ctx)
	} else {
		utils.CustomResponse{
			Status:   http.StatusCreated,
			Metadata: resObj,
		}.Send(ctx)
	}
}

func (t *TVideo) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := utils.GetUser(ctx)
	if err != nil {
		err.Send(ctx)
		return
	}

	if resObj, err := services.Video.GetById(id, user); err != nil {
		err.Send(ctx)
	} else {
		utils.CustomResponse{
			Status:   http.StatusCreated,
			Metadata: resObj,
		}.Send(ctx)
	}
}
