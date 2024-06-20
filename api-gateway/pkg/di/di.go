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
	notificationClient := client.NewNotificationClient(cfg)

	adminHandler := handler.NewAdminHandler(adminClient)
	userHandler := handler.NewUserHandler(userClient)
	postHandler := handler.NewPostHandler(postClient)
	chatHandler := handler.NewChatHandler(chatClient, helper.NewHelper(&cfg))
	videocallHandler := handler.NewVideoCallHandler()
	notificationhandler := handler.NewNotificationHandler(notificationClient)

	serverHTTP := server.NewServerHTTP(adminHandler, userHandler, postHandler, chatHandler, videocallHandler, notificationhandler)

	return serverHTTP, nil
}
