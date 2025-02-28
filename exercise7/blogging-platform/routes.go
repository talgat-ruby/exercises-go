package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /healthcheck", app.healthcheckHandler)

	router.HandleFunc("GET /posts", app.showPostsHandler)
	router.HandleFunc("POST /posts", app.createPostHandler)
	router.HandleFunc("GET /posts/{id}", app.showPostHandler)
	router.HandleFunc("PUT /posts/{id}", app.updatePostHandler)
	router.HandleFunc("DELETE /posts/{id}", app.deletePostHandler)
	return router
}
