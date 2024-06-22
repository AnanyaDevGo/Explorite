package interfaces

import "ExploriteGateway/pkg/utils/models"

type NotificationClient interface {
	GetNotification(userid int, mod models.NotificationPagination) ([]models.NotificationResponse, error)
	ReadNotification(id, user_id int) (bool, error)
	MarkAllAsRead(userId int) (bool, error)
	GetAllNotifications(userId int) ([]models.AllNotificationResponse, error)
}
