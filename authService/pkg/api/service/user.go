package service

import (
	pb "authservice/pkg/pb/user"
	interfaces "authservice/pkg/usecase/interface"
	"authservice/pkg/utils/models"
	"context"
)

type UserServer struct {
	userUseCase interfaces.UserUseCase
	pb.UnimplementedUserServer
}

func NewUserServer(useCase interfaces.UserUseCase) pb.UserServer {
	return &UserServer{
		userUseCase: useCase,
	}
}

func (us *UserServer) UserSignup(ctx context.Context, req *pb.UserSignupRequest) (*pb.UserSignupResponse, error) {
	userSignup := models.UserSignup{
		Email:       req.Email,
		Password:    req.Password,
		Firstname:   req.Firstname,
		Lastname:    req.Lastname,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		Bio:         req.Bio,
	}

	tokenUser, err := us.userUseCase.UserSignUp(userSignup)
	if err != nil {
		return nil, err
	}

	return &pb.UserSignupResponse{
		Status: 201,
		Token:  tokenUser.Token,
	}, nil
}

func (us *UserServer) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	userLogin := models.UserLogin{
		Email:    req.Email,
		Password: req.Password,
	}

	tokenUser, err := us.userUseCase.LoginHandler(userLogin)
	if err != nil {
		return nil, err
	}

	return &pb.UserLoginResponse{
		Status: 200,
		Token:  tokenUser.Token,
	}, nil
}
