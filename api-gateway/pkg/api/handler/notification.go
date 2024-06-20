package handler

import interfaces "ExploriteGateway/pkg/client/interface"

type NotificationHandler struct{
	GRPC_Client interfaces.NotificationClient
}
func NewNotificationHandler(notificationClient interfaces.NotificationClient) *NotificationHandler {
	return &NotificationHandler{
		GRPC_Client: notificationClient,
	}
}
