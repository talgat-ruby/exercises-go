package categories

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (c *Categories) GetAllPostsOfCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := c.logger.With("method", "GetAllPostsOfCategory")

	idString := r.URL.Path[len("/categories/"):]
	idString = idString[:len(idString)-6]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.ErrorContext(ctx, "invalid category ID", "error", err)
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}
	posts, err := c.db.GetAllPostsOfCategory(ctx, id)
	if err != nil {
		log.ErrorContext(ctx, "failed to get posts by category", "error", err)
		http.Error(w, fmt.Sprintf("Failed to get posts: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if len(posts) == 0 {
		log.InfoContext(ctx, "no posts found for category", "category_id", id)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if err := response.JSON(w, http.StatusOK, posts); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(
		ctx, "success get all posts of category",
		"category_id", id,
		"number of posts", len(posts))
}
