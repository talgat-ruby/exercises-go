package router

import (
	"context"
	"github.com/my/repo/internal/rest/handler"
	"net/http"
)

type Router struct {
	router  *http.ServeMux
	handler *handler.Handler
}

func NewRouter(handler *handler.Handler) *Router {
	mux := http.NewServeMux()
	return &Router{
		router:  mux,
		handler: handler,
	}
}

func (r *Router) Start(ctx context.Context) *http.ServeMux {
	r.player(ctx)
	return r.router
}
