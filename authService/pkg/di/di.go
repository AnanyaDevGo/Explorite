package di

import (
	server "authservice/pkg/api"
	"authservice/pkg/api/service"
	config "authservice/pkg/conifg"
	"authservice/pkg/db"
	"authservice/pkg/repository"
	"authservice/pkg/usecase"
)

func IntializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminServiceServer := service.NewAdminServer(adminUseCase)

	userRepository := repository.NewUserRepository(gormDB)
	userUsecase := usecase.NewUserUseCase(userRepository)
	userServiceServer := service.NewUserServer(userUsecase)
	
	notificationRepository := repository.NewNotificationRepository(gormDB)
	notificationUseCase := usecase.NewNotificationUseCase(notificationRepository)
	notificationServiceServer := service.NewNotificationServer(notificationUseCase)

	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer, userServiceServer, notificationServiceServer)
	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil

}
