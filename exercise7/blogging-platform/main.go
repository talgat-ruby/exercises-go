package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	_ = godotenv.Load()

	d, err := db.New(slog.With("service", "db"))
	if err != nil {
		slog.ErrorContext(
			ctx,
			"initialize service error",
			"service", "db",
			"error", err,
		)
		panic(err)
	}

	if err := d.Ping(ctx); err != nil {
		panic(err)
	}

	if err := d.Init(ctx); err != nil {
		panic(err)
	}

	a := api.New(slog.With("service", "api"), d)
	if err := a.Start(ctx); err != nil {
		slog.ErrorContext(
			ctx,
			"initialize service error",
			"service", "api",
			"error", err,
		)
		panic(err)
	}

	go func() {
		shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
		signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

		sig := <-shutdown
		slog.WarnContext(ctx, "signal received - shutting down...", "signal", sig)

		cancel()
	}()

	if err := a.Stop(ctx); err != nil {
		slog.ErrorContext(ctx, "service stop error", "error", err)
	}

	slog.InfoContext(ctx, "server was successfully shutdown.")
}
