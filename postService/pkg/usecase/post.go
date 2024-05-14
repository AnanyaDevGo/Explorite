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
	err = uc.postRepository.AddPost(userID, post.Caption, url)
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

// func (ps *PostUseCase) ListPost() ([]models.AddPost, error) {
//     posts, err := ps.postRepository.ListPost()
//     if err != nil {
//         return nil, err
//     }

//     return posts, nil
// }

// func (ps *PostUseCase) EditPost(postId int, post models.EditPost) error {
// 	if post.Caption == "" {
// 		return errors.New("caption cannot be empty")
// 	}
// 	if postId <= 0 {
// 		return errors.New("invalid post ID")
// 	}

// 	err := ps.postRepository.UpdatePostByID(postId, post)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (ps *PostUseCase) DeletePost(postID int) error {
// 	if postID <= 0 {
// 		return errors.New("invalid post ID")
// 	}

// 	err := ps.postRepository.DeletePostByID(postID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
