package model

import "time"

type Post struct {
	ID        int       `json:"id"`        // ID is the primary key for the blog post
	Title     string    `json:"title"`     // Title is the title of the blog post
	Content   string    `json:"content"`   // Content is the main content of the blog post
	Category  string    `json:"category"`  // Category is the category under which the blog post falls
	Tags      []string  `json:"tags"`      // Tags are associated tags for the blog post
	CreatedAt time.Time `json:"createdAt"` // CreatedAt is the timestamp when the post was created
	UpdatedAt time.Time `json:"updatedAt"` // UpdatedAt is the timestamp when the post was last updated
}
