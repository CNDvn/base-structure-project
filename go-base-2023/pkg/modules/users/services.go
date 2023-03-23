package users

import (
	"basego/pkg/models"
)

var usersRepo = UsersRepository{}

type UsersService struct {
}

func (UsersService) CreateUser(user models.User) (int64, error) {
	return usersRepo.CreateUser(user)
}

func (UsersService) FindUserByEmail(email string) (models.User, error) {
	return usersRepo.FindUserByEmail(email)
}

func (UsersService) FindUserById(id uint) (models.User, error) {
	return usersRepo.FindUserById(id)
}
