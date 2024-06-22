package usecase

import (
	"errors"
	"fmt"
	"postservice/pkg/config"
	"postservice/pkg/helper"
	interfaces "postservice/pkg/repository/interface"
	"postservice/pkg/utils/models"
	"strconv"

	"github.com/google/uuid"
)

type PostUseCase struct {
	postRepository interfaces.PostRepository
}

func NewPostUseCase(postRepository interfaces.PostRepository) *PostUseCase {
	return &PostUseCase{
		postRepository: postRepository,
	}
}

func (uc *PostUseCase) AddPost(post models.AddPost) error {
	fmt.Println("erorrrrrrr usecase")

	if post.UserId <= "" || post.Caption == "" || post.MediaURL == "" {
		return errors.New("invalid input data")
	}

	// img, err := jpeg.Decode(bytes.NewReader(post.Media))
	// if err != nil {
	// 	return err
	// }
	// data, err := helper.ImageToMultipartFile(img, post.Caption)
	// if err != nil {
	// 	return err
	// }
	fileUID := uuid.New()
	fileName := fileUID.String()
	h := helper.NewHelper(config.Config{})

	url, err := h.AddImageToAwsS3(post.Media, fileName)
	if err != nil {
		return err
	}
	userID, err := strconv.Atoi(post.UserId)
	if err != nil {
		return err
	}
	err = uc.postRepository.AddPost(userID, post.Caption, url, post.MediaURL)
	if err != nil {
		return err
	}

	// 	err = uc.postRepository.AddPost(userID, caption, url)
	// 	if err != nil {
	// 		return err
	// 	}

	return nil
	// }
}

func (ps *PostUseCase) ListPost() ([]models.AddPost, error) {
	posts, err := ps.postRepository.ListPost()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
func (ps *PostUseCase) EditPost(userID int, post models.EditPost) error {
	if userID <= 0 {
		return errors.New("invalid user ID")
	}

	exists, err := ps.postRepository.PostExists(post.PostId)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("post does not exist")
	}

	if post.Caption == "" {
		return errors.New("caption cannot be empty")
	}

	err = ps.postRepository.UpdatePostByID(post.PostId, post)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PostUseCase) DeletePost(postID int) error {
	if postID <= 0 {
		return errors.New("invalid post ID")
	}

	exists, err := ps.postRepository.PostExists(strconv.Itoa(postID))
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("post does not exist")
	}

	err = ps.postRepository.DeletePostByID(postID)
	if err != nil {
		return err
	}

	return nil
}
func (ps *PostUseCase) CreateCommentPost(postId, userId int, comment string) (bool, error) {

	if postId <= 0 {
		return false, errors.New("postId is required")
	}
	if comment == "" {
		return false, errors.New("comment is required")
	}

	okP, err := ps.postRepository.PostExists(strconv.Itoa(postId))

	if err != nil {
		return false, err
	}
	if !okP {
		return false, errors.New("post does not exist")
	}
	ok, err := ps.postRepository.CreateCommentPost(postId, userId, comment)
	if err != nil {
		return false, err
	}
	msg := fmt.Sprintf("%s Commented your post %d comment: %s", strconv.Itoa(userId), postId, comment)
	helper.SendNotification(models.Notification{
		UserID: userId,
		PostID: postId,
	}, []byte(msg))
	return ok, nil

}

func (ps *PostUseCase) UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error) {
	if postId <= 0 {
		return false, errors.New("postId is required")
	}
	if comment == "" {
		return false, errors.New("comment is required")
	}
	if commentId <= 0 {
		return false, errors.New("commentId is required")
	}

	okc, err := ps.postRepository.IsCommentIdExist(commentId)
	if err != nil {
		return false, err
	}
	if !okc {
		return false, errors.New("comment does not exist")
	}

	okcu, err := ps.postRepository.IsCommentIdBelongsUserId(commentId, userId)
	if err != nil {
		return false, err
	}
	if !okcu {
		return false, errors.New("comment does not belongs to you")
	}
	ok, err := ps.postRepository.UpdateCommentPost(commentId, postId, userId, comment)
	if err != nil {
		return false, err
	}
	return ok, nil

}

func (ps *PostUseCase) DeleteCommentPost(postId, userId, commentId int) (bool, error) {
	fmt.Println("postidddd", postId)
	if postId <= 0 {
		return false, errors.New("postId is required")
	}
	if commentId <= 0 {
		return false, errors.New("commentId is required")
	}
	okc, err := ps.postRepository.IsCommentIdExist(commentId)
	if err != nil {
		return false, err
	}
	if !okc {
		return false, errors.New("comment does not exist")
	}

	okcu, err := ps.postRepository.IsCommentIdBelongsUserId(commentId, userId)
	if err != nil {
		return false, err
	}
	if !okcu {
		return false, errors.New("comment does not belongs to you")
	}
	ok, err := ps.postRepository.DeleteCommentPost(commentId, postId, userId)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (ps *PostUseCase) UpvotePost(userID, postID int) error {
	if postID <= 0 {
		return errors.New("invalid post ID")
	}

	fmt.Println("posttttttt", postID)
	fmt.Println("userrrrr ", userID)
	exists, err := ps.postRepository.PostExists(strconv.Itoa(postID))
	if err != nil {
		return err
	}
	fmt.Println("exisssst", exists)
	if !exists {
		return errors.New("post does not exist")
	}
	// upvoted, err := ps.postRepository.IsPostvoted(postID, userID)
	// if err != nil {
	// 	return err
	// }
	// if upvoted {
	// 	return errors.New("post already upvoted")
	// }

	err = ps.postRepository.UpvotePost(postID, userID)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("%s liked your postid %d", strconv.Itoa(userID), postID)
	helper.SendNotification(models.Notification{
		UserID: userID,
		PostID: postID,
	}, []byte(msg))

	return nil
}

func (ps *PostUseCase) DownvotePost(userID, postID int) error {
	if postID <= 0 {
		return errors.New("invalid post ID")
	}

	exists, err := ps.postRepository.PostExists(strconv.Itoa(postID))
	if err != nil {
		return err
	}
	fmt.Println("posttt existtt", exists)
	if !exists {
		return errors.New("post does not exist")
	}

	// downvoted, err := ps.postRepository.IsPostvoted(postID, userID)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("voteeeee", downvoted)
	// if !downvoted {
	// 	return errors.New("post already downvoted")
	// }

	err = ps.postRepository.DownvotePost(postID, userID)
	if err != nil {
		return err
	}

	return nil
}
