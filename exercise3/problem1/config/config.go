package config

import (
	"log/slog"

	"github.com/talgat-ruby/exercises-go/exercise3/problem1/config/database"
)

type Config struct {
	Database *database.Config
	Port     string
}

func NewConfig(envPath string) *Config {
	slog.Info("hello from New Config")
	initConf(envPath)
	return &Config{
		Database: &database.Config{
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", "root"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 3306),
			DbName:   getEnv("DB_NAME", ""),
			SSLMode:  getEnv("DB_SSL_MODE", ""),
		},
		Port: getEnv("PORT", "8089"),
	}
}
