package handler

import (
 "encoding/json"
 "net/http"
 "strconv"

 "github.com/gorilla/mux"
 "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/service"
 "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

type PostHandler struct {
 postService *service.PostService
}

func NewPostHandler(svc *service.PostService) *PostHandler {
 return &PostHandler{
  postService: svc,
 }
}

// CreatePost decodes the incoming request and creates a new post.
func (ph *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
 var blogPost models.Post
 if err := json.NewDecoder(r.Body).Decode(&blogPost); err != nil {
  http.Error(w, "Invalid request format", http.StatusBadRequest)
  return
 }

 if err := ph.postService.CreatePost(&blogPost); err != nil {
  http.Error(w, err.Error(), http.StatusBadRequest)
  return
 }

 w.WriteHeader(http.StatusCreated)
 if err := json.NewEncoder(w).Encode(blogPost); err != nil {
  http.Error(w, "Error sending response", http.StatusInternalServerError)
 }
}

// GetAllPosts fetches all posts and returns them in JSON.
func (ph *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
 posts, err := ph.postService.GetAllPosts()
 if err != nil {
  http.Error(w, "Error retrieving posts", http.StatusInternalServerError)
  return
 }

 w.Header().Set("Content-Type", "application/json")
 if err := json.NewEncoder(w).Encode(posts); err != nil {
  http.Error(w, "Error sending response", http.StatusInternalServerError)
 }
}

// GetPostByID retrieves a specific post using its ID.
func (ph *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
 params := mux.Vars(r)
 id, err := strconv.Atoi(params["id"])
 if err != nil {
  http.Error(w, "Invalid ID format", http.StatusBadRequest)
  return
 }

 post, err := ph.postService.GetPostByID(int64(id))
 if err != nil {
  http.Error(w, "Post not found", http.StatusNotFound)
  return
 }

 w.Header().Set("Content-Type", "application/json")
 if err := json.NewEncoder(w).Encode(post); err != nil {
  http.Error(w, "Error sending response", http.StatusInternalServerError)
 }
}

// UpdatePost decodes the update request and applies changes to an existing post.
func (ph *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
 params := mux.Vars(r)
 id, err := strconv.Atoi(params["id"])
 if err != nil {
  http.Error(w, "Invalid request format", http.StatusBadRequest)
  return
 }

 var updatedPost models.Post
 if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
  http.Error(w, "Invalid request format", http.StatusBadRequest)
  return
 }
 updatedPost.ID = int64(id)

 if err := ph.postService.UpdatePost(&updatedPost); err != nil {
  http.Error(w, err.Error(), http.StatusBadRequest)
  return
 }

 w.WriteHeader(http.StatusOK)
 w.Write([]byte("Post updated successfully"))
}

// DeletePost removes a post based on its ID.
func (ph *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
 params := mux.Vars(r)
 id, err := strconv.Atoi(params["id"])
 if err != nil {
  http.Error(w, "Invalid ID format", http.StatusBadRequest)
  return
 }

 if err := ph.postService.DeletePost(int64(id)); err != nil {
  http.Error(w, "Error deleting post", http.StatusInternalServerError)
  return
 }

 w.WriteHeader(http.StatusOK)
 w.Write([]byte("Post deleted successfully"))
}