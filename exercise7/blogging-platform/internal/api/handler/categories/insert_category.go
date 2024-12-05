package categories

import (
	"fmt"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db/blog"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/request"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (c *Categories) InsertCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := c.logger.With("method", "InsertCategory")

	//w.Header().Set("Content-Type", "application/json")

	var cat blog.CategoryRequest
	if err := request.JSON(w, r, &cat); err != nil {
		log.ErrorContext(ctx, "invalid request body", "error", err)
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	category, err := c.db.InsertCategory(ctx, cat)
	if err != nil {
		log.ErrorContext(ctx, "failed to create category", "error", err)
		http.Error(w, fmt.Sprintf("Failed to create category: %v", err), http.StatusInternalServerError)
		return
	}

	if err := response.JSON(w, http.StatusCreated, category); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(
		ctx,
		"success create category",
		"category", category.String(),
	)

}
