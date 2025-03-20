package main

import (
	"log"
	"log/slog"
	"os"

	"tracker/internal/api"
	"tracker/internal/db"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("godotenverror:", err)
	}
	dbExpence, err := db.NewExpenceDB()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(os.Getenv("API_PORT"))
	err = api.StartServer(dbExpence)
	if err != nil {
		slog.Error("error start program")
		return
	}
}
