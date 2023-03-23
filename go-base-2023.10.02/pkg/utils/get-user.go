package utils

import (
	"encoding/json"
	"gobase/pkg/constants"
	"gobase/pkg/errormsg"
	"gobase/pkg/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) (*schemas.TUser, *CustomError) {
	userStr := ctx.Request.Header.Get(string(constants.REQUEST_HEADER_USER))
	var user *schemas.TUser
	if err := json.Unmarshal([]byte(userStr), &user); err != nil {
		return nil, &CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CANNOT_PARSE_USER,
		}
	}
	return user, nil
}
