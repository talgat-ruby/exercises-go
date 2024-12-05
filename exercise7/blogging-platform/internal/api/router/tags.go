package router

import (
	"context"
)

func (r *Router) tags(ctx context.Context) {
	r.router.HandleFunc("GET /tags", r.handler.GetAllTags)
	r.router.HandleFunc("GET /tags/{id}/posts", r.handler.GetAllPostsOfTag)
}
