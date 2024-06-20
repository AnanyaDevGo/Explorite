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

	notiRepository := repository.NewnotiRepository(gormDB)
	noticlient := client.NewAuthClient(&cfg)
	noriUseCase := usecase.NewnotiUsecase(notiRepository, noticlient)
	notiServiceServer := service.NewnotiServer(noriUseCase)
	grpcserver, err := server.NewGRPCServer(cfg, notiServiceServer)

	if err != nil {
		return &server.Server{}, err
	}
	go noriUseCase.ConsumeNotification()
	return grpcserver, nil
}
