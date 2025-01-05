package main

import (
	"log"
	"log/slog"
	"os"

	"tracker/internal/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("gogdotenverror:", err)
	}
	log.Println(os.Getenv("API_PORT"))
	err = api.StartServer()
	if err != nil {
		slog.Error("error start program")
		return
	}
}
