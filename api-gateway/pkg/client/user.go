package client

import (
	"context"
	"errors"
	"fmt"

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
	fmt.Println("client")
	fmt.Println("auth", cfg.ExploriteAuth)
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
	fmt.Println("usr", userDetails)
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
	fmt.Println("err", err)
	if err != nil {
		return models.TokenUser{}, err
	}
	fmt.Println("eeeeeeeeeeeeee", resp)
	return models.TokenUser{
		User: models.UserDetailResponse{
			ID:          int(resp.User.Id),
			Email:       resp.User.Email,
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
			ID:          int(resp.User.Id),
			Email:       resp.User.Email,
			Firstname:   resp.User.Firstname,
			Lastname:    resp.User.Lastname,
			PhoneNumber: resp.User.PhoneNumber,
			DateOfBirth: resp.User.DateOfBirth,
			Gender:      resp.User.Gender,
		},
		Token: resp.Token,
	}, nil
}
func (uc *userClient) UserOTPLogin(email string) (string, error) {
	resp, err := uc.Client.UserOTPLogin(context.Background(), &pb.UserOTPLoginRequest{
		Email: email,
	})
	if err != nil {
		return "", err
	}

	if resp.Status != 200 {
		return "", errors.New(resp.Error)
	}

	return resp.Otp, nil
}
func (uc *userClient) OtpVerification(email, otp string) (bool, error) {
	resp, err := uc.Client.OtpVerification(context.Background(), &pb.OtpVerificationRequest{
		Email: email,
		Otp:   otp,
	})
	if err != nil {
		return false, err
	}

	if resp.Status != 200 {
		return false, errors.New(resp.Error)
	}

	return true, nil
}
func (uc *userClient) AddProfile(id int, profile models.UserProfile) error {
	_, err := uc.Client.AddProfile(context.Background(), &pb.AddProfileRequest{
		Id:       int32(id),
		Name:     profile.Name,
		Username: profile.Username,
		Email:    profile.Email,
		Website:  profile.Website,
		Location: profile.Location,
		Phone:    profile.Phone,
		Bio:      profile.Bio,
	})
	if err != nil {
		return err
	}
	return nil
}
func (uc *userClient) GetProfile(id int) (models.UserProfile, error) {
	fmt.Println("id hereeee", id)
	resp, err := uc.Client.GetProfile(context.Background(), &pb.GetProfileRequest{Id: int32(id)})
	if err != nil {
		return models.UserProfile{}, err
	}

	var profile models.UserProfile
	profile.ID = uint(resp.Profile.Id)
	profile.Name = resp.Profile.Name
	profile.Email = resp.Profile.Email
	profile.Phone = resp.Profile.Phone
	profile.Bio = resp.Profile.Bio
	return profile, nil
}
func (uc *userClient) EditProfile(id int, profile models.EditProfile) error {
	_, err := uc.Client.EditProfile(context.Background(), &pb.EditProfileRequest{
		Id:       int32(id),
		Name:     profile.Name,
		Username: profile.Username,
		Email:    profile.Email,
		Website:  profile.Website,
		Location: profile.Location,
		Phone:    profile.Phone,
		Bio:      profile.Bio,
	})
	if err != nil {
		return err
	}
	return nil
}
func (uc *userClient) ChangePassword(userID int, oldPassword, newPassword, rePassword string) error {

	_, err := uc.Client.ChangePassword(context.Background(), &pb.ChangePasswordRequest{
		UserId:      int32(userID),
		OldPassword: oldPassword,
		NewPassword: newPassword,
		RePassword:  rePassword,
	})
	if err != nil {
		return err
	}
	return nil
}

// func (uc *userClient) SendFollowReq(Id,userID int) error{
// 	_, err := uc.Client.SendFollowReq(context.Background(), &pb.SendFollowReqRequest{
// 		Id: int32(Id),
// 		UserId:  int32(userID),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// }