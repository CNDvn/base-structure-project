package auth

import (
	"basego/pkg/models"
	"basego/pkg/utils"
	"context"
	"net/http"
)

type AuthService struct{}

func (AuthService) SignUp(idToken string) (*models.User, *utils.CustomError) {
	client, err := utils.FirebaseApp.Auth(context.Background())
	if err != nil {
		return nil, &utils.CustomError{
			Status:  http.StatusInternalServerError,
			Message: "Firebase Auth fail",
		}
	}
	tokenVerified, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: "Verify token fail",
		}
	}

	uid := tokenVerified.Claims["user_id"].(string)
	email := tokenVerified.Claims["email"].(string)
	name := tokenVerified.Claims["name"].(string)
	role := string(utils.ROLE_USER)

	user := &models.User{
		UID:   uid,
		Email: email,
		Name:  name,
		Role:  role,
	}
	if err := usersService.CreateUser(user); err != nil {
		return nil, err
	} else {
		return user, nil
	}

}
