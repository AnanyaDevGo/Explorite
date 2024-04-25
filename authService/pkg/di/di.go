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
	gornDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	adminRepository := repository.NewAdminRepository(gornDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminServiceServer := service.NewAdminServer(adminUseCase)

	userRepository := repository.NewUserRepository(gornDB)
	userUsecase := usecase.NewUserUseCase(userRepository)
	userServiceServer := service.NewUserServer(userUsecase)

	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer, userServiceServer)
	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil

}
