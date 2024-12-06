package posts

import (
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (h *Posts) DeleteNumberPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With(
		"method", "DeletePost",
	)
	idString, err := strconv.Atoi(r.PathValue("numberPost"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"faildet convert id to int",
			"error", err,
		)
	}
	err = h.db.DeletePost(int64(idString), ctx)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from data base",
			"error", err,
		)
		http.Error(w, "failed to query from data base", http.StatusInternalServerError)
		return
	}
	err = response.JSON(w, http.StatusNoContent, nil)
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}
	log.InfoContext(
		ctx,
		"succes delete post",
		"id", idString,
	)
	return
}
