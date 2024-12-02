package api

import (
	"context"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/router"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/service"
	"log"
	"net"
	"net/http"
)

type API struct {
	port string
	mux  *http.ServeMux
}

func New(cfg map[string]string, service *service.PostService) *API {
	return &API{
		port: cfg["API_PORT"],
		mux:  router.SetupRouter(service),
	}
}

func (a *API) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.port),
		Handler: a.mux,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		<-ctx.Done()
		_ = server.Shutdown(context.Background())
	}()

	log.Printf("API server is running on port %s", a.port)
	return server.ListenAndServe()
}
