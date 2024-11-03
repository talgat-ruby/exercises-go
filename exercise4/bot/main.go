package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/player"
)

func main() {
	ctx := context.Background()

	ready := startServer()
	<-ready

	player.JoinToGame(ctx)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Wait for SIGINT or SIGTERM
}
