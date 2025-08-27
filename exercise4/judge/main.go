package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/api"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	port := os.Getenv("PORT")

	a := api.New()

	if err := a.Start(ctx, port); err != nil {
		slog.ErrorContext(ctx, "api start error", "error", err)
		os.Exit(1)
	}

	go func() {
		shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
		signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

		sig := <-shutdown
		slog.WarnContext(ctx, "signal received - shutting down...", "signal", sig)

		cancel()
	}()

	<-ctx.Done()
	// Your cleanup tasks go here
	slog.InfoContext(ctx, "cleaning up ...")

	// Close api service
	if err := a.Stop(ctx); err != nil {
		slog.ErrorContext(ctx, "service stop error", "error", err)
	}

	slog.InfoContext(ctx, "server was successfully shutdown.")
}
