package main

import (
	"context"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api"
	config2 "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/config"
	db "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
	"log/slog"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	//config
	config, err := config2.NewConfig()
	if err != nil {
		slog.ErrorContext(
			ctx,
			"error on load config",
			"service", "config",
			"error", err,
		)
	}

	// db
	dbConnect, err := db.New()
	if err != nil {
		slog.ErrorContext(
			ctx,
			"initialize service error",
			"service", "db",
			"error", err,
		)
		panic(err)
	}

	// api
	a := api.NewApi(config, dbConnect)
	slog.InfoContext(ctx, "initialize service", "service", "api")
	a.Start(ctx, cancel)
	<-ctx.Done()
}
