package interfaces

import "authservice/pkg/utils/models"

type NotificationUsecase interface {
	UserData(userId int) (models.UserDatas, error)
}
