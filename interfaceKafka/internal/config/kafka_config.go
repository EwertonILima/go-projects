package config

import (
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/joho/godotenv"
)

// Config holds Kafka configuration details
type Config struct {
	BootstrapServers string `mapstructure:"bootstrap.servers"`
}

// LoadConfig reads Kafka configuration from environment variables
func LoadConfig() (*Config, error) {
	errLoad := godotenv.Load()
  if errLoad != nil {
    fmt.Print("Error loading .env file")
  }
	config := &Config{}
	err := mapstructure.Decode(map[string]string{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
	}, config)
	if err != nil {
		return nil, fmt.Errorf("error loading Kafka configuration: %w", err)
	}
	return config, nil
}
