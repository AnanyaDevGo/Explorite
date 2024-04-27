package interfaces

import "ExploriteGateway/pkg/utils/models"

type AdminClient interface {
	AdminSignUp(admindeatils models.AdminSignUp) (models.TokenAdmin, error)
	AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error)
	GetUsers(page int) ([]models.UserDetailsAtAdmin, error)
	BlockUser(id string) error
	UnBlockUser(id string) error
}
