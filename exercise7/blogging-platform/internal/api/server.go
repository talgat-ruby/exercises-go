package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/handler"
)

type API struct {
	server *http.Server
}

func New() *API {
	return &API{}
}
func (a *API) Start(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /post/{numberPost}", handler.GetNumberPost)
	mux.HandleFunc("GET /posts", handler.GetPosts)
	mux.HandleFunc("DELETE /post/{numberPost}", handler.DeleteNumberPost)
	mux.HandleFunc("POST /posts", handler.PostNumberPost)
	mux.HandleFunc("PUT /post/{numberPost}", handler.PutNumberPost)
	port := os.Getenv("PORT")
	newPort := fmt.Sprintf("%s:", port)
	a.server = &http.Server{
		Addr:    newPort,
		Handler: mux,
	}
	go func() {
		<-ctx.Done()
		if err := a.server.Shutdown(context.Background()); err != nil {
			fmt.Println("Error shutting down server:", err)
		}
	}()

	err := http.ListenAndServe(newPort, nil)
	if err != nil {
		return err
	}
	return nil
}
