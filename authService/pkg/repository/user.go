package repository

import (
	"authservice/pkg/domain"
	interfaces "authservice/pkg/repository/interface"
	"authservice/pkg/utils/models"
	"errors"
	"fmt"
	"regexp"
	"time"
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
	if err := ur.DB.Raw("INSERT INTO users (firstname, lastname, email, password, phone_number, date_of_birth, gender,bio) VALUES (?,?,?,  ?, ?, ?, ?, ?) RETURNING id, firstname, lastname, email, phone_number, date_of_birth, gender,bio", userDetails.Firstname, userDetails.Lastname, userDetails.Email, userDetails.Password, userDetails.PhoneNumber, userDetails.DateOfBirth, userDetails.Gender, userDetails.Bio).Scan(&model).Error; err != nil {
		return models.UserDetailResponse{}, err
	}
	fmt.Println("inside", model.Email)
	return model, nil
}
func (r *userRepository) CheckUserExistsByEmail(email string) (*domain.User, error) {
    var user domain.User
    result := r.DB.Where("email = ?", email).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, result.Error
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
func (ur *userRepository) GetUserName(email string) (string, error) {
	var name string
	err := ur.DB.Raw("select firstname from users where email = ?", email).Scan(&name).Error
	if err != nil {
		return "", err
	}
	return name, nil
}
func (ur *userRepository) DeleteRecentOtpRequestsBefore5min() error {
	query := "DELETE FROM user_otp_logins WHERE expiration < CURRENT_TIMESTAMP - INTERVAL '5 minutes';"
	err := ur.DB.Exec(query).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) TemporarySavingUserOtp(otp int, userEmail string, expiration time.Time) error {

	query := `INSERT INTO user_otp_logins (email, otp, expiration) VALUES ($1, $2, $3)`
	err := ur.DB.Exec(query, userEmail, otp, expiration).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) VerifyOTP(email, otp string) (bool, error) {
	query := "SELECT COUNT(*) FROM user_otp_logins WHERE email = ? AND otp = ? AND expiration > CURRENT_TIMESTAMP;"
	var count int64
	if err := ur.DB.Raw(query, email, otp).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
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
func (ur *userRepository) GetProfile(id int) (models.UserProfile, error) {

	fmt.Println("id repo", id)

	var profile models.UserProfile
	profile.ID = uint(id)
	if err := ur.DB.Raw("select firstname as name, email, phone_number as phone, bio from users where id = ?", id).Scan(&profile).Error; err != nil {
		return models.UserProfile{}, errors.New("error in getting profile")
	}
	fmt.Println("profile id", id)
	fmt.Println("profileeeeee", profile)
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
func (ur *userRepository) IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
func (ur *userRepository) IsValidWebsite(website string) bool {
	fmt.Println("wenbsite uuuuuuuuuu")
	websiteRegex := `^(https?|ftp):\/\/[^\s\/$.?#].[^\s]*$`
	match, _ := regexp.MatchString(websiteRegex, website)
	return match
}


func (ur *userRepository) EditProfile(id int, user models.EditProfile) (models.EditProfile, error) {
	var result models.EditProfile

	fmt.Println("edit profile")
	args := []interface{}{}
	query := "update users set"

	if user.Email != "" {
		query += " email = $1,"

		args = append(args, user.Email)
	}

	// if user.Name != "" {
	// 	query += " name = $2,"
	// 	args = append(args, user.Name)
	// }

	if user.Phone != "" {
		query += " phone_number = $2,"

		args = append(args, user.Phone)
	}

	query = query[:len(query)-1] + " where id = $3"

	args = append(args, id)

	err := ur.DB.Exec(query, args...).Error
	if err != nil {
		return models.EditProfile{}, err
	}
	query2 := "select firstname as name, email, phone_number as phone, bio from users where id = ?"
	if err := ur.DB.Raw(query2, id).Scan(&result).Error; err != nil {
		return models.EditProfile{}, err
	}

	return result, nil

}
func (ur *userRepository) ChangePassword(id int, password string) error {

	if id <= 0 {
		return errors.New("negative or zero values are not allowed")
	}

	err := ur.DB.Exec("UPDATE users SET password=$1 WHERE id=$2", password, id).Error
	if err != nil {
		return err
	}

	return nil

}

func (u *userRepository) GetPassword(id int) (string, error) {

	if id <= 0 {
		return "", errors.New("negative or zero values are not allowed")
	}

	var userPassword string
	err := u.DB.Raw("select password from users where id = ?", id).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil

}