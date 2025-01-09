package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"tracker/internal/db"
)

func StartServer(dbExpence *db.ExpencesDBSt) error {
	mux := http.NewServeMux()
	AuthorizationHandlers(mux, dbExpence)
	BasicHandlers(mux, *dbExpence)

	port := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	log.Printf("Server is running on port %s\n", port)
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	return server.ListenAndServe()
}
