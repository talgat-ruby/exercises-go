package router

import (
	"context"
)

func (r *Router) blogs(ctx context.Context) {
	r.router.HandleFunc("GET /blogs", r.handler.GetBlogs)
}
