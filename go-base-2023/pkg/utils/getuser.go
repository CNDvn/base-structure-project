package utils

import (
	"basego/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserReq(ctx *gin.Context) (*models.User, error) {
	userIdStr := ctx.Request.Header.Get(string(HEADER_USER_ID))
	userId, err := strconv.ParseUint(userIdStr, 10, 64)

	if err != nil {
		return nil, err
	}

	var user models.User
	result := MySqlDB.Find(&user, "id = ?", userId)
	return &user, result.Error
}
