package posts

import (
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

type GetPostResponse struct {
	Data *models.Blog `json:"data"`
}

func (h *Posts) GetNumberPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "GetPost")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to convert",
			"error", err,
		)
		http.Error(w, "failed to convert", http.StatusBadRequest)
		return
	}

	err, res := h.db.ServiceGetNumberPost(ctx, id)

	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to convert",
			"error", err,
		)
		return
	}
	resp := GetPostResponse{
		Data: res,
	}
	err = response.JSON(w, http.StatusOK, resp)
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
		"succes find posts",
		"id", resp.Data.Id,
	)
	return

}
