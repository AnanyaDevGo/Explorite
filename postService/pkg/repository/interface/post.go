package interfaces

import (
	"postservice/pkg/domain"
	"postservice/pkg/utils/models"
)

type PostRepository interface {
	AddPost(userID int, caption string, url string, mediaurl string) error
	ListPost() ([]models.AddPost, error)
	UpdatePostByID(postID string, updatedPost models.EditPost) error
	DeletePostByID(postID int) error
	PostExists(postID string) (bool, error)
	GetPostByID(postID int) (domain.Post, error)
	// CheckIfPostSaved(postID int) (bool, error)
	// SavePost(postID int) error
	// UnSavePost(postID int) error
	CreateCommentPost(postId, userId int, comment string) (bool, error)
	IsCommentIdExist(commentId int) (bool, error)
	IsCommentIdBelongsUserId(commentId, userId int) (bool, error)
	UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error)
	DeleteCommentPost(postId, userId, commentId int) (bool, error)
	UpvotePost(postID, userID int) error
	DownvotePost(postID, userID int) error
	IsPostvoted(postID int, userID int) (bool, error)
	GetPostedUserID(id int) (int, error)
}
