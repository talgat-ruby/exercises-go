package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//ctx := context.Background()

	url := "http://localhost:" + os.Getenv("PORT")
	gamehost := "http://localhost:4444"
	name := os.Getenv("NAME")
	ready := startServer()
	<-ready

	// TODO after server start

	requestURL := fmt.Sprintf(gamehost + "/status")
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: status code: %d\n", res.StatusCode)
	jsonBody := []byte(`{"name": "` + name + `", "url": "` + url + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL = fmt.Sprintf(gamehost + "/join")

	res, err = http.Post(requestURL, "application/json", bodyReader)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: status code: %d\n", res.StatusCode)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Wait for SIGINT or SIGTERM
}
