package posts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (p *Posts) GetInformationOfPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := p.logger.With("method", "GetInformationOfPost")

	idString := r.URL.Path[len("/posts/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.ErrorContext(ctx, "invalid post ID", "error", err)
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := p.db.GetInformationOfPost(ctx, id)
	if err != nil {
		log.ErrorContext(ctx, "post not found", "post_id", id)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if err := response.JSON(w, http.StatusOK, post); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(
		ctx, "success get information of post",
		"post_id", id)
}
