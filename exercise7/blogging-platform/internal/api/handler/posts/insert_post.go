package posts

import (
	"fmt"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/request"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (p *Posts) InsertPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := p.logger.With("method", "InsertPost")
	var req blog.PostRequest

	if err := request.JSON(w, r, &req); err != nil {
		log.ErrorContext(ctx, "invalid request body", "error", err)
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	post, err := p.db.InsertPost(ctx, req)
	if err != nil {
		log.ErrorContext(ctx, "failed to create post", "error", err)
		http.Error(w, fmt.Sprintf("Failed to create post: %v", err), http.StatusInternalServerError)
		return
	}

	if err := response.JSON(w, http.StatusCreated, post); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(
		ctx,
		"success create post",
		"post", post.String(),
	)
}
