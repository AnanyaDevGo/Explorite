package repository

import (
	"authservice/pkg/domain"
	interfaces "authservice/pkg/repository/interface"
	"authservice/pkg/utils/models"
	"errors"
	"fmt"
	"regexp"
	"unicode"

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
func (ur *userRepository) AddProfile(id int, profile models.UserProfile) error {
	err := ur.DB.Exec(`
		INSERT INTO user_profiles (user_id, name, username, email, website, location, phone, bio)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		id, profile.Name, profile.Username, profile.Email, profile.Website, profile.Location, profile.Phone, profile.Bio).Error
	if err != nil {
		return errors.New("could not add profile")
	}

	return nil
}
func (ur *userRepository) GetProfile(id int) ([]domain.UserProfile, error) {

	var profile []domain.UserProfile
	if err := ur.DB.Raw("select * from user_profiles where user_id = ?", id).Scan(&profile).Error; err != nil {
		return []domain.UserProfile{}, errors.New("error in getting profile")
	}
	return profile, nil
}

func (ur *userRepository) ValidatePhoneNumber(phone string) bool {
	phoneNumber := phone
	pattern := `^\d{10}$`
	regex := regexp.MustCompile(pattern)
	value := regex.MatchString(phoneNumber)
	return value
}
func (ur *userRepository) ValidateAlphabets(data string) (bool, error) {
	for _, char := range data {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) {
			return false, errors.New("data contains non-alphabetic characters")
		}
	}
	return true, nil
}
