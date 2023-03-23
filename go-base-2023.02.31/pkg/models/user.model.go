package models

type User struct {
	Base
	UID   string
	Email string `gorm:"unique"`
	Name  string
	Role  string
	Todo  []Todo `gorm:"foreignKey:UserID; references:ID"`
}
