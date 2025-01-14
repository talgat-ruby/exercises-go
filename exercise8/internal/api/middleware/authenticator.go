package middleware

import (
	"context"
	"expense_tracker/internal/auth"
	"net/http"
	"os"
	"strings"
)

func (m *Middleware) Authenticator(next http.Handler) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := m.logger.With("middleware", "Authorization")

		authenticatorHeader := r.Header.Get("Authorization")
		if authenticatorHeader == "" {
			log.ErrorContext(
				ctx,
				"authenticator header is empty",
			)
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		if !strings.HasPrefix(authenticatorHeader, "Bearer ") {
			log.ErrorContext(
				ctx,
				"invalid authenticator header",
			)
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		tokenString := authenticatorHeader[len("Bearer "):]
		userData, err := auth.ParseToken(tokenString, os.Getenv("TOKEN_SECRET"))
		if err != nil {
			log.ErrorContext(
				ctx,
				"fail authentication",
				"error", err,
			)
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		newCtx := context.WithValue(ctx, "user", userData)
		next.ServeHTTP(w, r.WithContext(newCtx))
	}
	return http.HandlerFunc(h)
}
