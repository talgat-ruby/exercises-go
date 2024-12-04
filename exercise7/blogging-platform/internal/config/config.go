package config

import "github.com/joho/godotenv"

type Config struct {
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	return config, nil
}
