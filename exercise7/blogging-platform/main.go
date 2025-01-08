package main

import (
	"log"
	"net/http"
	"blogging-platform/internal/posts" // Импорт обработчиков
	"blogging-platform/pkg/database"  // Импорт базы данных
)

func main() {
	// Подключаем базу данных
	database.Connect()

	// Настраиваем маршрут для /posts
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			posts.GetPosts(w, r)
		case http.MethodPost:
			posts.CreatePost(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	// Запускаем сервер
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
