package internal

import (
	"encoding/json"
	"net/http"

	"github.com/AbhiramiRajeev/Ingestion-Service/models"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Producer sarama.SyncProducer
	APIKey   string
}

func NewHandler(p sarama.SyncProducer, apiKey string) *Handler {
	return &Handler{
		Producer: p,
		APIKey:   apiKey,
	}
}

func (h *Handler) IngestEvent(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !ValidateAPIKey(authHeader, h.APIKey) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized"})
		return
	}

	var event models.Event
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body"})

		return
	}

	// jsonData , err := c.GetRawData()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Failed to read request body"})
	// }

	jsonData, err := json.Marshal(event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to marshal event"})
		return
	}
	msg := &sarama.ProducerMessage{
		Topic: "login_events",
		Value: sarama.ByteEncoder(jsonData),
	}

	_, _, err = h.Producer.SendMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send message to Kafka"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event ingested successfully"})
}
