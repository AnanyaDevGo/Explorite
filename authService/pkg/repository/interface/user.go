package interfaces

import (
	"authservice/pkg/domain"
	"authservice/pkg/utils/models"
	"time"
)

type UserRepository interface {
	UserSignUp(userDetails models.UserSignup) (models.UserDetailResponse, error)
	FindUserByEmail(user models.UserLogin) (models.UserSignup, error)
	CheckUserExistsByEmail(email string) (*domain.User, error)
	DeleteRecentOtpRequestsBefore5min() error
	TemporarySavingUserOtp(otp int, userEmail string, expiration time.Time) error
	VerifyOTP(email, otp string) (bool, error)
	GetUserName(email string) (string, error)

	AddProfile(id int, profile models.UserProfile) error
	GetProfile(id int) (models.UserProfile, error)
	EditProfile(id int, user models.EditProfile) (models.EditProfile, error)
	ValidatePhoneNumber(phone string) bool
	ValidateAlphabets(data string) (bool, error)
	IsValidEmail(email string) bool
	IsValidWebsite(website string) bool
	ChangePassword(id int, password string) error
	GetPassword(id int) (string, error)
}
