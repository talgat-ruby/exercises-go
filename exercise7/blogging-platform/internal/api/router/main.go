package router

import (
	"blogging-platform/internal/api/handler"
	"context"
	"net/http"
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
	r.PostRouter(ctx)

	return r.router
}
