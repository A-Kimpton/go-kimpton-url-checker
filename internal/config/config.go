package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	MessageQueueProvider string `env:"MESSAGE_QUEUE_PROVIDER" envDefault:"rabbitmq"`
	RabbitMQ             RabbitMQ
}

type RabbitMQ struct {
	Host     string `env:"RABBITMQ_HOST" envDefault:"localhost"`
	Port     int    `env:"RABBITMQ_PORT" envDefault:"5672"`
	User     string `env:"RABBITMQ_USER" envDefault:"guest"`
	Password string `env:"RABBITMQ_PASSWORD" envDefault:"guest"`
}

func NewConfig() (*Config, error) {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Parse environment variables into the Config struct
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
