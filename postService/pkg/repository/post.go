package repository

import (
	"errors"
	"fmt"
	"time"

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

func (pr *postRepository) UpdatePostByID(postID string, updatedPost models.EditPost) error {
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
func (pr *postRepository) PostExists(postID string) (bool, error) {
	fmt.Println("postid", postID)

	var count int
	query := "SELECT COUNT(*) FROM posts WHERE id = ?"

	if err := pr.DB.Raw(query, postID).Scan(&count).Error; err != nil {
		fmt.Println("errorrrrr", err)
		return false, err
	}
	fmt.Println("count", count)

	return count > 0, nil
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

func (pr *postRepository) CreateCommentPost(postId, userId int, comment string) (bool, error) {
	querry := `insert into comments (post_id,comment,user_id,created_at)
	values ($1,$2,$3,$4)`
	err := pr.DB.Exec(querry, postId, comment, userId, time.Now()).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (pr *postRepository) IsCommentIdExist(commentId int) (bool, error) {
	var ok int
	querry := `select count(*) from comments where id = ?`
	err := pr.DB.Raw(querry, commentId).Scan(&ok).Error
	if err != nil {
		return false, err
	}
	return ok > 0, nil
}

func (pr *postRepository) IsCommentIdBelongsUserId(commentId, userId int) (bool, error) {
	var ok int
	querry := `select count(*) from comments where id = ? and  user_id = ?`
	err := pr.DB.Raw(querry, commentId, userId).Scan(&ok).Error
	if err != nil {
		return false, err
	}
	return ok > 0, nil
}

func (pr *postRepository) UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error) {
	querry := `update comments set comment = ?, updated_at = ? where id = ? and  user_id = ? and post_id = ? `
	err := pr.DB.Exec(querry, comment, time.Now(), commentId, userId, postId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (pr *postRepository) DeleteCommentPost(postId, userId, commentId int) (bool, error) {
	querry := `delete from comments where id = ? and  user_id = ? and post_id = ? `
	err := pr.DB.Exec(querry, commentId, userId, postId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (pr *postRepository) UpvotePost(postID, userID int) error {
	query := "insert into post_votes (user_id, post_id, vote)values(?,?,?)"

	if err := pr.DB.Exec(query, postID, userID, true).Error; err != nil {
		return err
	}

	return nil
}
func (pr *postRepository) DownvotePost(postID, userID int) error {
	query := "UPDATE post_votes SET vote = ? WHERE user_id = ? AND post_id = ?"

	if err := pr.DB.Exec(query, false, userID, postID).Error; err != nil {
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

// func (pr *postRepository) IsPostUpvoted(postID int, userID int) (bool, error) {
// 	var count int64
// 	query := "SELECT COUNT(*) FROM post_votes WHERE post_id = ? AND user_id = ?"
// 	if err := pr.DB.Raw(query, postID, userID).Scan(&count).Error; err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }

func (pr *postRepository) IsPostvoted(postID int, userID int) (bool, error) {
	var vote bool
	query := "SELECT vote FROM post_votes WHERE post_id = ? AND user_id = ?"
	err := pr.DB.Raw(query, postID, userID).Scan(&vote).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return vote, nil
}
func (pr *postRepository) GetPostedUserID(id int) (int, error) {

	var idd int
	err := pr.DB.Raw("select user_id from posts where id = ?", id).Scan(&idd).Error
	if err != nil {
		return 0, err
	}
	return idd, nil
}
