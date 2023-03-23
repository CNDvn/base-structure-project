package middlewares

import (
	"encoding/json"
	"gobase/pkg/constants"
	"gobase/pkg/errormsg"
	"gobase/pkg/services"
	"gobase/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.Request.Header.Get(string(constants.REQUEST_HEADER_AUTHORIZATION))
		payload, err := services.Auth.ParseJwt(tokenString)
		if err != nil {
			utils.PrintLog("func AuthMiddleware() gin.HandlerFunc", err.Error())
			utils.CustomError{
				Status:  http.StatusUnauthorized,
				Message: errormsg.TOKEN_INVALID,
			}.Send(ctx)
			return
		}

		if payload == nil {
			utils.CustomError{
				Status:  http.StatusUnauthorized,
				Message: errormsg.PAYLOAD_INVALID,
			}.Send(ctx)
			return
		}

		user, err := services.User.FindOne(bson.M{"_id": payload.UID})
		if err != nil {
			utils.PrintLog("func AuthMiddleware() gin.HandlerFunc", err.Error())
			utils.CustomError{
				Status:  http.StatusUnauthorized,
				Message: errormsg.NOT_FOUND_USER,
			}.Send(ctx)
			return
		}

		if user == nil {
			utils.PrintLog("func AuthMiddleware() gin.HandlerFunc", "user is nil")
			utils.CustomError{
				Status:  http.StatusUnauthorized,
				Message: errormsg.NOT_FOUND_USER,
			}.Send(ctx)
			return
		}

		userByte, err := json.Marshal(user)
		if err != nil {
			utils.PrintLog("func AuthMiddleware() gin.HandlerFunc", "Can not parse user")
			utils.CustomError{
				Status:  http.StatusUnauthorized,
				Message: errormsg.NOT_FOUND_USER,
			}.Send(ctx)
			return
		}

		ctx.Request.Header.Set(string(constants.REQUEST_HEADER_USER), string(userByte))
		ctx.Next()
	}
}
