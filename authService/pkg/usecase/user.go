package usecase

import (
	"authservice/pkg/domain"
	"authservice/pkg/helper"
	"authservice/pkg/randomnumbergenerator"
	interfaces "authservice/pkg/repository/interface"
	services "authservice/pkg/usecase/interface"
	"authservice/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepository interfaces.UserRepository
}

func NewUserUseCase(repository interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepository: repository,
	}
}


var InternalError = "Internal Server Error"

func (uu *userUseCase) UserSignUp(userDetails models.UserSignup) (*domain.TokenUser, error) {
	email, err := uu.userRepository.CheckUserExistsByEmail(userDetails.Email)
	if err != nil {
		return &domain.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return &domain.TokenUser{}, errors.New("user with this email already exists")
	}

	hashPassword, err := helper.PasswordHash(userDetails.Password)
	if err != nil {
		return &domain.TokenUser{}, errors.New("error in hashing password")
	}
	userDetails.Password = hashPassword

	userData, err := uu.userRepository.UserSignUp(userDetails)
	if err != nil {
		return &domain.TokenUser{}, errors.New("could not add the user")
	}

	tokenString, err := helper.GenerateTokenUser(userData)
	if err != nil {
		return &domain.TokenUser{}, err
	}

	return &domain.TokenUser{
		User:  userData,
		Token: tokenString,
	}, nil
}

func (uu *userUseCase) LoginHandler(userDetails models.UserLogin) (*domain.TokenUser, error) {
	email, err := uu.userRepository.CheckUserExistsByEmail(userDetails.Email)
	if err != nil {
		return &domain.TokenUser{}, errors.New("error with server")
	}
	if email == nil {
		return &domain.TokenUser{}, errors.New("email doesn't exist")
	}

	userData, err := uu.userRepository.FindUserByEmail(userDetails)
	if err != nil {
		return &domain.TokenUser{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(userDetails.Password))
	if err != nil {
		return &domain.TokenUser{}, errors.New("password not matching")
	}

	var userDetailResponse models.UserDetailResponse
	err = copier.Copy(&userDetailResponse, &userData)
	if err != nil {
		return &domain.TokenUser{}, err
	}

	tokenString, err := helper.GenerateTokenUser(userDetailResponse)
	if err != nil {
		return &domain.TokenUser{}, err
	}

	return &domain.TokenUser{
		User:  userDetailResponse,
		Token: tokenString,
	}, nil
}
func (u *userUseCase) AddProfile(id int, profile models.UserProfile) error {

	if profile.Name == "" || profile.Username == "" || profile.Email == "" || profile.Website == "" || profile.Location == "" || profile.Phone == "" || profile.Bio == "" {
		return errors.New("fields cannot be empty")
	}
	ok, err := u.userRepository.ValidateAlphabets(profile.Name)
	if err != nil {
		return errors.New("invalid format for name")
	}
	if !ok {
		return errors.New("invalid format for name")
	}
	if !u.userRepository.IsValidEmail(profile.Email) {
		return errors.New("invalid email format")
	}

	if !u.userRepository.IsValidWebsite(profile.Website) {
		return errors.New("invalid website format")
	}

	phonenumber := u.userRepository.ValidatePhoneNumber(profile.Phone)
	if !phonenumber {
		return errors.New("invalid phone")
	}
	addProfileErr := u.userRepository.AddProfile(id, profile)
	if addProfileErr != nil {
		return errors.New("error in adding profile")
	}
	return nil
}
func (u *userUseCase) GetProfile(id int) (models.UserProfile, error) {
	fmt.Println("id usec ase", id)

	profile, err := u.userRepository.GetProfile(id)
	if err != nil {
		return models.UserProfile{}, errors.New("error in getting profile")
	}
	fmt.Println("profile usecase", profile)
	return profile, nil
}

func (u *userUseCase) EditProfile(id int, user models.EditProfile) (models.EditProfile, error) {
	fmt.Println("profileeee")

	if user.Name == "" || user.Username == "" || user.Email == "" || user.Website == "" || user.Location == "" || user.Phone == "" || user.Bio == "" {
		return models.EditProfile{}, errors.New("all fields must be filled")
	}
	fmt.Println("profile2222")

	ok, err := u.userRepository.ValidateAlphabets(user.Name)
	if err != nil {
		return models.EditProfile{}, errors.New("invalid format for name")
	}
	if !ok {
		return models.EditProfile{}, errors.New("invalid format for name")
	}
	fmt.Println("profile33333")
	if !u.userRepository.IsValidEmail(user.Email) {
		return models.EditProfile{}, errors.New("invalid email format")
	}
	fmt.Println("profile4444444")

	// if !u.userRepository.IsValidWebsite(user.Website) {
	// 	fmt.Println("websiteeeeeeeee")
	// 	err := errors.New("invalid website format")
	// 	return models.EditProfile{}, err
	// }
	fmt.Println("profile5555")

	phonenumber := u.userRepository.ValidatePhoneNumber(user.Phone)
	if !phonenumber {
		return models.EditProfile{}, errors.New("invalid phone")
	}
	fmt.Println("profile6666")
	result, err := u.userRepository.EditProfile(id, user)
	fmt.Println("resulttttt", result, err)

	if err != nil {
		fmt.Println("erorrrrrrrrrr")
		return models.EditProfile{}, err
	}

	return result, nil
}
func (r *userUseCase) UserOTPLogin(email string) (string, error) {
	otp := randomnumbergenerator.RandomNumber()

	otpString := strconv.Itoa(otp)

	errRemv := r.userRepository.DeleteRecentOtpRequestsBefore5min()
	if errRemv != nil {
		return "", errRemv
	}

	expiration := time.Now().Add(5 * time.Minute)
	errTempSave := r.userRepository.TemporarySavingUserOtp(otp, email, expiration)
	if errTempSave != nil {
		fmt.Println("Can't save temporary data for OTP verification in DB")
		return "", errors.New("OTP verification down, please try again later")
	}
	// name, err := r.userRepository.GetUserName(email)
	// if err != nil {
	// 	return "", err
	// }
	// err = helper.SendVerificationEmailWithOtp(otp, email, name)
	// if err != nil {
	// 	return "",err
	// }

	return otpString, nil
}

func (r *userUseCase) OtpVerification(email, otp string) (bool, error) {
	verified, err := r.userRepository.VerifyOTP(email, otp)
	if err != nil {
		return false, err
	}
	return verified, nil
}
func (r *userUseCase) ChangePassword(id int, old string, password string, repassword string) error {
    if password == "" {
        return errors.New("password cannot be empty")
    }

    userPassword, err := r.userRepository.GetPassword(id)
    if err != nil {
        return errors.New(InternalError)
    }

    err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(old))
    if err != nil {
        return errors.New("password incorrect")
    }

    if password != repassword {
        return errors.New("passwords does not match")
    }

    newpassword, err := helper.PasswordHash(password)
    if err != nil {
        return errors.New("error in hashing password")
    }

    return r.userRepository.ChangePassword(id, string(newpassword))
}
// func (uu *userUseCase) SavePost(postID int) error {
//     err := uu.userRepository.SavePost(postID)
//     if err != nil {
//         return err
//     }
//     return nil
// }

// func (uu *userUseCase) UnSavePost(postID int) error {
//     err := uu.userRepository.UnSavePost(postID)
//     if err != nil {
//         return err
//     }
//     return nil
// }
