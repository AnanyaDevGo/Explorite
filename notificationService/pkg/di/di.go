package di

import (
	server "notificationService/pkg/api"
	"notificationService/pkg/api/service"
	"notificationService/pkg/client"
	"notificationService/pkg/config"
	"notificationService/pkg/db"
	"notificationService/pkg/repository"
	"notificationService/pkg/usecase"
)

func InitializeApi(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	notificationRepository := repository.NewnotiRepository(gormDB)
	notificationclient := client.NewAuthClient(&cfg)
	notificationUseCase := usecase.NewnotiUsecase(notificationRepository, notificationclient)
	notificationServiceServer := service.NewnotiServer(notificationUseCase)
	grpcserver, err := server.NewGRPCServer(cfg, notificationServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	go notificationUseCase.ConsumeNotification()
	return grpcserver, nil
}
