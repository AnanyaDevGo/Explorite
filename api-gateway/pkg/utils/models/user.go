package models

type UserSignup struct {
	Email       string `json:"email" binding:"required" validate:"required,email"`
	Password    string `json:"password" binding:"required" validate:"min=6,max=20"`
	Firstname   string `json:"firstname" gorm:"validate:required"`
	Lastname    string `json:"lastname" gorm:"validate:required"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Bio         string `json:"bio"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type UserDetailResponse struct {
	Email       string `json:"email" binding:"required" validate:"required,email"`
	Password    string `json:"password" binding:"required" validate:"min=6,max=20"`
	Firstname   string `json:"firstname" gorm:"validate:required"`
	Lastname    string `json:"lastname" gorm:"validate:required"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
}

type User struct {
	ID           uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname    string `json:"firstname" gorm:"validate:required"`
	Lastname     string `json:"lastname" gorm:"validate:required"`
	Email        string `json:"email" gorm:"validate:required"`
	Password     string `json:"password" gorm:"validate:required"`
	Username     string `json:"username"`
	ProfileImage string `json:"profile_image"`
	PhoneNumber  string `json:"phone_number"`
	DateOfBirth  string `json:"date_of_birth"`
	Gender       string `json:"gender"`
	Bio          string `json:"bio"`
}

type TokenUser struct {
	User  UserDetailResponse
	Token string
}
