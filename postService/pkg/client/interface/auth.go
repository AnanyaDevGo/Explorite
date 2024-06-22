package client

import "postservice/pkg/utils/models"

type Newauthclient interface {
	UserData(userid int) (models.UserData, error)
}
