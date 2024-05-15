package domain

type Post struct {
	ID       uint   `json:"id" gorm:"uniquekey; not null"`
	UserID   uint   `json:"user_id"`
	Url      string `json:"url"`
	Caption  string `json:"caption"`
	MediaUrl string `json:"media_url"`
	Votes    int    `json:"votes"`
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
