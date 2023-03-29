package users

import (
	"basego/pkg/models"
	"basego/pkg/utils"
	"errors"
)

type UsersRepository struct{}

func (UsersRepository) CreateUser(user *models.User) error {
	result := utils.MySqlDB.Create(&user)
	return result.Error
}

func (UsersRepository) GetUser(userId uint) ([]models.User, error) {
	var user []models.User
	result := utils.MySqlDB.Find(&user)
	return user, result.Error
}

func (UsersRepository) FindUserByEmail(email string) (*models.User, error) {
	var user *models.User
	result := utils.MySqlDB.Find(&user, "email = ?", email)
	if result.RowsAffected == 0 {
		return user, errors.New("Not found user")
	}
	return user, result.Error
}

func (UsersRepository) FindUserById(id uint) (*models.User, error) {
	var user *models.User
	result := utils.MySqlDB.Find(&user, "id = ?", id)
	return user, result.Error
}
