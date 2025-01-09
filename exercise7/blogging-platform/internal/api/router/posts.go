package router

import (
	"context"
	"net/http"
)

func (r *Router) posts(ctx context.Context) {
	r.router.Handle("GET /posts", http.HandlerFunc(r.handler.FindPosts))
	r.router.Handle("GET /posts/{id}", http.HandlerFunc(r.handler.FindPost))
	r.router.Handle("POST /posts", r.midd.Authenticator(http.HandlerFunc(r.handler.CreatePost)))
	r.router.Handle("PUT /posts/{id}", http.HandlerFunc(r.handler.UpdatePost))
	r.router.Handle("DELETE /posts/{id}", http.HandlerFunc(r.handler.DeletePost))
}
