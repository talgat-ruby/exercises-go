package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Token string

const (
	Cols = 3
	Rows = 3
)

const (
	TokenEmpty Token = " "
	TokenX     Token = "x"
	TokenO     Token = "o"
)

type RequestMove struct {
	Board *Board `json:"board"`
	Token Token  `json:"token"`
}

func newRequestMove() *RequestMove {
	return &RequestMove{}
}

type Board [Cols * Rows]Token

type ResponseMove struct {
	Index int `json:"index"`
}

func newResponseMove() *ResponseMove {
	return &ResponseMove{}
}

func MoveBoard(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	requestMove := newRequestMove()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(r.Body)

	err = json.Unmarshal(body, &requestMove)
	if err != nil {
		fmt.Println("JSON unmarshalling error")
	}

	select {
	case <-ctx.Done():
		http.Error(w, "Execution time expired", http.StatusGatewayTimeout)
		return
	default:
		responseMove := newResponseMove()
		//smart bot loop
		for {
			i := rand.Intn(9)
			if requestMove.Board[i] == TokenEmpty {
				responseMove.Index = i
				break
			}
		}

		jsonData, err := json.Marshal(responseMove)
		if err != nil {
			http.Error(w, "JSON serialization error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonData)
		if err != nil {
			log.Println("Error writing response:", err)
			return
		}
	}
}
