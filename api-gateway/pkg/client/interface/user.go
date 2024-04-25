package interfaces

import "ExploriteGateway/pkg/utils/models"

type UserClient interface {
	UserSignUp(userDetails models.UserSignup) (models.TokenUser, error)
	UserLogin(userDetails models.UserLogin) (models.TokenUser, error)
}
