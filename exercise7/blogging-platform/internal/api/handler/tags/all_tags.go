package tags

import (
	"fmt"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (t *Tags) GetAllTags(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := t.logger.With("method", "GetAllTags")

	tags, err := t.db.GetAllTags(ctx)
	if err != nil {
		log.ErrorContext(ctx, "failed to get tags", "error", err)
		http.Error(w, fmt.Sprintf("Failed to get tags: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if len(tags) == 0 {
		log.InfoContext(ctx, "no tags found")
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if err := response.JSON(w, http.StatusOK, tags); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(ctx, "success get all tags", "number of tags", len(tags))

}
