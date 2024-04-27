package interfaces

import (
	"authservice/pkg/domain"
	"authservice/pkg/utils/models"
)

type AdminUseCase interface {
	AdminSignUp(admindeatils models.AdminSignUp) (*domain.TokenAdmin, error)
	LoginHandler(adminDetails models.AdminLogin) (*domain.TokenAdmin, error)
	GetUsers(page int) ([]models.UserDetailsAtAdmin, error)
	BlockUser(id string) error
	UnBlockUser(id string) error
}
