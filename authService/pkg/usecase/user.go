package usecase

import (
	"authservice/pkg/domain"
	"authservice/pkg/helper"
	interfaces "authservice/pkg/repository/interface"
	services "authservice/pkg/usecase/interface"
	"authservice/pkg/utils/models"
	"errors"

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
