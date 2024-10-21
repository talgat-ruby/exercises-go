package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type readyListener struct {
	net.Listener
	ready chan struct{}
	once  sync.Once
}

func (l *readyListener) Accept() (net.Conn, error) {
	l.once.Do(func() { close(l.ready) })
	return l.Listener.Accept()
}

type ResponseMove struct {
	Index int `json:"index"`
}

func servePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"message\":\"OK\"}")
}

type RequestMove struct {
	Board [9]string `json:"board"`
	Token string    `json:"token"`
}

func getIndex(b [9]string, me string, rival string) int {
	index := 0
	sum := 0
	diagCells := [4]int{0, 2, 6, 8}
	midCells := [4]int{1, 3, 5, 7}
	winCases := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	if b[4] == " " {
		return 4
	} else {
		for _, v := range winCases {
			sum = 0
			if b[v[0]] == me {
				sum++
			} else {
				if b[v[0]] == rival {
					continue
				}
				index = v[0]
			}
			if b[v[1]] == me {
				sum++
			} else {
				if b[v[1]] == rival {
					continue
				}
				index = v[1]
			}
			if b[v[2]] == me {
				sum++
			} else {
				if b[v[2]] == rival {
					continue
				}
				index = v[2]
			}

			if sum == 2 {
				if b[index] == " " {
					return index
				}
			}
		}

		for _, v := range winCases {
			sum = 0
			if b[v[0]] == rival {
				sum++
			} else {
				index = v[0]
			}
			if b[v[1]] == rival {
				sum++
			} else {
				index = v[1]
			}
			if b[v[2]] == rival {
				sum++
			} else {
				index = v[2]
			}

			if sum == 2 {
				if b[index] == " " {
					return index
				}
			}
		}

		for _, v := range diagCells {
			if b[v] == " " {
				return v
			}
		}

		for _, v := range midCells {
			if b[v] == " " {
				return v
			}
		}
	}

	return 0
}

func serveMove(w http.ResponseWriter, r *http.Request) {
	reqJson := RequestMove{}
	respJson := ResponseMove{}
	var me string
	var rival string
	if reqJson.Token == "x" {
		me = "x"
		rival = "o"
	} else {
		me = "o"
		rival = "x"
	}

	if err := json.NewDecoder(r.Body).Decode(&reqJson); err != nil {
		log.Fatal(err)
	}

	index := getIndex(reqJson.Board, me, rival)

	respJson = ResponseMove{
		Index: index,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(respJson)
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() <-chan struct{} {
	ready := make(chan struct{})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		panic(err)
	}

	list := &readyListener{Listener: listener, ready: ready}
	srv := &http.Server{
		IdleTimeout: 2 * time.Minute,
	}

	http.HandleFunc("GET /ping", servePing)
	http.HandleFunc("POST /move", serveMove)

	go func() {
		err := srv.Serve(list)
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	return ready
}
