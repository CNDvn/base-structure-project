package helpers

import (
	"context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func InitialFirebase(credentialsJSON []byte) (*firebase.App, error) {
	opt := option.WithCredentialsJSON(credentialsJSON)
	return firebase.NewApp(context.Background(), nil, opt)
}
