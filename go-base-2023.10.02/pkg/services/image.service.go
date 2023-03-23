package services

import (
	"context"
	"gobase/pkg/databases"
	"gobase/pkg/errormsg"
	"gobase/pkg/integrates"
	"gobase/pkg/reqdto"
	"gobase/pkg/resdto"
	"gobase/pkg/schemas"
	"gobase/pkg/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TImage struct{}

func (t *TImage) Create(dto *reqdto.TCreateImageReqDto, userReq *schemas.TUser) (*resdto.TCreateImageResDto, *utils.CustomError) {
	mongoClient := databases.MongoClient{}
	mongoClient.Connection()
	defer mongoClient.Close()

	nameFile := uuid.New()
	ex := strings.Split(dto.FileName, ".")
	key := utils.PATH_BASE_IMAGE_ORIGIN + nameFile.String() + "." + ex[len(ex)-1]

	result, errUpfile := integrates.Aws.UploadFileToBucket(dto.File, key)
	if errUpfile != nil {
		utils.PrintLog("func (t *TImage) Create(dto *reqdto.TCreateImageReqDto, userReq *schemas.TUser)", errUpfile.Error())
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CANNOT_UPLOAD_FILE,
		}
	}

	newImage := schemas.TImage{
		Title:       dto.Title,
		Description: dto.Description,
		Tags:        dto.Tags,
		Views:       0,
		URL:         result.Location,
		Owner:       userReq.ID,
	}
	utils.SetDefaultInsert(&newImage)

	insertResult, err := schemas.GetImageCollection(mongoClient.Client).InsertOne(context.Background(), &newImage)
	if err != nil {
		utils.PrintLog("func (t *TImage) Create(dto *reqdto.TCreateImageReqDto, userReq *schemas.TUser)", err.Error())
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CANNOT_CREATE_IMAGE,
		}
	}
	if insertResult == nil {
		utils.PrintLog("func (t *TImage) Create(dto *reqdto.TCreateImageReqDto, userReq *schemas.TUser)", "insertResult nil")
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CANNOT_CREATE_IMAGE,
		}
	}

	newImage.ID = insertResult.InsertedID.(primitive.ObjectID)
	returnObj := &resdto.TCreateImageResDto{
		TImage: newImage,
	}
	return returnObj, nil
}

func (t *TImage) Get(userReq *schemas.TUser, pagination *utils.TPagination) ([]schemas.TImage, *utils.CustomError) {
	mongoClient := databases.MongoClient{}
	mongoClient.Connection()
	defer mongoClient.Close()
	ctx := context.Background()

	opts := options.Find().SetSort(bson.D{{"created_at", -1}}).SetLimit(pagination.GetLimit()).SetSkip(pagination.GetSkip())
	cursor, err := schemas.GetImageCollection(mongoClient.Client).Find(ctx, bson.M{
		"owner": userReq.ID,
	}, opts)
	if err != nil {
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.GET_LIST_IMAGE_FAIL,
		}
	}
	defer cursor.Close(ctx)

	var listImage []schemas.TImage
	for cursor.Next(ctx) {
		var image *schemas.TImage
		if err = cursor.Decode(&image); err != nil {
			utils.PrintLog("func (t *TImage) Get(userReq *schemas.TUser)", err.Error())
		}
		listImage = append(listImage, *image)
	}
	return listImage, nil
}

func (t *TImage) GetById(id string, userReq *schemas.TUser) (*schemas.TImage, *utils.CustomError) {
	mongoClient := databases.MongoClient{}
	mongoClient.Connection()
	defer mongoClient.Close()
	ctx := context.Background()

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.PrintLog("func (t *TImage) GetById(id string, userReq *schemas.TUser)", err.Error())
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.GET_VIDEO_FAIL,
		}
	}

	singleResult := schemas.GetImageCollection(mongoClient.Client).FindOne(ctx, bson.M{
		"owner": userReq.ID,
		"_id":   idObj,
	})

	if singleResult == nil {
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.GET_VIDEO_FAIL,
		}
	}
	if singleResult.Err() != nil {
		utils.PrintLog("func (t *TUser) FindOne", singleResult.Err().Error())
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.GET_VIDEO_FAIL,
		}
	}

	var image *schemas.TImage
	if err := singleResult.Decode(&image); err != nil {
		utils.PrintLog("func (t *TUser) FindOne", err.Error())
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.GET_VIDEO_FAIL,
		}
	}

	return image, nil

}
