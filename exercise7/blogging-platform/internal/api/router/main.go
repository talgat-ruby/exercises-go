package router

import (
	"context"
	"net/http"

	"github.com/UAssylbek/blogging-platform/internal/api/handler"
)

type Router struct {
	router  *http.ServeMux
	handler *handler.Handler
}

func New(handler *handler.Handler) *Router {
	mux := http.NewServeMux()

	return &Router{
		router:  mux,
		handler: handler,
	}
}

func (r *Router) Start(ctx context.Context) *http.ServeMux {
	r.posts(ctx)

	return r.router
}
