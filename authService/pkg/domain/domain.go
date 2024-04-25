package domain

import "authservice/pkg/utils/models"

type Admin struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
}
type TokenAdmin struct {
	Admin models.AdminDetailsResponse
	Token string
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
	User  models.UserDetailResponse
	Token string
}

