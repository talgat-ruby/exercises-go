package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ready := startServer()
	<-ready

	// Присоединяемся к игре
	err := joinGame()
	if err != nil {
		log.Fatalf("Failed to join the game: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Ожидание завершения по SIGINT или SIGTERM
}
