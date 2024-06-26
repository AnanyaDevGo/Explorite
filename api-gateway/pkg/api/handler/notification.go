package handler

import (
	"net/http"
	"os"
	"strconv"

	logging "ExploriteGateway/Logging"
	interfaces "ExploriteGateway/pkg/client/interface"
	"ExploriteGateway/pkg/utils/models"
	"ExploriteGateway/pkg/utils/response"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

type NotificationHandler struct {
	GRPC_Client interfaces.NotificationClient
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewNotificationHandler(grpc_client interfaces.NotificationClient) *NotificationHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/ExploriteGateway.log")
	return &NotificationHandler{
		GRPC_Client: grpc_client,
		Logger:      logger,
		LogFile:     logFile,
	}
}

func (n *NotificationHandler) GetNotification(c *gin.Context) {
	n.Logger.Info("GetNotification at NotificationHandler started")

	pageStr := c.DefaultQuery("limit", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		n.Logger.Error("page number not in right format", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	countStr := c.DefaultQuery("offset", "10")
	pageSize, err := strconv.Atoi(countStr)
	if err != nil {
		n.Logger.Error("user count in a page not in right format", err)
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	notificationRequest := models.NotificationPagination{
		Limit:  page,
		Offset: pageSize,
	}

	userID, exists := c.Get("id")
	if !exists {
		n.Logger.Error("User ID not found in JWT claims")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	UserID, _ := userID.(int)

	data, err := n.GRPC_Client.GetNotification(UserID, notificationRequest)
	if err != nil {
		n.Logger.Error("Error during GetNotification rpc call", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to get notification details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	n.Logger.Info("Successfully retrieved Notifications")
	res := response.ClientResponse(http.StatusOK, "Successfully retrieved Notifications", data, nil)
	c.JSON(http.StatusOK, res)
}

func (n *NotificationHandler) ReadNotification(c *gin.Context) {
	noificationIdStr := c.Query("notification_id")
	noificationId, err := strconv.Atoi(noificationIdStr)
	if err != nil {
		n.Logger.Error("conversion error notification id ", err)
		errs := response.ClientResponse(http.StatusBadRequest, "conversion error notification id ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	userId, ok := c.Get("id")
	if !ok {
		n.Logger.Error("User ID not found in JWT claims")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	UserID, _ := userId.(int)

	result, err := n.GRPC_Client.ReadNotification(noificationId, UserID)
	if err != nil {
		n.Logger.Error("Error during ReadNotification rpc call", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to read notification ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	n.Logger.Info(" Notification Successfully marked as read")
	res := response.ClientResponse(http.StatusOK, "Notification Successfully marked as read", result, nil)
	c.JSON(http.StatusOK, res)

}

func (n *NotificationHandler) MarkAllAsRead(c *gin.Context) {
	n.Logger.Info("MarkAllAsRead at NotificationHandler started")

	userId, ok := c.Get("id")
	if !ok {
		n.Logger.Error("User ID not found in JWT claims")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	UserID, _ := userId.(int)
	n.Logger.Info("MarkAllAsRead at client rpc  started")
	result, err := n.GRPC_Client.MarkAllAsRead(UserID)
	if err != nil {
		n.Logger.Error("Error during ReadNotification rpc call", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to read notification ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	n.Logger.Info("MarkAllAsRead at client rpc  finished")
	n.Logger.Info(" Notifications Successfully marked as read")
	res := response.ClientResponse(http.StatusOK, "Notification Successfully marked as read", result, nil)
	c.JSON(http.StatusOK, res)

}

func (n *NotificationHandler) GetAllNotifications(c *gin.Context) {
	n.Logger.Info("GetAllNotifications at NotificationHandler started")

	userId, ok := c.Get("id")
	if !ok {
		n.Logger.Error("User ID not found in JWT claims")
		errs := response.ClientResponse(http.StatusBadRequest, "User ID not found in JWT claims", nil, "")
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	UserID, _ := userId.(int)
	n.Logger.Info("GetAllNotifications at client rpc  started")
	result, err := n.GRPC_Client.GetAllNotifications(UserID)
	if err != nil {
		n.Logger.Error("Error during GetAllNotifications rpc call", err)
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to GetAllNotifications ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	n.Logger.Info("GetAllNotifications at client rpc  finished")
	n.Logger.Info(" Notifications Successfully Fetched")
	res := response.ClientResponse(http.StatusOK, "Get All Notifications Successfull ", result, nil)
	c.JSON(http.StatusOK, res)

}
