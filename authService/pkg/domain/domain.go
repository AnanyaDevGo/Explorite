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
	ID          uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname   string `json:"firstname" gorm:"validate:required"`
	Lastname    string `json:"lastname" gorm:"validate:required"`
	Email       string `json:"email" gorm:"validate:required"`
	Password    string `json:"password" gorm:"validate:required"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Bio         string `json:"bio"`
}

type TokenUser struct {
	User  models.UserDetailResponse
	Token string
}

type Users struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8,max=20"`
	Phone    string `json:"phone"`
	Blocked  bool   `json:"blocked" gorm:"default:false"`
	Isadmin  bool   `json:"isadmin" gorm:"default:false;check:isadmin IN (true, false)"`
}
type UserProfile struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	UserId   int    `json:"user_id"`
	Users    Users  `json:"-" gorm:"foreignkey:UserId"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email" validate:"email"`
	Website  string `json:"website" validate:"website"`
	Location string `json:"location"`
	Phone    string `json:"phone"`
	Bio      string `json:"bio"`
}
