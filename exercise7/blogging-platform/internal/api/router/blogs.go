package router

import (
	"context"
)

func (r *Router) blogs(ctx context.Context) {
	r.router.HandleFunc("GET /blogs", r.handler.GetBlogs)
	r.router.HandleFunc("GET /blog/{id}", r.handler.GetBlogById)
	r.router.HandleFunc("POST /blog", r.handler.AddBlog)
	r.router.HandleFunc("PUT /blog/{id}", r.handler.UpdateBlog)
	r.router.HandleFunc("DELETE /blog/{id}", r.handler.DeleteBlog)
}
