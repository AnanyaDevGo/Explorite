package domain

type Post struct {
	ID            uint      `json:"id" gorm:"uniquekey; not null"`
	UserID        uint      `json:"user_id"`
	Url           string    `json:"url"`
	Caption       string    `json:"caption"`
}
