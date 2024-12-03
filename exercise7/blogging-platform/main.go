package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
)

func main() {
	ctx := context.Background()

	_ = godotenv.Load()

	fmt.Println(os.Getenv("API_PORT"))

	d, err := db.New(slog.With("service", "db"))
	if err != nil {
		panic(err)
	}
	if err := d.Init(ctx); err != nil {
		panic(err)
	}

	a := api.New(slog.With("service", "api"), d)
	if err := a.Start(ctx); err != nil {
		panic(err)
	}
}
