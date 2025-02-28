package post

import (
	"github.com/lib/pq"
	"time"
)

type Model struct {
	ID        int            `json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Category  string         `json:"category"`
	Tags      pq.StringArray `json:"tags"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
}
