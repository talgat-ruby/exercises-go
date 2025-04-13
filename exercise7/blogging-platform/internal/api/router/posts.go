package router

import (
	"context"
)

func (r *Router) posts(ctx context.Context) {
	r.router.HandleFunc("POST /posts", r.handler.InsertPost)
	r.router.HandleFunc("DELETE /posts/{id}", r.handler.DeletePost)
	r.router.HandleFunc("PUT /posts/{id}", r.handler.UpdatePost)
	r.router.HandleFunc("GET /posts/{id}", r.handler.GetInformationOfPost)
	r.router.HandleFunc("GET /posts", r.handler.GetAllPosts)
}
