package posts

import (
    "blogging-platform/pkg/database"
    "encoding/json"
    "net/http"
)

type Post struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Content   string `json:"content"`
    CreatedAt string `json:"created_at"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT id, title, content, created_at FROM posts")
    if err != nil {
        http.Error(w, "Ошибка получения постов", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
            http.Error(w, "Ошибка сканирования данных", http.StatusInternalServerError)
            return
        }
        posts = append(posts, post)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(posts)
}
func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, "Неверный формат данных", http.StatusBadRequest)
        return
    }

    query := "INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING id"
    err := database.DB.QueryRow(query, post.Title, post.Content).Scan(&post.ID)
    if err != nil {
        http.Error(w, "Ошибка создания поста", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(post)
}
