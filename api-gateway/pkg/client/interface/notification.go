package interfaces

import "ExploriteGateway/pkg/utils/models"

type NotificationClient interface {
	GetNotification(userid int, req models.NotificationPagination) ([]models.NotificationResponse, error)
}
