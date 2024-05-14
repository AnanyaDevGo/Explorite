package repository

import (
	"errors"
	interfaces "postservice/pkg/repository/interface"

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
func (pr *postRepository) AddPost(userID int, caption string, url string) error {
	err := pr.DB.Exec("INSERT INTO posts (user_id, caption, url) VALUES (?, ?, ?)", userID, caption, url).Error
	if err != nil {
		return errors.New("error on updating data base")
	}
	return nil
}
// func (pr *postRepository) ListPost() ([]models.AddPost, error) {
//     var posts []models.AddPost

//     query := `
//         SELECT caption, user_id, media_url, media_data
//         FROM posts
//     `
//     if err := pr.DB.Raw(query).Scan(&posts).Error; err != nil {
//         return []models.AddPost{}, err
//     }

//     return posts, nil
// }

// func (pr *postRepository) UpdatePostByID(postID int, updatedPost models.EditPost) error {
// 	err := pr.DB.Exec("UPDATE posts SET caption = ? WHERE id = ?", updatedPost.Caption, postID).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func (pr *postRepository) DeletePostByID(postID int) error {
// 	err := pr.DB.Exec("DELETE FROM posts WHERE id = ?", postID).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
