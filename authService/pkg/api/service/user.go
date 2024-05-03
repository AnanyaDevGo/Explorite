package service

import (
	pb "authservice/pkg/pb/user"
	interfaces "authservice/pkg/usecase/interface"
	"authservice/pkg/utils/models"
	"context"
	"errors"
	"fmt"
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

func (us *UserServer) UserSignUp(ctx context.Context, req *pb.UserSignupRequest) (*pb.UserSignupResponse, error) {
	fmt.Println("usrrrrrrrrrrrrrrrrrr", req)
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

	fmt.Println("token user", tokenUser)
	return &pb.UserSignupResponse{
		Status: 201,
		User: &pb.UserDetails{
			Id:          uint64(tokenUser.User.ID),
			Email:       tokenUser.User.Email,
			Firstname:   tokenUser.User.Firstname,
			Lastname:    tokenUser.User.Lastname,
			PhoneNumber: tokenUser.User.PhoneNumber,
			DateOfBirth: tokenUser.User.DateOfBirth,
			Gender:      tokenUser.User.Gender,
		},
		Token: tokenUser.Token,
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
		User: &pb.UserDetails{
			Id:          uint64(tokenUser.User.ID),
			Email:       tokenUser.User.Email,
			Firstname:   tokenUser.User.Firstname,
			Lastname:    tokenUser.User.Lastname,
			PhoneNumber: tokenUser.User.PhoneNumber,
			DateOfBirth: tokenUser.User.DateOfBirth,
			Gender:      tokenUser.User.Gender,
		},
		Token: tokenUser.Token,
	}, nil
}
func (us *UserServer) AddProfile(ctx context.Context, req *pb.AddProfileRequest) (*pb.AddProfileResponse, error) {
	profile := models.UserProfile{
		ID:       uint(req.Id),
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Website:  req.Website,
		Location: req.Location,
		Phone:    req.Phone,
		Bio:      req.Bio,
	}

	err := us.userUseCase.AddProfile(int(req.Id), profile)
	if err != nil {
		return nil, err
	}

	return &pb.AddProfileResponse{
		Status: 200,
	}, nil
}

func (us *UserServer) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	fmt.Println("hereeeee", req.Id)
	profile, err := us.userUseCase.GetProfile(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.GetProfileResponse{
		Status: 200,
		Profile: &pb.UserProfile{
			Id:    uint64(profile.ID),
			Name:  profile.Name,
			Email: profile.Email,
			Phone: profile.Phone,
			Bio:   profile.Bio,
		},
	}, nil
}
func (us *UserServer) EditProfile(ctx context.Context, req *pb.EditProfileRequest) (*pb.EditProfileResponse, error) {
	profile := models.EditProfile{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Website:  req.Website,
		Location: req.Location,
		Phone:    req.Phone,
		Bio:      req.Bio,
	}

	res, err := us.userUseCase.EditProfile(int(req.Id), profile)
	if err != nil {
		fmt.Println("error @ editprofile")
		return nil, err
	}

	return &pb.EditProfileResponse{
		Status:   200,
		Name:     res.Name,
		Username: res.Username,
		Email:    res.Email,
		Website:  res.Website,
		Location: res.Location,
		Phone:    res.Phone,
		Bio:      res.Bio,
	}, nil
}
func (us *UserServer) UserOTPLogin(ctx context.Context, req *pb.UserOTPLoginRequest) (*pb.UserOTPLoginResponse, error) {
	otp, err := us.userUseCase.UserOTPLogin(req.Email)
	if err != nil {
		return &pb.UserOTPLoginResponse{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}

	return &pb.UserOTPLoginResponse{
		Status: 200,
		Otp:    otp,
	}, nil
}

func (us *UserServer) OtpVerification(ctx context.Context, req *pb.OtpVerificationRequest) (*pb.OtpVerificationResponse, error) {
	verified, err := us.userUseCase.OtpVerification(req.Email, req.Otp)
	if err != nil {
		return &pb.OtpVerificationResponse{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}

	return &pb.OtpVerificationResponse{
		Status:   200,
		Verified: verified,
	}, nil
}
func (us *UserServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	userID := int(req.UserId)
	oldPassword := req.OldPassword
	newPassword := req.NewPassword
	rePassword := req.RePassword

	err := us.userUseCase.ChangePassword(userID, oldPassword, newPassword, rePassword)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			return &pb.ChangePasswordResponse{
				Status: 400,
				Error:  "user not found",
			}, nil
		}
		if errors.Is(err, models.ErrInvalidPassword) {
			return &pb.ChangePasswordResponse{
				Status: 400,
				Error:  "invalid password",
			}, nil
		}
		return &pb.ChangePasswordResponse{
			Status: 500,
			Error:  "internal server error",
		}, nil
	}

	return &pb.ChangePasswordResponse{
		Status: 200,
	}, nil
}
