package router

import "context"

func (r *Router) blogs(ctx context.Context) {
	r.router.HandleFunc("GET /blog/posts", r.handler.FindBlogPosts)
	r.router.HandleFunc("GET /blog/{id}", r.handler.FindBlogPost)
}
