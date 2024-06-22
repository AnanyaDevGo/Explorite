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

	notiRepository := repository.NewNotificationRepository(gormDB)
	noticlient := client.NewAuthClient(&cfg)
	notiUseCase := usecase.NewNotificationUsecase(notiRepository, noticlient)
	notiServiceServer := service.NewNotificationServer(notiUseCase)
	grpcserver, err := server.NewGRPCServer(cfg, notiServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	go notiUseCase.ConsumeNotification()
	return grpcserver, nil
}
