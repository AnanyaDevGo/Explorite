package client

import (
	"context"

	"google.golang.org/grpc"

	interfaces "ExploriteGateway/pkg/client/interface"
	"ExploriteGateway/pkg/config"
	pb "ExploriteGateway/pkg/pb/user"
	"ExploriteGateway/pkg/utils/models"
)

type userClient struct {
	Client pb.UserClient
}

func NewUserClient(cfg config.Config) (interfaces.UserClient, error) {
	grpcConnection, err := grpc.Dial(cfg.ExploriteAuth, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	grpcClient := pb.NewUserClient(grpcConnection)

	return &userClient{
		Client: grpcClient,
	}, nil
}

func (uc *userClient) UserSignUp(userDetails models.UserSignup) (models.TokenUser, error) {
	resp, err := uc.Client.UserSignUp(context.Background(), &pb.UserSignupRequest{
		Email:       userDetails.Email,
		Password:    userDetails.Password,
		Firstname:   userDetails.Firstname,
		Lastname:    userDetails.Lastname,
		PhoneNumber: userDetails.PhoneNumber,
		DateOfBirth: userDetails.DateOfBirth,
		Gender:      userDetails.Gender,
		Bio:         userDetails.Bio,
	})
	if err != nil {
		return models.TokenUser{}, err
	}

	return models.TokenUser{
		User: models.UserDetailResponse{
			Email:       resp.User.Email,
			Password:    resp.User.Password,
			Firstname:   resp.User.Firstname,
			Lastname:    resp.User.Lastname,
			PhoneNumber: resp.User.PhoneNumber,
			DateOfBirth: resp.User.DateOfBirth,
			Gender:      resp.User.Gender,
		},
		Token: resp.Token,
	}, nil
}

func (uc *userClient) UserLogin(userDetails models.UserLogin) (models.TokenUser, error) {
	resp, err := uc.Client.UserLogin(context.Background(), &pb.UserLoginRequest{
		Email:    userDetails.Email,
		Password: userDetails.Password,
	})
	if err != nil {
		return models.TokenUser{}, err
	}

	return models.TokenUser{
		User: models.UserDetailResponse{
			Email:       resp.User.Email,
			Password:    resp.User.Password,
			Firstname:   resp.User.Firstname,
			Lastname:    resp.User.Lastname,
			PhoneNumber: resp.User.PhoneNumber,
			DateOfBirth: resp.User.DateOfBirth,
			Gender:      resp.User.Gender,
		},
		Token: resp.Token,
	}, nil
}
