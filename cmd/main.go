package main

import (
	"log"
	"os"

	"github.com/AbhiramiRajeev/Ingestion-Service/internal"
	"github.com/gin-gonic/gin"
	"k8s.io/klog"
)

func main() {
	brokers := []string{"localhost:9092"}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API KEY is not set")
	}
	producer, err := internal.NewKafkaProducer(brokers)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	r := gin.Default()
	h := internal.NewHandler(producer, apiKey)

	r.POST("/ingest", h.IngestEvent)

	defer func() {
		if err := producer.Close(); err != nil {
			klog.Fatal("Unable to close kafka ", err)
		}
		klog.Info("kafka closed successfully")
	}()
	r.Run(":8080")

}
