package categories

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (c *Categories) GetInformationOfCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := c.logger.With("method", "GetInformationOfCategory")

	idString := r.URL.Path[len("/categories/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.ErrorContext(ctx, "invalid category ID", "error", err)
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	category, err := c.db.GetInformationOfCategory(ctx, id)
	if err != nil || category == nil {
		log.ErrorContext(ctx, "category not found", "category_id", id)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if err := response.JSON(w, http.StatusOK, category); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(
		ctx, "success get information of category",
		"category_id", id)
}
