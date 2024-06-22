package client

import (
	"fmt"
	"os"

	logging "notificationService/Logging"
	"notificationService/pkg/config"
	pb "notificationService/pkg/pb/auth"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type authClient struct {
	Client  pb.NotificationAuthServiceClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewAuthClient(cfg *config.Config) *authClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/notificationService.log")
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
