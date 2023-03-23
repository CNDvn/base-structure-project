package services

import (
	"context"
	"fmt"
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

type TVideo struct{}

func (t *TVideo) Create(dto *reqdto.TCreateVideoReqDto, userReq *schemas.TUser) (*resdto.TCreateVideoResDto, *utils.CustomError) {
	mongoClient := databases.MongoClient{}
	mongoClient.Connection()
	defer mongoClient.Close()

	nameFile := uuid.New()
	ex := strings.Split(dto.FileName, ".")
	key := utils.PATH_BASE_VIDEO_ORIGIN + nameFile.String() + "." + ex[len(ex)-1]

	result, errUpfile := integrates.Aws.UploadFileToBucket(dto.File, key)
	if errUpfile != nil {
		utils.PrintLog("func (t *TVideo) Create(dto *reqdto.TCreateVideo, userReq *schemas.TUser)", errUpfile.Error())
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CANNOT_UPLOAD_FILE,
		}
	}
	fmt.Println(result)
	// url := strings.Split(preSignURL, "?")[0]
	newVideo := schemas.TVideo{
		Title:       dto.Title,
		Description: dto.Description,
		Tags:        dto.Tags,
		Views:       0,
		URL:         result.Location,
		Owner:       userReq.ID,
	}
	utils.SetDefaultInsert(&newVideo)

	insertResult, err := schemas.GetVideoCollection(mongoClient.Client).InsertOne(context.Background(), &newVideo)
	if err != nil {
		utils.PrintLog("func (t *TVideo) Create(dto *reqdto.TCreateVideo, userReq *schemas.TUser)", err.Error())
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CANNOT_CREATE_VIDEO,
		}
	}
	if insertResult == nil {
		utils.PrintLog("func (t *TVideo) Create(dto *reqdto.TCreateVideo, userReq *schemas.TUser)", "insertResult nil")
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CANNOT_CREATE_VIDEO,
		}
	}

	newVideo.ID = insertResult.InsertedID.(primitive.ObjectID)
	returnObj := &resdto.TCreateVideoResDto{
		TVideo: newVideo,
		// PreSignUrl: preSignURL,
	}
	return returnObj, nil
}

func (t *TVideo) Get(userReq *schemas.TUser, pagination *utils.TPagination) ([]schemas.TVideo, *utils.CustomError) {
	mongoClient := databases.MongoClient{}
	mongoClient.Connection()
	defer mongoClient.Close()
	ctx := context.Background()

	opts := options.Find().SetSort(bson.D{{"created_at", -1}}).SetLimit(pagination.GetLimit()).SetSkip(pagination.GetSkip())
	cursor, err := schemas.GetVideoCollection(mongoClient.Client).Find(ctx, bson.M{
		"owner": userReq.ID,
	}, opts)
	if err != nil {
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.GET_LIST_VIDEO_FAIL,
		}
	}
	defer cursor.Close(ctx)

	var listVideo []schemas.TVideo
	for cursor.Next(ctx) {
		var video *schemas.TVideo
		if err = cursor.Decode(&video); err != nil {
			utils.PrintLog("func (t *TVideo) Get(userReq *schemas.TUser)", err.Error())
		}
		listVideo = append(listVideo, *video)
	}
	return listVideo, nil
}

func (t *TVideo) GetById(id string, userReq *schemas.TUser) (*schemas.TVideo, *utils.CustomError) {
	mongoClient := databases.MongoClient{}
	mongoClient.Connection()
	defer mongoClient.Close()
	ctx := context.Background()

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.PrintLog("func (t *TVideo) GetById(id string, userReq *schemas.TUser)", err.Error())
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.GET_VIDEO_FAIL,
		}
	}

	singleResult := schemas.GetVideoCollection(mongoClient.Client).FindOne(ctx, bson.M{
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

	var video *schemas.TVideo
	if err := singleResult.Decode(&video); err != nil {
		utils.PrintLog("func (t *TUser) FindOne", err.Error())
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.GET_VIDEO_FAIL,
		}
	}

	return video, nil

}
