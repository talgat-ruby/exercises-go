package workout_tracker

import (
	"context"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())

	_ = godotenv.Load()

	slog.Info(os.Getenv("API_PORT")) // what better way to log information about api, should i store it in a file or send to differenct service inside my server ?

	//dbVar, err := db.New(slog.With
}
