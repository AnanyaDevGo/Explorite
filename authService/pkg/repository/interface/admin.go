package interfaces

import (
	"authservice/pkg/domain"
	"authservice/pkg/utils/models"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)
	GetUsers(page int) ([]models.UserDetailsAtAdmin, error)
	GetUserByID(id string) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
}
