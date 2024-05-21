package di

import (
	server "ExploriteGateway/pkg/api"
	"ExploriteGateway/pkg/api/handler"
	"ExploriteGateway/pkg/client"
	"ExploriteGateway/pkg/config"
	"ExploriteGateway/pkg/helper"
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
	postClient, err := client.NewPostClient(cfg)
	if err != nil {
		return nil, err
	}
	chatClient := client.NewChatClient(cfg)

	adminHandler := handler.NewAdminHandler(adminClient)
	userHandler := handler.NewUserHandler(userClient)
	postHandler := handler.NewPostHandler(postClient)
	chatHandler := handler.NewChatHandler(chatClient,helper.NewHelper(&cfg))

	serverHTTP := server.NewServerHTTP(adminHandler, userHandler, postHandler, chatHandler)

	return serverHTTP, nil
}
