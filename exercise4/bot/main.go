package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	ready := startServer()
	<-ready

	// TODO after server start

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Wait for SIGINT or SIGTERM
}
