package posts

import (
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/request"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

type CreatePostRequest struct {
	Data *models.Blog `json:"data"`
}

type CreatePostResponse struct {
	Data *models.Blog `json:"data"`
}

func (h *Posts) PostNumberPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "CreatePost")
	requestBody := &CreatePostRequest{}

	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	dbresp, err := h.db.NewPostCreator(ctx, *requestBody.Data)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, "failed to query from db", http.StatusBadRequest)
		return
	}
	if dbresp == nil {
		log.ErrorContext(
			ctx,
			"row is empty",
			"error", err,
		)
		http.Error(w, "row is empty", http.StatusBadRequest)
		return
	}
	resp := CreatePostResponse{
		Data: dbresp,
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
	// err = json.NewDecoder(r.Body).Decode(&newPost)
	// if err != nil {
	// 	response.JSON(w, http.StatusBadRequest, "")
	// 	return
	// }
	// err = posts.NewPostCreator(newPost)
	// if err != nil {
	// 	response.JSON(w, http.StatusBadRequest, "")
	// 	return
	// }
	// fmt.Fprintln(w, "POST succes")
	log.InfoContext(
		ctx,
		"succes insert post",
		"post id", resp.Data.Id,
	)
}



// func CreationPost(body models.Blog) error {
// 	if err := db.NewPostCreator(body); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func ServiceGetPosts() ([]models.Blog, error) {
// 	return db.DBGetPosts()
// }
