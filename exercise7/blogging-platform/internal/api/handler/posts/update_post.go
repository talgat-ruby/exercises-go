package posts

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/request"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (p *Posts) UpdatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := p.logger.With("method", "UpdatePost")

	idString := r.URL.Path[len("/posts/"):]
	id_post, err := strconv.Atoi(idString)
	if err != nil {
		log.ErrorContext(ctx, "invalid post ID", "error", err)
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var req blog.PostRequest
	if err := request.JSON(w, r, &req); err != nil {
		log.ErrorContext(ctx, "invalid request body", "error", err)
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	post, err := p.db.UpdatePost(ctx, id_post, req)
	if err != nil {
		if err == sql.ErrNoRows {
			log.ErrorContext(ctx, "post not found", "post_id", id_post)
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}
		log.ErrorContext(ctx, "failed to update post", "error", err)
		http.Error(w, fmt.Sprintf("Failed to update post: %v", err), http.StatusInternalServerError)
		return
	}

	if err := response.JSON(w, http.StatusCreated, post); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(
		ctx,
		"success update post",
		"post", post.String(),
	)
}
