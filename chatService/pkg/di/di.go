package di

import (
	server "chatservice/pkg/api"
	"chatservice/pkg/api/service"
	"chatservice/pkg/client"
	"chatservice/pkg/config"
	"chatservice/pkg/db"
	"chatservice/pkg/repository"
	"chatservice/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {
	database, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	chatRepository := repository.NewChatRepository(database)
	authClient := client.NewAuthClient(&cfg)

	chatUseCase := usecase.NewChatUseCase(chatRepository, authClient.Client)

	serviceServer := service.NewChatServer(chatUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, serviceServer)
	if err != nil {
		return nil, err
	}

	go chatUseCase.MessageConsumer()
	return grpcServer, nil
}
