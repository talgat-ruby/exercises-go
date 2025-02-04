package main

import (
	"context"
	"expense_tracker/internal/api"
	"expense_tracker/internal/db"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	_ = godotenv.Load()

	d, err := db.New(slog.With("service", "db"))
	if err != nil {
		panic(err)
	}

	a := api.New(slog.With("service", "api"), d)

	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		if err := a.Start(ctx); err != nil {
			slog.ErrorContext(ctx, "failed to start api", "error", err.Error())
		}
		cancelFunc()
	}(ctx, cancel)

	go func(cancelFunc context.CancelFunc) {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)

		sig := <-shutdown
		slog.WarnContext(ctx, "signal received - shutting down...", "signal", sig)

		cancelFunc()
	}(cancel)

	<-ctx.Done()
}
