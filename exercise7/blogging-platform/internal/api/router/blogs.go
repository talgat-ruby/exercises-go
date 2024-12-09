package router

import "context"

func (r *Router) blogs(ctx context.Context) {
	//r.router.HandleFunc("GET /movies", r.handler.FindMovies)
	r.router.HandleFunc("GET /blog/{id}", r.handler.FindBlogPost)
}
