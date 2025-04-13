package categories

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/request"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (c *Categories) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := c.logger.With("method", "UpdateCategory")

	idString := r.URL.Path[len("/categories/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.ErrorContext(ctx, "invalid category ID", "error", err)
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var cat blog.CategoryRequest
	if err := request.JSON(w, r, &cat); err != nil {
		log.ErrorContext(ctx, "invalid request body", "error", err)
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	category, err := c.db.UpdateCategory(ctx, id, cat)
	if err != nil {
		if err == sql.ErrNoRows {
			log.ErrorContext(ctx, "category not found", "category_id", id)
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}
		log.ErrorContext(ctx, "failed to update category", "error", err)
		http.Error(w, fmt.Sprintf("Failed to update category: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if err := response.JSON(w, http.StatusCreated, category); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(ctx, "success update category", "category", category.String())
}
