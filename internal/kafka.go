package internal


import "github.com/IBM/sarama"


func NewKafkaProducer(brokers []string)(sarama.SyncProducer,error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true // By default, Sarama async producers don’t block waiting for an ack , make sure it returns after Kafka confirms it has received and stored the message.
    return sarama.NewSyncProducer(brokers, config)
}