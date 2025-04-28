package main

import (
	"fmt"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "pong")
	if err != nil {
		log.Println(err)
		return
	}
}
