package controllers

import (
	"gobase/pkg/reqdto"
	"gobase/pkg/services"
	"gobase/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TAuth struct{}

func (t *TAuth) SignUpWithUsername(ctx *gin.Context) {
	var dto *reqdto.TSignUpWithUsername
	ctx.BindJSON(&dto)

	if signUpSuccess, err := services.Auth.SignUp(dto); err != nil {
		err.Send(ctx)
		return
	} else {
		utils.CustomResponse{
			Status:   http.StatusCreated,
			Metadata: signUpSuccess,
		}.Send(ctx)
	}
}

func (t *TAuth) SignInWithUsername(ctx *gin.Context) {
	var dto *reqdto.TSignInWithUsername
	ctx.BindJSON(&dto)

	if signInSuccess, err := services.Auth.SignIn(dto); err != nil {
		err.Send(ctx)
		return
	} else {
		utils.CustomResponse{
			Status:   http.StatusOK,
			Metadata: signInSuccess,
		}.Send(ctx)
	}
}
