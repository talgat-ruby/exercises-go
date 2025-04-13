package tags

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

func (t *Tags) GetAllPostsOfTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := t.logger.With("method", "GetAllPostsOfTag")

	idString := r.URL.Path[len("/tags/"):]
	idString = idString[:len(idString)-6]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.ErrorContext(ctx, "invalid tag ID", "error", err)
		http.Error(w, "Invalid tag ID", http.StatusBadRequest)
		return
	}
	posts, err := t.db.GetAllPostsOfTag(ctx, id)
	if err != nil {
		log.ErrorContext(ctx, "failed to get posts by tag", "error", err)
		http.Error(w, fmt.Sprintf("Failed to get posts: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if len(posts) == 0 {
		log.InfoContext(ctx, "no posts found for tag", "tag_id", id)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if err := response.JSON(w, http.StatusOK, posts); err != nil {
		log.ErrorContext(ctx, "failed to encode response", "error", err)
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	log.InfoContext(
		ctx, "success get all posts of tag",
		"tag_id", id,
		"number of posts", len(posts))
}
