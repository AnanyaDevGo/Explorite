package helper

import (
	cfg "ExploriteGateway/pkg/config"
	"fmt"
	"log"
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
	log.Println("Opening file for upload...")
	f, openErr := file.Open()
	if openErr != nil {
		log.Println("Error opening file:", openErr)
		return "", openErr
	}
	defer func() {
		log.Println("Closing file after upload.")
		f.Close()
	}()

	log.Println("Creating new AWS session...")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(h.cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			h.cfg.Access_key_ID,
			h.cfg.Secret_access_key,
			"",
		),
	})
	if err != nil {
		log.Println("Error creating AWS session:", err)
		return "", err
	}

	log.Println("Uploading file to S3...")
	uploader := s3manager.NewUploader(sess)
	bucketName := "crocsclub"

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file.Filename),
		Body:   f,
	})
	if err != nil {
		log.Println("Error uploading file to S3:", err)
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, file.Filename)
	log.Println("File uploaded successfully. URL:", url)
	return url, nil
}
