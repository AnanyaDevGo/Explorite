package interfaces

import "ExploriteGateway/pkg/utils/models"

type UserClient interface {
	UserSignUp(userDetails models.UserSignup) (models.TokenUser, error)
	UserLogin(userDetails models.UserLogin) (models.TokenUser, error)
	UserOTPLogin(email string) (string, error)
	OtpVerification(email, otp string) (bool, error)
	AddProfile(id int, profile models.UserProfile) error
	GetProfile(id int) (models.UserProfile, error)
	EditProfile(id int, profile models.EditProfile) error
	ChangePassword(userID int, oldPassword, newPassword, rePassword string) error
}
