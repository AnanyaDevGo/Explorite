package models

type AddPost struct {
	Caption string `form:"caption" validate:"lte=60"`
	Media   []byte `form:"media" validate:"required"`

	UserId   string `validate:"required"`
	MediaURL string `form:"mediaurl" validate:"required"`
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
