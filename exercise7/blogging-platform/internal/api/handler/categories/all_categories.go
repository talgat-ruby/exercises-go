package categories

import (
	"fmt"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (c *Categories) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := c.logger.With("method", "GetAllCategories")

	categories, err := c.db.GetAllCategories(ctx)
	if err != nil {
		log.ErrorContext(ctx, "failed to get categories", "error", err)
		http.Error(w, fmt.Sprintf("Failed to get categories: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if err := response.JSON(w, http.StatusOK, categories); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(ctx, "success get all categories", "number of categories", len(categories))

}
