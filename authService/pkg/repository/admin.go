package repository

import (
	"authservice/pkg/domain"
	interfaces "authservice/pkg/repository/interface"
	"authservice/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"

	//"strconv"

	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (ad *adminRepository) AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error) {
	var model models.AdminDetailsResponse

	fmt.Println("email", model.Email)

	fmt.Println("models", model)
	if err := ad.DB.Raw("INSERT INTO admins (firstname, lastname, email, password) VALUES (?, ?, ?, ?) RETURNING id, firstname, lastname, email", adminDetails.Firstname, adminDetails.Lastname, adminDetails.Email, adminDetails.Password).Scan(&model).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}
	fmt.Println("inside", model.Email)
	return model, nil
}
func (ad *adminRepository) CheckAdminExistsByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	res := ad.DB.Where(&domain.Admin{Email: email}).First(&admin)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.Admin{}, res.Error
	}
	return &admin, nil
}

func (ad *adminRepository) FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error) {
	var user models.AdminSignUp
	err := ad.DB.Raw("SELECT * FROM admins WHERE email=? ", admin.Email).Scan(&user).Error
	if err != nil {
		return models.AdminSignUp{}, errors.New("error checking user details")
	}
	return user, nil
}
func (ad *adminRepository) GetUsers(page int) ([]models.UserDetailsAtAdmin, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 20
	var userDetails []models.UserDetailsAtAdmin

	query := `
        SELECT id, CONCAT(firstname, ' ', lastname) AS name, email, phone_number as phone, blocked as block_status
        FROM users
        LIMIT ? OFFSET ?
    `
	if err := ad.DB.Raw(query, 20, offset).Scan(&userDetails).Error; err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil
}

func (ad *adminRepository) GetUserByID(id string) (domain.User, error) {

	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.User{}, err
	}

	var count int
	if err := ad.DB.Raw("select count(*) from users where id = ?", user_id).Scan(&count).Error; err != nil {
		return domain.User{}, err
	}
	if count < 1 {
		return domain.User{}, errors.New("user for the given id does not exist")
	}

	query := fmt.Sprintf("select * from users where id = '%d'", user_id)
	var userDetails domain.User

	if err := ad.DB.Raw(query).Scan(&userDetails).Error; err != nil {
		return domain.User{}, err
	}

	return userDetails, nil
}

func (ad *adminRepository) UpdateBlockUserByID(user domain.User) error {

	err := ad.DB.Exec("update users set blocked = ? where id = ?", user.Blocked, user.ID).Error
	if err != nil {
		return err
	}

	return nil

}
