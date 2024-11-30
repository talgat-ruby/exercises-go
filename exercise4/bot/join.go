package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Player struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func NewPlayer(name, port string) *Player {
	return &Player{
		Name: name,
		URL:  "http://localhost:" + port,
	}
}

func JoinHandler() {
	player := NewPlayer(PlayerName, Port)

	path := "http://localhost:4444/join"
	jsonData, err := json.Marshal(player)
	if err != nil {
		log.Fatalf("Error during serialization: %v", err)
	}

	resp, err := http.Post(path, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error ping request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("Success join")
	} else {
		log.Println("Not success join", resp.StatusCode)
	}
}
