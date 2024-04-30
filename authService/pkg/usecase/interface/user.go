package interfaces

import (
	"authservice/pkg/domain"
	"authservice/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(userdeatils models.UserSignup) (*domain.TokenUser, error)
	LoginHandler(userDetails models.UserLogin) (*domain.TokenUser, error)
	UserOTPLogin(email string) (string, error)
	OtpVerification(email, otp string) (bool, error)

	AddProfile(id int, profile models.UserProfile) error
	GetProfile(id int) (models.UserProfile, error)
	EditProfile(id int, user models.EditProfile) (models.EditProfile, error)
}
