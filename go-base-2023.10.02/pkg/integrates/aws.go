package integrates

import (
	"context"
	"gobase/pkg/helpers"
	"io"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type tAWS struct {
	s3Client *s3.Client
}

var Aws tAWS

func InitAWS() error {
	var err error
	Aws.s3Client = s3.New(s3.Options{
		Region:      helpers.GetENV().AWS_REGION,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(helpers.GetENV().AWS_ACCESS_KEY_ID, helpers.GetENV().AWS_SECRET_ACCESS_KEY, "")),
	})
	return err
}

func (t *tAWS) GetS3() *s3.Client {
	return t.s3Client
}

func (t *tAWS) UploadFileToBucket(file multipart.FileHeader, key string) (*manager.UploadOutput, error) {
	fileContent, err := file.Open()
	if err != nil {
		return nil, err
	}

	body := io.Reader(fileContent)
	uploader := manager.NewUploader(t.s3Client)
	params := &s3.PutObjectInput{
		Bucket: aws.String(helpers.GetENV().AWS_BUCKET_NAME),
		Key:    aws.String(key),
		Body:   body,
		ACL:    "public-read",
	}
	resultUpload, err := uploader.Upload(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	return resultUpload, nil
}
