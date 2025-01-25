package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"tracker/internal/auth"
)

// type Middleware struct {
// 	log *slog.Logger
// }

// func New(log *slog.Logger) *Middleware {
// 	return &Middleware{
// 		log: log,
// 	}
// }

func AuthMiddleware(next http.Handler) http.Handler {
	a := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
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
		tokenSecret := os.Getenv("SECRET_TOKEN")
		fmt.Printf("this token secret: (%s)\n", tokenSecret)
		fmt.Printf("this acces token : (%s)\n", tokenStr)
		userData, err := auth.ParseToken(tokenStr, tokenSecret)
		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			log.Println("Unauthorized", "error", err)
			return
		}
		newCtx := context.WithValue(ctx, "user", userData)
		next.ServeHTTP(w, r.WithContext(newCtx))
	}
	return http.HandlerFunc(a)
}
