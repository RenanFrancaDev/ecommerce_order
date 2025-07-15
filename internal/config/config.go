package config

import "os"

type Config struct {
    RabbitMQURL string
}
func Load() *Config {
	
    return &Config{
        RabbitMQURL: os.Getenv("RABBITMQ_URL"),
    }
}