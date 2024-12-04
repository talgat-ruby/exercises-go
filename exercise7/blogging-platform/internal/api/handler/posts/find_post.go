package posts

import (
	"net/http"
	"strconv"

	"github.com/UAssylbek/blogging-platform/internal/db/post"
	"github.com/UAssylbek/blogging-platform/pkg/httputils/response"
)

type FindPostResponse struct {
	Data *post.ModelPost `json:"data"`
}

func (h *Posts) FindPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "FindPost")

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

	dbResp, err := h.db.FindPost(ctx, int64(id))

	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := FindPostResponse{
		Data: dbResp,
	}

	if err := response.JSON(
		w,
		http.StatusOK,
		resp,
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
		"success find post",
		"post id", resp.Data.ID,
	)
	return
}
