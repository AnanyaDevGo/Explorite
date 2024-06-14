package interfaces

import (
	"postservice/pkg/utils/models"
)

type PostUsecase interface {
	AddPost(post models.AddPost) error
	ListPost() ([]models.AddPost, error)
	EditPost(postId int, post models.EditPost) error
	DeletePost(postID int) error
	//SavePost(postID int) error
	//UnSavePost(postID int) error

	CreateCommentPost(postId, userId int, comment string) (bool, error)
	UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error)
	DeleteCommentPost(postId, userId, commentId int) (bool, error)
	UpvotePost(userID, postID int) error
	DownvotePost(userID, postID int) error
}
