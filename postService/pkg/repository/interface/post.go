package interfaces

import (
	"postservice/pkg/domain"
	"postservice/pkg/utils/models"
)

type PostRepository interface {
	AddPost(userID int, caption string, url string, mediaurl string) error
	ListPost() ([]models.AddPost, error)
	UpdatePostByID(postID int, updatedPost models.EditPost) error
	DeletePostByID(postID int) error
	PostExists(postID int) (bool, error)
	GetPostByID(postID int) (domain.Post, error)
	// CheckIfPostSaved(postID int) (bool, error)
	// SavePost(postID int) error
	// UnSavePost(postID int) error
	UpvotePost(postID, userID int) error
	DownvotePost(postID, userID int) error
	IsPostvoted(postID int, userID int) (bool, error)
}
