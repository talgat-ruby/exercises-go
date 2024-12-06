package posts

import (
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/request"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

type UpdatePostRequest struct {
	Data *models.Blog `json:"data"`
}

func (h Posts) PutNumberPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With(
		"method", "UpdatePost",
	)
	id, err := strconv.Atoi(r.PathValue("numberPost"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to convert id to int",
			"error", err,
		)
	}
	requestBody := &UpdatePostRequest{}
	err = request.JSON(w, r, requestBody)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}
	err = h.db.UpdatePost(ctx, int64(id), *requestBody.Data)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed query from db",
			"error", err,
		)
		http.Error(w, "failed query from db", http.StatusInternalServerError)
		return
	}
	err = response.JSON(w, http.StatusNoContent, nil)
	if err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		http.Error(w, "fail json", http.StatusInternalServerError)
		return
	}
	log.InfoContext(
		ctx,
		"succes update post",
		"id", id,
	)
	return
}
