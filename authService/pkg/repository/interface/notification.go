package interfaces

import "authservice/pkg/utils/models"

type NotificationRepository interface {
	UserData(userId int) (models.UserDatas, error)
}
