package helper

import (
	config "ExploriteGateway/pkg/config"
	"ExploriteGateway/pkg/utils/models"
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
)

type Helper struct {
	config *config.Config
	Groups map[string][]*websocket.Conn
}

func NewHelper(config *config.Config) *Helper {
	return &Helper{
		config: config,
		Groups: make(map[string][]*websocket.Conn),
	}
}

func (h *Helper) SendMessageToUser(User map[string]*websocket.Conn, msg []byte, userID string) {
	var message models.Message
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println("Error while unmarshaling:", err)
		return
	}

	message.SenderID = userID
	recipientConn, ok := User[message.RecipientID]
	log.Println("Recipient ID:", message.RecipientID)
	if ok {
		if err := recipientConn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Error sending message to recipient:", err)
		}
	}
	if err := KafkaProducer(message); err != nil {
		log.Println("Error sending message to Kafka:", err)
	}
	log.Println("Message sent successfully to user")
}

func KafkaProducer(message models.Message) error {
	log.Println("Sending message to Kafka:", message)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("Error loading config:", err)
		return err
	}
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		log.Println("Error creating Kafka producer:", err)
		return err
	}
	log.Println("Kafka producer created successfully")

	result, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling message:", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: cfg.KafkaTopic,
		Key:   sarama.StringEncoder("Friend message"),
		Value: sarama.StringEncoder(result),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println("Error sending message to Kafka:", err)
		return err
	}

	log.Printf("[producer] Partition ID: %d; Offset: %d; Value: %v\n", partition, offset, msg)
	log.Println("Message sent successfully to Kafka")
	return nil
}

func (h *Helper) ValidateToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("123456789"), nil
	})
	if err != nil {
		log.Println("Error parsing token:", err)
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["id"].(float64)
		if !ok {
			log.Println("User ID not found in token")
			return 0, errors.New("user_id not found in token")
		}
		return int(userID), nil
	}
	log.Println("Invalid token")
	return 0, errors.New("invalid token")
}

func (h *Helper) AddUserToGroup(groupID string, conn *websocket.Conn) {
	h.Groups[groupID] = append(h.Groups[groupID], conn)
	log.Println("User added to group:", groupID)
}

func (h *Helper) RemoveUserFromGroups(userID int, conn *websocket.Conn) {
	for groupID, conns := range h.Groups {
		for i, c := range conns {
			if c == conn {
				h.Groups[groupID] = append(h.Groups[groupID][:i], h.Groups[groupID][i+1:]...)
				log.Println("User removed from group:", groupID)
				break
			}
		}
	}
}

func (h *Helper) SendMessageToGroup(msg []byte, userID int, groupID string) {
	var message models.Message
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println("Error while unmarshaling:", err)
		return
	}

	message.SenderID = strconv.Itoa(userID)
	for _, conn := range h.Groups[groupID] {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Error sending message to group member:", err)
		}
	}
	if err := KafkaProducer(message); err != nil {
		log.Println("Error sending message to Kafka:", err)
	}
	log.Println("Message sent successfully to group:", groupID)
}
