package posts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/request"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (h *Posts) GetNumberPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "GetPost")
	id, err := strconv.Atoi(r.PathValue("numberPost"))
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

	err = request.JSON(w, r, res)
	if err != nil {
		response.JSON(w, http.StatusConflict, "")
		return
	}
	fmt.Fprintln(w, "GETIDpost succes")
}
