package users

import (
	"basego/pkg/models"
	"basego/pkg/utils"
	"net/http"
)

var usersRepo = UsersRepository{}

type UsersService struct {
}

func (UsersService) CreateUser(user *models.User) *utils.CustomError {
	if err := usersRepo.CreateUser(user); err != nil {
		return &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		return nil
	}
}

func (UsersService) FindUserByEmail(email string) (*models.User, *utils.CustomError) {
	if user, err := usersRepo.FindUserByEmail(email); err != nil {
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		return user, nil
	}
}

func (UsersService) FindUserById(id uint) (*models.User, *utils.CustomError) {
	if user, err := usersRepo.FindUserById(id); err != nil {
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		return user, nil
	}
}
