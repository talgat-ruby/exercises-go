package router

import (
	"context"
)

func (r *Router) posts(ctx context.Context) {
	r.router.HandleFunc("GET /posts", r.handler.FindPosts)
	r.router.HandleFunc("GET /posts/{id}", r.handler.FindPost)
	r.router.HandleFunc("POST /posts", r.handler.CreatePost)
	r.router.HandleFunc("PUT /posts/{id}", r.handler.UpdatePost)
	r.router.HandleFunc("DELETE /posts/{id}", r.handler.DeletePost)
}
