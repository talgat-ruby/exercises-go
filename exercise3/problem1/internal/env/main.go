package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	AppPort string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := &Env{
		AppPort: os.Getenv("PORT"),
		DBHost:  os.Getenv("DB_HOST"),
		DBPort:  os.Getenv("DB_PORT"),
		DBUser:  os.Getenv("DB_USER"),
		DBPass:  os.Getenv("DB_PASSWORD"),
		DBName:  os.Getenv("DB_NAME"),
	}

	return env
}
