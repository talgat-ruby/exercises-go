package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/config"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/handler"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/repository"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/service"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	r := mux.NewRouter()

	r.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	r.HandleFunc("/posts", postHandler.GetAllPosts).Methods("GET")
	r.HandleFunc("/posts/{id:[0-9]+}", postHandler.GetPostByID).Methods("GET")
	r.HandleFunc("/posts/{id:[0-9]+}", postHandler.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id:[0-9]+}", postHandler.DeletePost).Methods("DELETE")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the blogging platform!"))
	})

	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
