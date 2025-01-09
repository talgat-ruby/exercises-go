package posts

import (
	"net/http"
	"strconv"

	"github.com/webaiz/exercise7/blogging-platform/internal/db/post"
	"github.com/webaiz/exercise7/blogging-platform/pkg/httputils/request"
	"github.com/webaiz/exercise7/blogging-platform/pkg/httputils/response"
)

type UpdatePostRequest struct {
	Data *post.ModelPost `json:"data"`
}

func (h *Posts) UpdatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "UpdatePost")

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

	// request parse
	requestBody := &UpdatePostRequest{}

	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	// db request
	if err := h.db.UpdatePost(ctx, int64(id), requestBody.Data); err != nil {
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
		"success update post",
		"id", id,
	)
	return
}
