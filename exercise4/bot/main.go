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

	name := os.Getenv("NAME")
	port := os.Getenv("PORT")
	url := "http://locolhost:" + port
	client.SendJoinRequest(ctx, name, url)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Wait for SIGINT or SIGTERM
}
