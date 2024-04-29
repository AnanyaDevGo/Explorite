package interfaces

import (
	"authservice/pkg/domain"
	"authservice/pkg/utils/models"
)

type UserRepository interface {
	UserSignUp(userDetails models.UserSignup) (models.UserDetailResponse, error)
	FindUserByEmail(user models.UserLogin) (models.UserSignup, error)
	CheckUserExistsByEmail(email string) (*domain.User, error)
	AddProfile(id int, profile models.UserProfile) error
	GetProfile(id int) ([]domain.UserProfile, error)
	ValidatePhoneNumber(phone string) bool
	ValidateAlphabets(data string) (bool, error)
}
