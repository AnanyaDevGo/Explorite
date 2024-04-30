package domain

import (
	"authservice/pkg/utils/models"
	"time"
)

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
	ID          uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname   string `json:"firstname" gorm:"validate:required"`
	Lastname    string `json:"lastname" gorm:"validate:required"`
	Email       string `json:"email" gorm:"validate:required"`
	Password    string `json:"password" gorm:"validate:required"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Bio         string `json:"bio"`
	Blocked     bool   `json:"blocked" gorm:"default:false"`
	Isadmin     bool   `json:"isadmin" gorm:"default:false;check:isadmin IN (true, false)"`
}

type TokenUser struct {
	User  models.UserDetailResponse
	Token string
}
type UserProfile struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	UserId   int    `json:"user_id"`
	Users    User   `json:"-" gorm:"foreignkey:UserId"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email" validate:"email"`
	Website  string `json:"website" validate:"website"`
	Location string `json:"location"`
	Phone    string `json:"phone"`
	Bio      string `json:"bio"`
}
type UserOTPLogin struct {
	Email      string    `json:"email" validate:"email"`
	Expiration time.Time `json:"expiration"`
	Otp        int       `json:"otp"`
}
