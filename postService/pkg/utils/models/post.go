package models

type AddPost struct {
	Caption string `form:"caption" validate:"lte=60"`
	Media   []byte `form:"media" validate:"required"`
	//Votes    uint   `json:"votes"`
	UserId   string `validate:"required"`
	MediaURL string `form:"mediaurl" validate:"required"`
	Votes    int    `json:"votes"`
}

type LikeRequest struct {
	PostID string `json:"postid"`
	UserID string
}

type EditPost struct {
	Caption string `form:"caption" validate:"lte=60"`
	UserId  string `validate:"required"`
	PostId  string `json:"postid" validate:"required,number"`
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
