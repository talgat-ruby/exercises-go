package router

import (
	"context"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/handler"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/api/middleware"
	"net/http"
)

type Router struct {
	router  *http.ServeMux
	handler *handler.Handler
	midd    *middleware.Middleware
}

func New(h *handler.Handler, m *middleware.Middleware) (*Router, error) {
	mux := http.NewServeMux()

	return &Router{
		router:  mux,
		handler: h,
		midd:    m,
	}, nil
}

func (r *Router) Start(ctx context.Context) *http.ServeMux {
	r.auth(ctx)
	r.blogs(ctx)

	return r.router
}
