package helper

import (
	cfg "ExploriteGateway/pkg/config"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type helper struct {
	cfg cfg.Config
}

func (h *helper) AddImageToAwsS3(file *multipart.FileHeader) (string, error) {

	f, openErr := file.Open()

	if openErr != nil {
		return "", openErr
	}

	defer f.Close()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(h.cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			h.cfg.Access_key_ID,
			h.cfg.Secret_access_key,
			"",
		),
	})
	if err != nil {
		return "", err
	}
	uploader := s3manager.NewUploader(sess)
	bucketName := "crocsclub"

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file.Filename),
		Body:   f,
	})

	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, file.Filename)
	return url, nil
}
