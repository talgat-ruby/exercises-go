package main

import (
	"time"
)

// Todo представляет структуру задачи
type Todo struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TodoCreate представляет структуру для создания новой задачи
type TodoCreate struct {
	Description string `json:"description"`
}

// TodoUpdate представляет структуру для обновления существующей задачи
type TodoUpdate struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
