package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	MODE     string
	API_PORT string

	JWT_SECRET_KEY           string
	REFRESH_TOKEN_SECRET_KEY string

	MONGO_URI string

	AWS_ACCESS_KEY_ID     string
	AWS_SECRET_ACCESS_KEY string

	AWS_REGION            string
	AWS_BUCKET_NAME       string
	MONGO_INITDB_DATABASE string
}

var env *ENV

func LoadENV() error {

	if os.Getenv("API_PORT") == "" {
		if err := godotenv.Load(); err != nil {
			if err := godotenv.Load("../../.env"); err != nil {
				return err
			}
		}
	}

	env = &ENV{
		MODE:     os.Getenv("MODE"),
		API_PORT: os.Getenv("API_PORT"),

		JWT_SECRET_KEY:           os.Getenv("JWT_SECRET_KEY"),
		REFRESH_TOKEN_SECRET_KEY: os.Getenv("REFRESH_TOKEN_SECRET_KEY"),

		MONGO_URI: os.Getenv("MONGO_URI"),

		AWS_ACCESS_KEY_ID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWS_SECRET_ACCESS_KEY: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AWS_REGION:            os.Getenv("AWS_REGION"),
		AWS_BUCKET_NAME:       os.Getenv("AWS_BUCKET_NAME"),
		MONGO_INITDB_DATABASE: os.Getenv("MONGO_INITDB_DATABASE"),
	}
	return nil
}

func GetENV() *ENV {
	if env == nil {
		LoadENV()
	}
	return env
}
