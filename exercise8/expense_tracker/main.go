package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {

}
func StartServer() error {
	mux := http.NewServeMux()
	mux.Handle("POST /expenses", http.HandlerFunc(ExpensesHandler))
	mux.Handle("GET /expenses", http.HandlerFunc(ExpensesHandlerGet))
	mux.Handle("GET /balance", http.HandlerFunc(BalanceHandler))
	log.Println("Server is running on port 8080")
	port := ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	return server.ListenAndServe()
}

func AuthMiddleware(next http.Handler) http.Handler {
	a := func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Authorization header is missing")
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer") {
			log.Println("Authorization header is invalid")
			http.Error(w, "Authorization header is invalid", http.StatusUnauthorized)
			return
		}
		tokenStr := authHeader[len("Bearer "):]

	}
	return http.HandlerFunc(a)
}
func ExpensesHandler(w http.ResponseWriter, r *http.Request) {

}
func ExpensesHandlerGet(w http.ResponseWriter, r *http.Request) {

}
func BalanceHandler(w http.ResponseWriter, r *http.Request) {

}

func ParseToken(token, secret string) (Tokens, error) {
	t, err := jwt.ParseToken(
		token,
	)
	return Tokens{}, nil
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
