package repository

import (
	"authservice/pkg/domain"
	interfaces "authservice/pkg/repository/interface"
	"authservice/pkg/utils/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (ur *userRepository) UserSignUp(userDetails models.UserSignup) (models.UserDetailResponse, error) {
	var model models.UserDetailResponse

	fmt.Println("email", model.Email)

	fmt.Println("models", model)
	if err := ur.DB.Raw("INSERT INTO users (firstname, lastname, email, password, phone_number, date_of_birth, gender) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id, firstname, lastname, email, phone_number, date_of_birth, gender", userDetails.Firstname, userDetails.Lastname, userDetails.Email, userDetails.Password, userDetails.PhoneNumber, userDetails.DateOfBirth, userDetails.Gender).Scan(&model).Error; err != nil {
		return models.UserDetailResponse{}, err
	}
	fmt.Println("inside", model.Email)
	return model, nil
}

func (ur *userRepository) CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := ur.DB.Where(&domain.User{Email: email}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &domain.User{}, res.Error
	}
	return &user, nil
}

func (ur *userRepository) FindUserByEmail(user models.UserLogin) (models.UserSignup, error) {
	var userDetail models.UserSignup
	err := ur.DB.Raw("SELECT * FROM users WHERE email=? ", user.Email).Scan(&userDetail).Error
	if err != nil {
		return models.UserSignup{}, errors.New("error checking user details")
	}
	return userDetail, nil
}
