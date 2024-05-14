package interfaces

import (
	"ExploriteGateway/pkg/utils/models"
	"mime/multipart"
)

type PostClient interface {
	AddPost(post models.AddPost, file *multipart.FileHeader) error
	//ListPost() ([]models.AddPost, error)
	//EditPost(id int, post models.EditPost) error
	//DeletePost(id int) error
}