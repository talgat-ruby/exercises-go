package router

import (
	"context"
)

func (r *Router) BloggingRouter(_ context.Context) {
	r.router.HandleFunc("GET /post/{numberPost}", r.handler.GetNumberPost)
	r.router.HandleFunc("GET /posts", r.handler.GetPosts)
	r.router.HandleFunc("DELETE /post/{numberPost}", r.handler.DeleteNumberPost)
	r.router.HandleFunc("POST /posts", r.handler.PostNumberPost)
	r.router.HandleFunc("PUT /post/{numberPost}", r.handler.PutNumberPost)
}
