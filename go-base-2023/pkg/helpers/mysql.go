package helpers

import (
	"fmt"

	"basego/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.MYSQL_USER,
		env.MYSQL_PASSWORD,
		env.MYSQL_HOST,
		env.MYSQL_PORT,
		env.MYSQL_DB_NAME)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func MysqlAutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})
}
