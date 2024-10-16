package main

import (
	"context"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/api"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	ready := startServer()
	<-ready

	port := os.Getenv("PORT_API")
	a := api.New()

	if err := a.Start(ctx, port); err != nil {
		slog.ErrorContext(ctx, "bot start error", "error", err)
		os.Exit(1)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Wait for SIGINT or SIGTERM
}
