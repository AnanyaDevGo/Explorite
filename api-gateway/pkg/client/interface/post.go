package interfaces

import (
	"ExploriteGateway/pkg/utils/models"
	"mime/multipart"
)

type PostClient interface {
	AddPost(post models.AddPost, file *multipart.FileHeader) error
	ListPost() ([]models.AddPost, error)
	EditPost(id int, post models.EditPost) error
	DeletePost(id int) error
	// SavePost(postID int) error
	// UnSavePost(postID int) error

	CreateCommentPost(postId, userId int, comment string) (bool, error)
	UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error)
	DeleteCommentPost(postId, userId, commentId int) (bool, error)
	UpvotePost(userID, postID int) error
	DownvotePost(userID, postID int) error
}
