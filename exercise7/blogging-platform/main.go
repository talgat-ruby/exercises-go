package main

import (
	"context"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/service"
	"log/slog"
	"os"
	"os/signal"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/config"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.LoadConfig("config/app.properties")
	if err != nil {
		slog.ErrorContext(ctx, "failed to load configuration", "error", err)
		panic(err)
	}

	dbConn := db.InitDatabase(cfg, ctx)

	postService := service.NewPostService(dbConn)

	a := api.New(cfg, postService)
	if err := a.Start(ctx); err != nil {
		slog.ErrorContext(ctx, "initialize service error", "service", "api", "error", err)
		panic(err)
	}

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)

		sig := <-shutdown
		slog.WarnContext(ctx, "signal received - shutting down...", "signal", sig)

		cancel()
	}()
}
