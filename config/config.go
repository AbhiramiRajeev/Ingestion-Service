package config

type Config struct {
	Kafka struct {
		Brokers []string
		Topic   string
	}
	Server struct {
		Port string
	}
	Auth struct {
		ApiKey string
	}
}
