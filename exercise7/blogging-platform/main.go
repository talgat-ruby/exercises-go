package main

import (
	"context"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api"
	conf "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/config"
	db "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	//config
	config, err := conf.NewConfig(ctx)
	if err != nil {
		slog.ErrorContext(
			ctx,
			"error on load config",
			"service", "config",
			"error", err,
		)
	}

	// db
	dbConnect, err := db.New(config, slog.With("service", "db"))
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
	a := api.NewApi(slog.With("service", "api"), dbConnect)
	slog.InfoContext(ctx, "initialize service", "service", "api")

	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		if err := a.Start(ctx); err != nil {
			slog.ErrorContext(ctx, "failed to start api", "error", err.Error())
		}
		cancelFunc()
	}(ctx, cancel)

	go func(cancelFunc context.CancelFunc) {
		shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
		signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

		sig := <-shutdown
		slog.WarnContext(ctx, "signal received - shutting down...", "signal", sig)

		cancelFunc()
	}(cancel)

	<-ctx.Done()
}
