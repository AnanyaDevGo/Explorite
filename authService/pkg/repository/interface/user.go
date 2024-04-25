package interfaces

import (
	"authservice/pkg/domain"
	"authservice/pkg/utils/models"
)

type UserRepository interface {
	UserSignUp(userDetails models.UserSignup) (models.UserDetailResponse, error)
	FindUserByEmail(user models.UserLogin) (models.UserSignup, error)
	CheckUserExistsByEmail(email string) (*domain.User, error)
}
