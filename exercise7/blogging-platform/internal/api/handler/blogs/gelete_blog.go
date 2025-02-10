package blogs

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
	"net/http"
	"strconv"
)

func (h *Blogs) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "DeleteBlog")

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to convert id to int",
			"error", err,
		)
		http.Error(w, "failed to convert id to int", http.StatusBadRequest)
		return
	}

	// db request
	if err := h.db.DeleteBlog(ctx, int64(id)); err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := response.JSON(
		w,
		http.StatusNoContent,
		nil,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}

	log.InfoContext(
		ctx,
		"success delete blog",
		"id", id,
	)
	return
}
