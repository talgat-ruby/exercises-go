package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/service"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/request"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

type CreatedPost struct {
	newPost models.Blog
}

func GetNumberPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("numberPost"))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "")
		return
	}
	err, res := service.ServiceGetNumberPost(id)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, "")
		return
	}
	err = request.JSON(w, r, res)
	if err != nil {
		response.JSON(w, http.StatusConflict, "")
		return
	}
	fmt.Fprintln(w, "GETIDpost succes")
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	res, err := service.ServiceGetPosts()
	if err != nil {
		response.JSON(w, http.StatusConflict, "")
		return
	}
	err = request.JSON(w, r, res)
	if err != nil {
		response.JSON(w, http.StatusConflict, "")
		return
	}

}
func DeleteNumberPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete succes")

}
func PostNumberPost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Blog
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		request.JSON(w, r, http.StatusBadRequest)
		return
	}
	// err := response.JSON(w, http.StatusCreated, &newPost)
	err = service.CreationPost(newPost)
	if err != nil {
		request.JSON(w, r, http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "POST succes")

}
func PutNumberPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PUt succes")

}
