package interfaces

import (
	"authservice/pkg/domain"
	"authservice/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(userdeatils models.UserSignup) (*domain.TokenUser, error)
	LoginHandler(userDetails models.UserLogin) (*domain.TokenUser, error)
	AddProfile(id int, profile models.UserProfile) error
	GetProfile(id int) ([]domain.UserProfile, error)
}
