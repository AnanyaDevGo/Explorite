package handler

import (
	interfaces "ExploriteGateway/pkg/client/interface"
	"ExploriteGateway/pkg/utils/models"
	"ExploriteGateway/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	GRPC_Client interfaces.NotificationClient
}

func NewNotificationHandler(notificationClient interfaces.NotificationClient) *NotificationHandler {
	return &NotificationHandler{
		GRPC_Client: notificationClient,
	}
}

func (ad *NotificationHandler) GetNotification(c *gin.Context) {
	var notificationRequest models.NotificationPagination
	if err := c.ShouldBindJSON(&notificationRequest); err != nil {
		errorres := response.ClientResponse(http.StatusBadRequest, "details give in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
		return
	}

	id_string, _ := c.Get("id")
	id, _ := id_string.(int)

	result, errs := ad.GRPC_Client.GetNotification(id, notificationRequest)
	if errs != nil {
		errss := response.ClientResponse(http.StatusBadRequest, "error in getting notification", nil, errs.Error())
		c.JSON(http.StatusBadRequest, errss)
		return
	}
	succesres := response.ClientResponse(http.StatusOK, "successfully got all notification", result, nil)
	c.JSON(http.StatusOK, succesres)
}
