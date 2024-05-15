package repository

import (
	"errors"
	//"postservice/pkg/domain"
	"postservice/pkg/domain"
	interfaces "postservice/pkg/repository/interface"
	"postservice/pkg/utils/models"

	"gorm.io/gorm"
)

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository(DB *gorm.DB) interfaces.PostRepository {
	return &postRepository{
		DB: DB,
	}
}
func (pr *postRepository) AddPost(userID int, caption string, url string, mediaurl string) error {
	err := pr.DB.Exec("INSERT INTO posts (user_id, caption, url, media_url) VALUES (?, ?, ?, ?)", userID, caption, url, mediaurl).Error
	if err != nil {
		return errors.New("error on updating data base")
	}
	return nil
}
func (pr *postRepository) ListPost() ([]models.AddPost, error) {
	var posts []models.AddPost

	query := `
        SELECT id,caption, user_id, url as media, media_url
        FROM posts
    `
	if err := pr.DB.Raw(query).Scan(&posts).Error; err != nil {
		return []models.AddPost{}, err
	}

	return posts, nil
}

func (pr *postRepository) UpdatePostByID(postID int, updatedPost models.EditPost) error {
	err := pr.DB.Exec("UPDATE posts SET caption = ? WHERE id = ?", updatedPost.Caption, postID).Error
	if err != nil {
		return err
	}

	return nil
}

func (pr *postRepository) DeletePostByID(postID int) error {
	err := pr.DB.Exec("DELETE FROM posts WHERE id = ?", postID).Error
	if err != nil {
		return err
	}

	return nil
}
func (pr *postRepository) PostExists(postID int) (bool, error) {
	var count int64
	query := "SELECT COUNT(*) FROM posts WHERE id = ?"

	if err := pr.DB.Raw(query, postID).Count(&count).Error; err != nil {
		return false, err
	}

	return count == 1, nil
}

// func (pr *postRepository) SavePost(postID int) error {
// 	saved, err := pr.CheckIfPostSaved(postID)
// 	if err != nil {
// 		return err
// 	}
// 	if saved {
// 		return errors.New("post already saved")
// 	}
// 	_, err = pr.DB.Exec("INSERT INTO saved_posts (user_id, post_id) VALUES (?, ?)", pr.UserID, postID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (pr *postRepository) UnSavePost(postID int) error {
//     saved, err := pr.CheckIfPostSaved(postID)
//     if err != nil {
//         return err
//     }
//     if !saved {
//         return errors.New("post already unsaved")
//     }
//     _, err = pr.DB.Exec("DELETE FROM saved_posts WHERE user_id = ? AND post_id = ?", pr.UserID, postID)
//     if err != nil {
//         return err
//     }
//     return nil
// }

// func (pr *postRepository) CheckIfPostSaved(postID int) (bool, error) {
// 	var count int64
// 	var savedPost domain.SavedPost

// 	if err := pr.DB.Model(&savedPost).Where("user_id = ? AND post_id = ?", pr.UserID, postID).Count(&count).Error; err != nil {
// 		return false, err
// 	}

// 	return count > 0, nil
// }

func (pr *postRepository) UpvotePost(postID int) error {
	query := "UPDATE posts SET votes = votes + 1 WHERE id = ?"

	if err := pr.DB.Exec(query, postID).Error; err != nil {
		return err
	}

	return nil
}
func (pr *postRepository) DownvotePost(postID int) error {
	query := "UPDATE posts SET votes = votes - 1 WHERE id = ?"

	if err := pr.DB.Exec(query, postID).Error; err != nil {
		return err
	}

	return nil
}
func (pr *postRepository) GetPostByID(postID int) (domain.Post, error) {
	var post domain.Post
	if err := pr.DB.First(&post, postID).Error; err != nil {
		return domain.Post{}, err
	}
	return post, nil
}
func (pr *postRepository) IsPostUpvoted(postID int, userID int) (bool, error) {
	var count int64
	query := "SELECT COUNT(*) FROM upvotes WHERE post_id = ? AND user_id = ?"
	if err := pr.DB.Raw(query, postID, userID).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (pr *postRepository) IsPostDownvoted(postID int, userID int) (bool, error) {
	var count int64
	query := "SELECT COUNT(*) FROM downvotes WHERE post_id = ? AND user_id = ?"
	if err := pr.DB.Raw(query, postID, userID).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
