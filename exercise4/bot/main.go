package main

import (
	"context"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/client"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	ready := startServer()
	<-ready

	client.SendJoinRequest(ctx)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Wait for SIGINT or SIGTERM
}
