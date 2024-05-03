package di

import (
	server "ExploriteGateway/pkg/api"
	"ExploriteGateway/pkg/api/handler"
	"ExploriteGateway/pkg/client"
	"ExploriteGateway/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {
	adminClient, err := client.NewAdminClient(cfg)
	if err != nil {
		return nil, err
	}
	userClient, err := client.NewUserClient(cfg)
	if err != nil {
		return nil, err
	}

	adminHandler := handler.NewAdminHandler(adminClient)

	userHandler := handler.NewUserHandler(userClient)

	serverHTTP := server.NewServerHTTP(adminHandler, userHandler)

	return serverHTTP, nil
}