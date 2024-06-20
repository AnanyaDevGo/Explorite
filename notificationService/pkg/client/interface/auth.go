package client

import "notificationService/pkg/utils/models"

type Newauthclient interface {
	UserData(userid int) (models.UserData, error)
}
