package posts

import (
	"fmt"
	"net/http"
	"strconv"
)

func (p *Posts) DeletePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := p.logger.With("method", "DeletePost")

	idString := r.URL.Path[len("/posts/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.ErrorContext(ctx, "invalid post ID", "error", err)
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	err = p.db.DeletePost(ctx, id)
	if err != nil {
		if err.Error() == fmt.Sprintf("post with id %d not found", id) {
			log.ErrorContext(ctx, "post not found", "category_id", id)
			http.Error(w, "404 Not Found", http.StatusNotFound)
		} else {
			log.ErrorContext(ctx, "failed to delete post", "error", err)
			http.Error(w, fmt.Sprintf("Failed to delete post: %s", err.Error()), http.StatusInternalServerError)
		}
		return
	}

	log.InfoContext(ctx, "success delete post", "post_id", id)
	w.WriteHeader(http.StatusNoContent)
}
