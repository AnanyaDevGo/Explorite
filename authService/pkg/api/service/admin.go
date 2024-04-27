package service

import (
	pb "authservice/pkg/pb/admin"
	interfaces "authservice/pkg/usecase/interface"
	"authservice/pkg/utils/models"
	"context"
	"fmt"
	"strconv"
)

type AdminServer struct {
	adminUseCase interfaces.AdminUseCase
	pb.UnimplementedAdminServer
}

func NewAdminServer(useCase interfaces.AdminUseCase) pb.AdminServer {

	return &AdminServer{
		adminUseCase: useCase,
	}

}
func (ad *AdminServer) AdminSignup(ctx context.Context, req *pb.AdminSignupRequest) (*pb.AdminSignupResponse, error) {
	adminSignup := models.AdminSignUp{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
	}

	fmt.Println("service", adminSignup)

	res, err := ad.adminUseCase.AdminSignUp(adminSignup)
	if err != nil {
		return &pb.AdminSignupResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(res.Admin.ID),
		Firstname: res.Admin.Firstname,
		Lastname:  res.Admin.Lastname,
		Email:     res.Admin.Email,
	}
	return &pb.AdminSignupResponse{
		Status:       201,
		AdminDetails: adminDetails,
		Token:        res.Token,
	}, nil
}

func (ad *AdminServer) AdminLogin(ctx context.Context, Req *pb.AdminLoginInRequest) (*pb.AdminLoginResponse, error) {
	adminLogin := models.AdminLogin{
		Email:    Req.Email,
		Password: Req.Password,
	}
	admin, err := ad.adminUseCase.LoginHandler(adminLogin)
	if err != nil {
		return &pb.AdminLoginResponse{}, err
	}
	adminDetails := &pb.AdminDetails{
		Id:        uint64(admin.Admin.ID),
		Firstname: admin.Admin.Firstname,
		Lastname:  admin.Admin.Lastname,
		Email:     admin.Admin.Email,
	}
	return &pb.AdminLoginResponse{
		Status:       200,
		AdminDetails: adminDetails,
		Token:        admin.Token,
	}, nil
}

func (as *AdminServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	page := int(req.Page)

	users, err := as.adminUseCase.GetUsers((page))
	if err != nil {
		return nil, err
	}

	var userDetails []*pb.UserDetailsAtAdmin
	for _, user := range users {
		userID := int32(user.Id)
		userDetails = append(userDetails, &pb.UserDetailsAtAdmin{
			Id:          userID,
			Name:        user.Name,
			Email:       user.Email,
			Phone:       user.Phone,
			BlockStatus: user.BlockStatus,
		})
	}

	return &pb.GetUsersResponse{
		Users: userDetails,
	}, nil
}

func (as *AdminServer) BlockUser(ctx context.Context, req *pb.BlockUserRequest) (*pb.BlockUserResponse, error) {
	userID := strconv.FormatUint(req.UserId, 10)

	err := as.adminUseCase.BlockUser(userID)
	if err != nil {
		return &pb.BlockUserResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.BlockUserResponse{
		Success: true,
	}, nil
}

func (as *AdminServer) UnBlockUser(ctx context.Context, req *pb.UnBlockUserRequest) (*pb.UnBlockUserResponse, error) {
	userID := strconv.FormatUint(req.UserId, 10)

	err := as.adminUseCase.BlockUser(userID)
	if err != nil {
		return &pb.UnBlockUserResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.UnBlockUserResponse{
		Success: true,
	}, nil
}
