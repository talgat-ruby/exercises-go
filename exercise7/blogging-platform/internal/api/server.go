package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/handler"
)

type API struct {
	router *http.ServeMux
}

func New() *API {
	mux := http.NewServeMux()
	return &API{
		router: mux,
	}
}
func (a *API) Start(ctx context.Context) error {
	a.BloggingRouter(ctx)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	decimalPort, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	stringPort := fmt.Sprintf(":%d", decimalPort)
	server := &http.Server{
		Addr:    stringPort,
		Handler: a.router,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}
	// go func() {
	// 	<-ctx.Done()
	// 	if err := server.Shutdown(context.Background()); err != nil {
	// 		fmt.Println("Error shutting down server:", err)
	// 	}
	// }()
	log.Printf("Start server: %d\n", decimalPort)
	err = server.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (a API) BloggingRouter(_ context.Context) {
	a.router.HandleFunc("GET /post/{numberPost}", handler.GetNumberPost)
	a.router.HandleFunc("GET /posts", handler.GetPosts)
	a.router.HandleFunc("DELETE /post/{numberPost}", handler.DeleteNumberPost)
	a.router.HandleFunc("POST /posts", handler.PostNumberPost)
	a.router.HandleFunc("PUT /post/{numberPost}", handler.PutNumberPost)
}
