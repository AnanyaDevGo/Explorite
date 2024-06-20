package domain

import "time"

type Post struct {
	ID       uint   `json:"id" gorm:"uniquekey; not null"`
	UserID   uint   `json:"user_id"`
	Url      string `json:"url"`
	Caption  string `json:"caption"`
	MediaUrl string `json:"media_url"`
	//Votes    int    `json:"votes"`
}
type SavedPost struct {
	UserID int
	PostID int
}
type Upvote struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}

type Downvote struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}
type PostVote struct {
	ID     uint `json:"id" gorm:"uniquekey; not null"`
	UserID int  `json:"user_id"`
	PostID int  `json:"post_id"`
	Vote   bool `json:"vote" gorm:"default:false"`
}
type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	PostID    uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	Comment   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
