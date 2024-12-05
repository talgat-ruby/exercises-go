package router

import (
	"context"
)

func (r *Router) server(ctx context.Context) {
	r.router.HandleFunc("GET /ping", r.handler.Ping)
}
