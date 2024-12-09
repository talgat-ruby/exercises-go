package blog

import "time"

type BlogPost struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Category  string    `json:"category" db:"category"`
	Tags      []string  `json:"tags" db:"tags"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

const BlogPostSchema = `
CREATE TABLE IF NOT EXISTS blog_posts (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    category VARCHAR(100) NOT NULL,
    tags TEXT[], -- PostgreSQL array type for storing string arrays
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Index for faster querying by category
CREATE INDEX IF NOT EXISTS idx_blog_posts_category ON blog_posts(category);

-- Index for timestamp-based queries
CREATE INDEX IF NOT EXISTS idx_blog_posts_created_at ON blog_posts(created_at);
`
