package client

import (
	"context"
	"fmt"
	"os"

	logging "postservice/Logging"
	"postservice/pkg/config"
	pb "postservice/pkg/pb/auth"
	"postservice/pkg/utils/models"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type authClient struct {
	Client  pb.NotificationAuthServiceClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewAuthClient(cfg *config.Config) *authClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Post.log")
	grpcConnection, err := grpc.Dial(cfg.Explorite_Auth, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect", err)
	}
	grpcClient := pb.NewNotificationAuthServiceClient(grpcConnection)

	return &authClient{
		Client:  grpcClient,
		Logger:  logger,
		LogFile: logFile,
	}
}

func (ad *authClient) UserData(userid int) (models.UserData, error) {
	ad.Logger.Info("UserData at authClient started")

	data, err := ad.Client.UserData(context.Background(), &pb.UserDataRequest{
		Id: int64(userid),
	})

	if err != nil {
		ad.Logger.Error("error from authClient", err)
		return models.UserData{}, err
	}
	ad.Logger.Info("UserData at authClient success")
	return models.UserData{
		UserId:   int(data.Id),
		Username: data.Username,
	}, nil
}
