package repository

import (
	"fmt"
	"os"

	logging "notificationService/Logging"
	interfaces "notificationService/pkg/repository/interface"
	"notificationService/pkg/utils/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type notificationRepository struct {
	DB      *gorm.DB
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewNotificationRepository(DB *gorm.DB) interfaces.NotificationRepository {
	logger, logFile := logging.InitLogrusLogger("./Logging/notificationService.log")
	return &notificationRepository{
		DB:      DB,
		Logger:  logger,
		LogFile: logFile,
	}
}

func (c *notificationRepository) StoreNotificationReq(noti models.NotificationReq) error {
	c.Logger.Info("StoreNotificationReq at notificationRepository started")
	querry := `
	INSERT INTO notifications(
	user_id,
	sender_id,
	sender_name,
	post_id,
	message,
	created_at,
	read) 
	VALUES(?,?,?,?,?,?,?)
	`
	err := c.DB.Exec(querry, noti.UserID, noti.SenderID, noti.SenderName, noti.PostID, noti.Message, noti.CreatedAt, false).Error
	if err != nil {
		c.Logger.Error("Error while storing notification request: ", err)
		return err
	}
	c.Logger.Info("StoreNotificationReq at notificationRepository completed successfully")
	return nil
}

func (c *notificationRepository) GetNotification(id int, pag models.Pagination) ([]models.Notification, error) {
	c.Logger.Info("GetNotification at notificationRepository started")

	var data []models.Notification
	if pag.Offset <= 0 {
		pag.Offset = 1
	}
	querry := `
	SELECT id,
	sender_id, message, created_at, 
	sender_id, post_id
	FROM notifications 
	WHERE user_id = ? AND read = false 
	ORDER BY created_at DESC 
	
	`
	err := c.DB.Raw(querry, id).Scan(&data).Error
	if err != nil {
		c.Logger.Error("Error while getting notifications: ", err)
		return nil, err
	}
	fmt.Println("data", data)
	c.Logger.Info("GetNotification at notificationRepository completed successfully")
	return data, nil
}

func (c *notificationRepository) ReadNotification(id int) (bool, error) {
	c.Logger.Info("ReadNotification at notificationRepository started")

	querry := `
	update notifications 
	set read = ?
	where id = ?
	`
	err := c.DB.Exec(querry, "true", id).Error
	if err != nil {
		c.Logger.Error("Error while reading notification: ", err)
		return false, err
	}
	c.Logger.Info("ReadNotification at notificationRepository completed successfully")
	return true, nil
}

func (c *notificationRepository) IsNotificationExistOnUser(id, userId int) (bool, error) {
	c.Logger.Info("IsNotificationExistOnUser at notificationRepository started")
	var count int
	fmt.Println("iddd", id, userId)
	querry := `
	select count(*) from notifications 
	where id = ? and user_id =  ?
	`
	err := c.DB.Raw(querry, id, userId).Scan(&count).Error
	if err != nil {
		c.Logger.Error("Error while reading notification: ", err)
		return false, err
	}
	c.Logger.Info("IsNotificationExistOnUser at notificationRepository completed successfully")
	return count > 0, nil
}

func (c *notificationRepository) MarkAllAsRead(userId int) (bool, error) {
	c.Logger.Info("MarkAllAsRead at notificationRepository started")

	querry := `
	update notifications 
	set read = ?
	where user_id = ?
	`
	err := c.DB.Exec(querry, "true", userId).Error
	if err != nil {
		c.Logger.Error("Error while reading notifications: ", err)
		return false, err
	}
	c.Logger.Info("MarkAllAsRead at notificationRepository completed successfully")
	return true, nil
}

func (c *notificationRepository) UnreadedNotificationExist(userId int) (bool, error) {
	c.Logger.Info("UnreadedNotificationExist at notificationRepository started")
	var count int
	querry := `
	select count(*)  from notifications 
	where user_id = ? and read = ?
	`
	err := c.DB.Raw(querry, userId, "false").Scan(&count).Error
	if err != nil {
		c.Logger.Error("Error while reading notifications: ", err)
		return false, err
	}
	c.Logger.Info("UnreadedNotificationExist at notificationRepository completed successfully")
	return count > 0, nil
}

func (c *notificationRepository) GetAllNotifications(userId int) ([]models.AllNotificationResponse, error) {
	c.Logger.Info("GetAllNotifications at notificationRepository started")
	var data []models.AllNotificationResponse
	querry := `
	SELECT id,
	sender_id, message, created_at, 
	sender_name as username, post_id,read
	FROM notifications 
	WHERE user_id = ? 
	ORDER BY created_at DESC
	`
	err := c.DB.Raw(querry, userId).Scan(&data).Error
	if err != nil {
		c.Logger.Error("Error while reading notifications: ", err)
		return nil, err
	}
	c.Logger.Info("GetAllNotifications at notificationRepository completed successfully")
	return data, nil
}
