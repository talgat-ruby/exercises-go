package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type Api struct {
	router *http.ServeMux
}

func NewApi() *Api {
	mux := http.NewServeMux()

	return &Api{
		router: mux,
	}
}

func (a *Api) Start(ctx context.Context) error {
	a.router.HandleFunc("GET /post", GetPost)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	fmt.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a test post"))
}
