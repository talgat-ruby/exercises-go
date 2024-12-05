package posts

import (
	"fmt"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (p *Posts) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := p.logger.With("method", "GetAllPosts")
	perm := r.URL.Query().Get("term")
	var posts []blog.Post
	var err error

	if perm != "" {

		log := p.logger.With("method", "SearchPosts")

		posts, err = p.db.SearchPosts(ctx, perm)
		if err != nil {
			log.ErrorContext(ctx, "failed to search posts", "error", err)
			http.Error(w, fmt.Sprintf("Failed to search posts: %s", err.Error()), http.StatusInternalServerError)
			return
		}

	} else {

		posts, err = p.db.GetAllPosts(ctx)
		if err != nil {
			log.ErrorContext(ctx, "failed to get posts", "error", err)
			http.Error(w, fmt.Sprintf("Failed to get posts: %s", err.Error()), http.StatusInternalServerError)
			return
		}

	}

	if len(posts) == 0 {
		log.InfoContext(ctx, "no posts found")
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if err := response.JSON(w, http.StatusOK, posts); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(
		ctx, "success get all posts",
		"number of posts", len(posts))
}
