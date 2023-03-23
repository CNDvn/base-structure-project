package services

import (
	"context"
	"errors"
	"gobase/pkg/databases"
	"gobase/pkg/errormsg"
	"gobase/pkg/schemas"
	"gobase/pkg/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TUser struct{}

func (t *TUser) Create(user *schemas.TUser) *utils.CustomError {
	mongoClient := databases.MongoClient{}
	mongoClient.Connection()
	defer mongoClient.Close()

	utils.SetDefaultInsert(user)
	insertResult, err := schemas.GetUserCollection(mongoClient.Client).InsertOne(context.Background(), user)
	if err != nil {
		utils.PrintLog("func (t *TUser) Create(user *schemas.TUser)", err.Error())
		return &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CANNOT_CREATE_USER,
		}
	}
	if insertResult == nil {
		utils.PrintLog("func (t *TUser) Create(user *schemas.TUser)", "insertResult nil")
		return &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CANNOT_CREATE_USER,
		}
	}
	user.ID = insertResult.InsertedID.(primitive.ObjectID)
	return nil
}

func (t *TUser) FindOne(filter interface{}, opts ...*options.FindOneOptions) (*schemas.TUser, error) {
	mongoClient := databases.MongoClient{}
	mongoClient.Connection()
	defer mongoClient.Close()

	singleResult := schemas.GetUserCollection(mongoClient.Client).FindOne(context.Background(), filter, opts...)

	if singleResult == nil {
		return nil, errors.New(errormsg.NOT_FOUND_USER)
	}
	if singleResult.Err() != nil {
		utils.PrintLog("func (t *TUser) FindOne", singleResult.Err().Error())
		return nil, errors.New(errormsg.NOT_FOUND_USER)
	}

	var user *schemas.TUser
	if err := singleResult.Decode(&user); err != nil {
		utils.PrintLog("func (t *TUser) FindOne", err.Error())
		return nil, errors.New(errormsg.NOT_FOUND_USER)
	}

	return user, nil
}
