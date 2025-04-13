package router

import (
	"context"
)

func (r *Router) categories(ctx context.Context) {
	r.router.HandleFunc("POST /categories", r.handler.InsertCategory)
	r.router.HandleFunc("DELETE /categories/{id}", r.handler.DeleteCategory)
	r.router.HandleFunc("PUT /categories/{id}", r.handler.UpdateCategory)

	r.router.HandleFunc("GET /categories", r.handler.GetAllCategories)
	r.router.HandleFunc("GET /categories/{id}", r.handler.GetInformationOfCategory)
	r.router.HandleFunc("GET /categories/{id}/posts", r.handler.GetAllPostsOfCategory)
}
