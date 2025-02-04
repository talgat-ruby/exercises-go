package config

import (
	"context"
	"flag"
	"github.com/joho/godotenv"
)

type SharedConfig struct {
	Port int    `env:"PORT"`
	Host string `env:"HOST,default=localhost"`
}

type Config struct {
	DB *DBConfig
}

func NewConfig(ctx context.Context) (*Config, error) {
	config := &Config{}
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	if c, err := NewConfigDB(ctx); err != nil {
		return nil, err
	} else {
		config.DB = c
	}
	flag.Parse()
	return config, nil
}
