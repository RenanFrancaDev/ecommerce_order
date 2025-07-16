package config

import "os"

type Config struct {
	RabbitMQURL         string
	RabbitMQOrdersQueue string
	GinMode             string
}

func Load() *Config {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug"
	}
	
	return &Config{
		RabbitMQURL:         os.Getenv("RABBITMQ_URL"),
		RabbitMQOrdersQueue: os.Getenv("ORDERS_QUEUE"),
		GinMode:             ginMode,
	}
}
