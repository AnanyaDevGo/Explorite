package interfaces

import "chatservice/pkg/utils/models"

type ChatUseCase interface {
	MessageConsumer()
	GetFriendChat(string, string, models.Pagination) ([]models.Message, error)
}
