package client

import (
	interfaces "ExploriteGateway/pkg/client/interface"
	"ExploriteGateway/pkg/config"
	pb "ExploriteGateway/pkg/pb/admin"
	"ExploriteGateway/pkg/utils/models"
	"context"
	"strconv"

	"google.golang.org/grpc"
)

type adminClient struct {
	Client pb.AdminClient
}

func NewAdminClient(cfg config.Config) (interfaces.AdminClient, error) {
	grpcConnection, err := grpc.Dial(cfg.ExploriteAuth, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	grpcClient := pb.NewAdminClient(grpcConnection)

	return &adminClient{
		Client: grpcClient,
	}, nil
}

func (ad *adminClient) AdminSignUp(admindeatils models.AdminSignUp) (models.TokenAdmin, error) {
	admin, err := ad.Client.AdminSignup(context.Background(), &pb.AdminSignupRequest{
		Firstname: admindeatils.Firstname,
		Lastname:  admindeatils.Lastname,
		Email:     admindeatils.Email,
		Password:  admindeatils.Password,
	})
	if err != nil {
		return models.TokenAdmin{}, err
	}
	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID:        uint(admin.AdminDetails.Id),
			Firstname: admin.AdminDetails.Firstname,
			Lastname:  admin.AdminDetails.Lastname,
			Email:     admin.AdminDetails.Email,
		},
		Token: admin.Token,
	}, nil
}

func (ad *adminClient) AdminLogin(adminDetails models.AdminLogin) (models.TokenAdmin, error) {
	admin, err := ad.Client.AdminLogin(context.Background(), &pb.AdminLoginInRequest{
		Email:    adminDetails.Email,
		Password: adminDetails.Password,
	})

	if err != nil {
		return models.TokenAdmin{}, err
	}
	return models.TokenAdmin{
		Admin: models.AdminDetailsResponse{
			ID:        uint(admin.AdminDetails.Id),
			Firstname: admin.AdminDetails.Firstname,
			Lastname:  admin.AdminDetails.Lastname,
			Email:     admin.AdminDetails.Email,
		},
		Token: admin.Token,
	}, nil
}
func (ad *adminClient) GetUsers(page int) ([]models.UserDetailsAtAdmin, error) {
	res, err := ad.Client.GetUsers(context.Background(), &pb.GetUsersRequest{Page: int32(page)})
	if err != nil {
		return nil, err
	}

	var userDetails []models.UserDetailsAtAdmin
	for _, user := range res.Users {
		userDetails = append(userDetails, models.UserDetailsAtAdmin{
			Id:          int(user.Id),
			Name:        user.Name,
			Email:       user.Email,
			Phone:       user.Phone,
			BlockStatus: user.BlockStatus,
		})
	}

	return userDetails, nil
}
func (ad *adminClient) BlockUser(id string) error {
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	_, err = ad.Client.BlockUser(context.Background(), &pb.BlockUserRequest{
		UserId: userId,
	})
	if err != nil {
		return err
	}
	return nil
}
func (ad *adminClient) UnBlockUser(id string) error {
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	_, err = ad.Client.UnBlockUser(context.Background(), &pb.UnBlockUserRequest{
		UserId: userId,
	})
	if err != nil {
		return err
	}
	return nil
}
