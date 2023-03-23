package userdto

import "basego/pkg/models"

type UserInfo struct {
	ID    uint
	UID   string
	Email string
	Name  string
	Role  string
}

func (userDto *UserInfo) MapFrom(user models.User) {
	userDto.ID = user.ID
	userDto.UID = user.UID
	userDto.Email = user.Email
	userDto.Name = user.Name
	userDto.Role = user.Role
}
