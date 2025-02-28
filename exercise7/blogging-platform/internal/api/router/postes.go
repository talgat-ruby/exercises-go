package router

import "context"

func (r *Router) PostRouter(ctx context.Context) {
	r.router.HandleFunc("GET /posts", r.handler.PostsIndex)
	r.router.HandleFunc("GET /posts/{id}", r.handler.PostsShow)
	r.router.HandleFunc("POST /posts", r.handler.PostsCreate)
	r.router.HandleFunc("PUT /posts/{id}", r.handler.PostsUpdate)
	r.router.HandleFunc("DELETE /posts/{id}", r.handler.PostsDelete)
}
