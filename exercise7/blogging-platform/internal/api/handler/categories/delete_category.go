package categories

import (
	"fmt"
	"net/http"
	"strconv"
)

func (c *Categories) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := c.logger.With("method", "DeleteCategory")

	idString := r.URL.Path[len("/categories/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.ErrorContext(ctx, "invalid category ID", "error", err)
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	err = c.db.DeleteCategory(ctx, id)
	if err != nil {
		if err.Error() == fmt.Sprintf("category with id %d not found", id) {
			log.ErrorContext(ctx, "category not found", "category_id", id)
			http.Error(w, "404 Not Found", http.StatusNotFound)
		} else {
			log.ErrorContext(ctx, "failed to delete category", "error", err)
			http.Error(w, fmt.Sprintf("Failed to delete category: %s", err.Error()), http.StatusInternalServerError)
		}
		return
	}

	log.InfoContext(ctx, "success delete category", "category_id", id)
	w.WriteHeader(http.StatusNoContent)
}
