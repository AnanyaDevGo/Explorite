package helper

import (
	"bytes"
	"fmt"
	cfg "postservice/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Helper struct {
	cfg cfg.Config
}

func NewHelper(cfg cfg.Config) *Helper {
	return &Helper{cfg: cfg}
}

func (h *Helper) AddImageToAwsS3(file []byte, filename string) (string, error) {

	// f, openErr := file.Open()

	// if openErr != nil {
	// 	fmt.Println("erorrrrrrr 1 helper", openErr)
	// 	return "", openErr
	// }

	// defer f.Close()
	config, err := cfg.LoadConfig()
	if err != nil {
		return "", err
	}

	fmt.Println("pppppppp", config.DBHost)

	fmt.Println("print1", config.AWSRegion)
	fmt.Println("print2", config.Access_key_ID)
	fmt.Println("print3", config.Secret_access_key)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			config.Access_key_ID,
			config.Secret_access_key,
			"",
		),
	})
	if err != nil {
		fmt.Println("erorrrr here", err)
		return "", err
	}
	uploader := s3manager.NewUploader(sess)
	bucketName := "crocsclub"

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("test1"),
		Body:   bytes.NewReader(file),
	})

	expo := "explorite"
	if err != nil {
		fmt.Println("erroorrrr 2", err)
		return "", err
	}
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", expo, filename)
	return url, nil
}

// func ImageToMultipartFile(imageData image.Image, filename string) (*multipart.FileHeader, error) {
// 	fmt.Println("here image to multi")
// 	tempFile, err := ioutil.TempFile("", "image_*.jpg")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer os.Remove(tempFile.Name())

// 	err = jpeg.Encode(tempFile, imageData, nil)
// 	if err != nil {
// 		fmt.Println("erorrrrr 3", err)
// 		return nil, err
// 	}

// 	err = tempFile.Close()
// 	if err != nil {
// 		fmt.Println("erorrrrr 4", err)
// 		return nil, err
// 	}

// 	fileInfo, err := os.Stat(tempFile.Name())
// 	if err != nil {
// 		fmt.Println("erorrrrr 5", err)
// 		return nil, err
// 	}

// 	fileHeader := &multipart.FileHeader{
// 		Filename: fileInfo.Name(),
// 		Size:     fileInfo.Size(),
// 	}

// 	return fileHeader, nil
// }
