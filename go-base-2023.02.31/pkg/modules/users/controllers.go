package users

import (
	"basego/pkg/modules/users/userdto"
	"basego/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var usersService = UsersService{}

type UsersController struct {
}

func (UsersController) GetMyInfo(ctx *gin.Context) {
	if user, err := utils.GetUserReq(ctx); err != nil {
		utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}.Send(ctx)
		return
	} else {
		userDto := userdto.UserInfo{}
		userDto.MapFrom(*user)

		utils.CustomResponse{
			Status:   http.StatusOK,
			Metadata: userDto,
		}.Send(ctx)
		return
	}
}

func (UsersController) FindUserByEmail(ctx *gin.Context) {
	email := ctx.Request.URL.Query().Get("email")

	if user, err := usersService.FindUserByEmail(email); err != nil {
		err.Send(ctx)
		return
	} else {
		userDto := userdto.UserInfo{}
		userDto.MapFrom(*user)

		utils.CustomResponse{
			Status:   http.StatusOK,
			Metadata: userDto,
		}.Send(ctx)
		return
	}
}
