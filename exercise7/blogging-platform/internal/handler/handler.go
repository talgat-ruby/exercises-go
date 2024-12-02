package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/service"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/request"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
)

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
	fmt.Fprintln(w, "GETposts succes")

}
func DeleteNumberPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete succes")

}
func PostNumberPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST succes")

}
func PutNumberPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PUt succes")

}
