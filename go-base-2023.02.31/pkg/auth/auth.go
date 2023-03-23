package auth

import (
	firebaseContext "context"
	"net/http"
	"strconv"

	"basego/pkg/modules/users"
	"basego/pkg/modules/users/userdto"
	"basego/pkg/utils"

	"github.com/gin-gonic/gin"
)

var usersService = users.UsersService{}
var authService = AuthService{}

func Auth(ctx *gin.Context) {
	tokenString := ctx.GetHeader(string(utils.HEADER_AUTHORIZATION))

	client, err := utils.FirebaseApp.Auth(firebaseContext.Background())
	if err != nil {
		utils.CustomError{
			Status:  http.StatusUnauthorized,
			Message: err.Error()}.Send(ctx)
		return
	}
	token, err := client.VerifyIDToken(firebaseContext.Background(), tokenString)
	if err != nil {
		utils.CustomError{
			Status:  http.StatusUnauthorized,
			Message: err.Error()}.Send(ctx)
		return
	}

	email := token.Claims["email"].(string)
	user, errFind := usersService.FindUserByEmail(email)
	if errFind != nil {
		errFind.Send(ctx)
		return
	}
	ctx.Request.Header.Add(string(utils.HEADER_USER_EMAIL), email)
	ctx.Request.Header.Add(string(utils.HEADER_USER_ID), strconv.FormatInt(int64(user.ID), 10))
	ctx.Request.Header.Add(string(utils.HEADER_USER_UID), user.UID)
	ctx.Request.Header.Add(string(utils.HEADER_USER_ROLE), string(user.Role))
	ctx.Next()
}

func SignUp(ctx *gin.Context) {
	idToken := ctx.GetHeader(string(utils.HEADER_AUTHORIZATION))

	if user, err := authService.SignUp(idToken); err != nil {
		err.Send(ctx)
	} else {
		userDto := userdto.UserInfo{}
		userDto.MapFrom(*user)
		utils.CustomResponse{
			Status:   http.StatusCreated,
			Metadata: userDto,
		}.Send(ctx)
	}
}
