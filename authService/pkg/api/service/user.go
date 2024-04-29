package service

import (
	pb "authservice/pkg/pb/user"
	interfaces "authservice/pkg/usecase/interface"
	"authservice/pkg/utils/models"
	"context"
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
	profiles, err := us.userUseCase.GetProfile(int(req.Id))
	if err != nil {
		return nil, err
	}

	var pbProfiles []*pb.UserProfile
	for _, profile := range profiles {
		pbProfile := &pb.UserProfile{
			Id:       uint64(profile.ID),
			Name:     profile.Name,
			Username: profile.Username,
			Email:    profile.Email,
			Website:  profile.Website,
			Location: profile.Location,
			Phone:    profile.Phone,
			Bio:      profile.Bio,
		}
		pbProfiles = append(pbProfiles, pbProfile)
	}

	return &pb.GetProfileResponse{
		Status:   200,
		Profiles: pbProfiles,
	}, nil
}
