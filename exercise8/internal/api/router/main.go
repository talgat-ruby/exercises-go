package router

import (
	"context"
	"expense_tracker/internal/api/handler"
	"expense_tracker/internal/api/middleware"
	"log/slog"
	"net/http"
)

type Router struct {
	logger  *slog.Logger
	router  *http.ServeMux
	handler *handler.Handler
	mid     *middleware.Middleware
}

func New(logger *slog.Logger, handler *handler.Handler, mid *middleware.Middleware) *Router {
	mux := http.NewServeMux()

	return &Router{
		logger:  logger,
		router:  mux,
		handler: handler,
		mid:     mid,
	}
}

func (r *Router) Start(ctx context.Context) *http.ServeMux {
	r.expense(ctx)
	r.auth(ctx)
	return r.router
}
