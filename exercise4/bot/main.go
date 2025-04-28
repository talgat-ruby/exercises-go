package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	PlayerName string
	Port       string
)

func main() {
	Port = os.Getenv("PORT")
	PlayerName = os.Getenv("NAME")

	if Port == "" {
		Port = "8080"
	}

	if PlayerName == "" {
		PlayerName = "PLayer:" + Port
	}

	ready := startServer()
	<-ready

	go JoinHandler()

	http.HandleFunc("GET /ping", ping)
	http.HandleFunc("POST /move", MoveBoard)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Wait for SIGINT or SIGTERM
}
