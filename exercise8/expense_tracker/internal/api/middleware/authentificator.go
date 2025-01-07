package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"tracker/internal/auth"
)

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
		userData, err := auth.ParseToken(tokenStr, os.Getenv("TOKEN_SECRET"))
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
