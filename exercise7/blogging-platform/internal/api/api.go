package api

import (
	"blogging-platform/internal/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// API - структура для работы с API
type API struct {
	server *http.Server
}

// BlogPost - структура поста в блоге
type BlogPost struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

// New создаёт новый API
func New() *API {
	mux := http.NewServeMux()

	// Проверка статуса сервера
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running!")
	})

	// Получение всех постов с фильтрацией
	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			term := r.URL.Query().Get("term")
			posts := getAllPosts(term)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(posts)
			return
		}

		// Создание поста
		if r.Method == http.MethodPost {
			var post BlogPost
			if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// Валидация данных
			if post.Title == "" || post.Content == "" || post.Category == "" {
				http.Error(w, "Title, Content и Category не могут быть пустыми", http.StatusBadRequest)
				return
			}

			id := createPost(post)
			post.ID = id

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(post)
			return
		}

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	// Обновление и удаление постов
	mux.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/posts/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodPut {
			var updatedPost BlogPost
			if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			if updatedPost.Title == "" || updatedPost.Content == "" || updatedPost.Category == "" {
				http.Error(w, "Title, Content и Category не могут быть пустыми", http.StatusBadRequest)
				return
			}

			updatePost(id, updatedPost)
			updatedPost.ID = id
			json.NewEncoder(w).Encode(updatedPost)
			return
		}

		if r.Method == http.MethodDelete {
			deletePost(id)
			w.WriteHeader(http.StatusNoContent)
			return
		}

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	return &API{
		server: &http.Server{
			Addr:    ":8081",
			Handler: mux,
		},
	}
}

// Start запускает сервер
func (a *API) Start(ctx context.Context) error {
	fmt.Println("🚀 Starting server on :8080...")
	go func() {
		if err := a.server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("❌ Server error:", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("⚠️ Shutting down server...")
	return a.server.Shutdown(context.Background())
}

// Функции для работы с БД

// getAllPosts получает все посты с фильтрацией
func getAllPosts(term string) []BlogPost {
	query := "SELECT id, title, content, category, tags FROM posts"
	var args []interface{}

	if term != "" {
		query += " WHERE LOWER(title) LIKE LOWER($1) OR LOWER(content) LIKE LOWER($1) OR LOWER(category) LIKE LOWER($1)"
		args = append(args, "%"+term+"%")
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		log.Fatalf("❌ Ошибка запроса постов: %v", err)
	}
	defer rows.Close()

	var posts []BlogPost
	for rows.Next() {
		var post BlogPost
		var tags []string
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, pq.Array(&tags)); err != nil {
			log.Fatalf("❌ Ошибка чтения поста: %v", err)
		}
		post.Tags = tags
		posts = append(posts, post)
	}
	return posts
}

// createPost создаёт новый пост
func createPost(post BlogPost) int {
	query := "INSERT INTO posts (title, content, category, tags) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := db.DB.QueryRow(query, post.Title, post.Content, post.Category, pq.Array(post.Tags)).Scan(&id)
	if err != nil {
		log.Fatalf("❌ Ошибка вставки поста: %v", err)
	}
	return id
}

// updatePost обновляет существующий пост
func updatePost(id int, post BlogPost) {
	query := "UPDATE posts SET title=$1, content=$2, category=$3, tags=$4 WHERE id=$5"
	_, err := db.DB.Exec(query, post.Title, post.Content, post.Category, pq.Array(post.Tags), id)
	if err != nil {
		log.Fatalf("❌ Ошибка обновления поста: %v", err)
	}
}

// deletePost удаляет пост
func deletePost(id int) {
	query := "DELETE FROM posts WHERE id=$1"
	_, err := db.DB.Exec(query, id)
	if err != nil {
		log.Fatalf("❌ Ошибка удаления поста: %v", err)
	}
}
