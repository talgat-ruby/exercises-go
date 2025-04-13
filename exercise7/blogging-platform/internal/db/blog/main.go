package blog

import (
	"encoding/json"
	"fmt"
	"time"
)

type Category struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Posts     []*Post    `json:"posts,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type CategoryRequest struct {
	Name string `json:"name"`
}

func (c *Category) MarshalJSON() ([]byte, error) {
	type Alias Category
	return json.Marshal(&struct {
		*Alias
		CreatedAt *time.Time `json:"created_at,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}{
		Alias:     (*Alias)(c),
		CreatedAt: filterNullTime(c.CreatedAt),
		UpdatedAt: filterNullTime(c.UpdatedAt),
	})
}

func filterNullTime(t *time.Time) *time.Time {
	if t == nil || t.IsZero() {
		return nil
	}
	return t
}

type Tag struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Posts     []*Post    `json:"posts,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (t *Tag) MarshalJSON() ([]byte, error) {
	type Alias Tag
	return json.Marshal(&struct {
		*Alias
		CreatedAt *time.Time `json:"created_at,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}{
		Alias:     (*Alias)(t),
		CreatedAt: filterNullTime(t.CreatedAt),
		UpdatedAt: filterNullTime(t.UpdatedAt),
	})
}

type Post struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Category  *Category  `json:"category"`
	Tags      []*Tag     `json:"tags,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (p *Post) MarshalJSON() ([]byte, error) {
	type Alias Post
	return json.Marshal(&struct {
		*Alias
		CreatedAt *time.Time `json:"created_at,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}{
		Alias:     (*Alias)(p),
		CreatedAt: filterNullTime(p.CreatedAt),
		UpdatedAt: filterNullTime(p.UpdatedAt),
	})
}

type PostRequest struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

// для правильного вывода логов
func (c *Category) String() string {
	return fmt.Sprintf("Category: {ID: %d, Name: %s }", c.ID, c.Name)
}

func (t *Tag) String() string {
	return fmt.Sprintf("Tag: {ID: %d, Name: %s}", t.ID, t.Name)
}

func (p *Post) String() string {
	return fmt.Sprintf("Post:{ID: %d, Title: %s, Content: %s, %s, Tags: %v}", p.ID, p.Title, p.Content, p.Category.String(), p.Tags)
}
