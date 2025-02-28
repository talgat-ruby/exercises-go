package service

import (
 "errors"

 "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/repository"
 "github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

type PostService struct {
 repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
 return &PostService{
  repo: repo,
 }
}

// CreatePost validates and creates a new post.
func (ps *PostService) CreatePost(post *models.Post) error {
 if post.Title == "" || post.Content == "" || post.Category == "" {
  return errors.New("all fields (title, content, category) are required")
 }
 return ps.repo.Create(post)
}

// GetAllPosts retrieves every post.
func (ps *PostService) GetAllPosts() ([]models.Post, error) {
 return ps.repo.GetAll()
}

// GetPostByID fetches a post using its unique ID.
func (ps *PostService) GetPostByID(id int64) (*models.Post, error) {
 p, err := ps.repo.GetPostByID(int(id))
 if err != nil {
  return nil, errors.New("post not found")
 }
 return p, nil
}

// UpdatePost validates and updates an existing post.
func (ps *PostService) UpdatePost(post *models.Post) error {
 if post.Title == "" || post.Content == "" || post.Category == "" {
  return errors.New("all fields (title, content, category) are required")
 }

 existingPost, err := ps.repo.GetPostByID(int(post.ID))
 if err != nil {
  return errors.New("post not found")
 }

 existingPost.Title = post.Title
 existingPost.Content = post.Content
 existingPost.Category = post.Category
 existingPost.Tags = post.Tags

 return ps.repo.Update(existingPost)
}

// DeletePost removes a post given its ID.
func (ps *PostService) DeletePost(id int64) error {
 _, err := ps.repo.GetPostByID(int(id))
 if err != nil {
  return errors.New("post not found")
 }
 return ps.repo.Delete(int(id))
}