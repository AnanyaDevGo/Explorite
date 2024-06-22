package di

import (
	server "postservice/pkg/api"
	"postservice/pkg/api/service"
	"postservice/pkg/client"
	"postservice/pkg/config"
	"postservice/pkg/db"
	"postservice/pkg/repository"
	"postservice/pkg/usecase"
)

func IntializeAPI(cfg config.Config) (*server.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	postRepository := repository.NewPostRepository(gormDB)
	postclient := client.NewAuthClient(&cfg)
	postUseCase := usecase.NewPostUseCase(postRepository, postclient)
	postServiceServer := service.NewPostServer(postUseCase)

	grpcServer, err := server.NewGRPCServer(cfg, postServiceServer)
	if err != nil {
		return &server.Server{}, err
	}
	return grpcServer, nil

}
