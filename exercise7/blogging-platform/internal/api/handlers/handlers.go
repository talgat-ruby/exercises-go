package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/model"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/service"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/pkg/httputils/response"
	"net/http"
	"strconv"
)

type Handler struct {
	service *service.PostService
}

func New(service *service.PostService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	var posts []model.Post
	var err error
	if term != "" {
		posts, err = h.service.GetPost(term)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if posts == nil || len(posts) == 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
	} else {
		posts, err = h.service.GetPosts()
		if err != nil {
			http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
			return
		}
	}

	if errJson := response.JSON(w, http.StatusOK, posts); errJson != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
	}
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdPost, err := h.service.CreatePost(&post)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdPost); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.URL.Path[len("/posts/"):]
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	post.ID = postID
	updatedPost, err := h.service.UpdatePost(&post)
	if err != nil {
		if err.Error() == "post not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedPost); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.URL.Path[len("/posts/"):]
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeletePost(postID); err != nil {
		if err.Error() == "post not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Post with ID %d deleted successfully", postID)
}
