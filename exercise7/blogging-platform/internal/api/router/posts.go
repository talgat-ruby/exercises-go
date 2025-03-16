package router

import (
	"context"
)

func (r *Router) posts(ctx context.Context) {
	r.router.HandleFunc("GET /posts", r.handler.FindPosts)
}
