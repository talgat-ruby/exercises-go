package router

import (
	"context"
	"net/http"
)

func (r *Router) blogs(ctx context.Context) {
	r.router.Handle("GET /blogs", http.HandlerFunc(r.handler.GetBlogs))
	r.router.Handle("GET /blogs/{id}", http.HandlerFunc(r.handler.GetBlog))
	r.router.Handle("POST /blogs" /*r.midd.Authenticator*/, http.HandlerFunc(r.handler.CreateBlog))
	r.router.Handle("PUT /blogs/{id}" /*r.midd.Authenticator*/, http.HandlerFunc(r.handler.UpdateBlog))
	r.router.Handle("DELETE /blogs/{id}", http.HandlerFunc(r.handler.DeleteBlog))
}
