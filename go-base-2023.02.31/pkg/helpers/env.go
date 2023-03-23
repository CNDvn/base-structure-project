package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	MODE string
	PORT string

	MYSQL_HOST     string
	MYSQL_PORT     string
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_DB_NAME  string

	FIREBASE_SERVICE_ACCOUNT string
}

var env ENV

func LoadENV() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	env = ENV{
		MODE: os.Getenv("MODE"),
		PORT: os.Getenv("APP_PORT"),

		MYSQL_HOST:     os.Getenv("MYSQL_HOST"),
		MYSQL_PORT:     os.Getenv("MYSQL_PORT"),
		MYSQL_USER:     os.Getenv("MYSQL_USER"),
		MYSQL_PASSWORD: os.Getenv("MYSQL_PASSWORD"),
		MYSQL_DB_NAME:  os.Getenv("MYSQL_DB_NAME"),

		FIREBASE_SERVICE_ACCOUNT: os.Getenv("FIREBASE_SERVICE_ACCOUNT"),
	}
	return nil
}

func GetENV() ENV {
	return env
}
