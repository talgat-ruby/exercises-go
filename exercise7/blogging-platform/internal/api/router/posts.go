package router

import (
	"context"
)

func (r *Router) BloggingRouter(_ context.Context) {
	r.router.HandleFunc("GET /post/{id}", r.handler.GetNumberPost)
	r.router.HandleFunc("GET /posts", r.handler.GetPosts)
	r.router.HandleFunc("DELETE /post/{id}", r.handler.DeleteNumberPost)
	r.router.HandleFunc("POST /post", r.handler.PostNumberPost)
	r.router.HandleFunc("PUT /post/{id}", r.handler.PutNumberPost)
}
