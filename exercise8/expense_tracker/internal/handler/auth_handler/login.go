package authhandler

import (
	"net/http"
	"os"
	"tracker/internal/auth"
	"tracker/internal/models"
	"tracker/utils/request"
	"tracker/utils/respone"
)

func (h *AuthentificatorHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "Login")
	requestBody := &models.LoginRequest{}
	err := request.RequestJSON(w, r, requestBody)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed parse request body",
			"error", err,
		)
		http.Error(w, "failed parse request body", http.StatusBadRequest)
		return
	}
	if requestBody.Data == nil {
		log.ErrorContext(ctx, "request body is empty")
		http.Error(w, "request body is empty", http.StatusBadRequest)
		return
	}
	if requestBody.Data.Email == "" || requestBody.Data.Password == "" {
		log.ErrorContext(ctx, "email or password is empty")
		http.Error(w, "email or password is empty", http.StatusBadRequest)
		return
	}
	user, err := h.database.DbLogin(requestBody.Data.Email, ctx)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to get user",
			"error", err,
		)
		http.Error(w, "failed to get user", http.StatusUnauthorized)
		return
	}
	if user == nil {
		log.ErrorContext(ctx, "user not found")
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}
	valid, err := auth.VerifyPassword(requestBody.Data.Password, user.Salt, os.Getenv("PEPPER"), user.PasswordHash)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to verify password",
			"error", err,
		)
		http.Error(w, "failed to verify password", http.StatusUnauthorized)
		return
	}
	if !valid {
		log.ErrorContext(ctx, "password is invalid")
		http.Error(w, "password is invalid", http.StatusUnauthorized)
		return
	}
	tokens, err := auth.GenerateToken(
		&models.UserData{
			ID:    user.ID,
			Email: user.Email,
		},
		os.Getenv("SECRET"),
	)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to generate token",
			"error", err,
		)
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}
	respBody := &models.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	err = respone.ResponseJSON(w, respBody, http.StatusOK)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to response",
			"error", err,
		)
		http.Error(w, "failed to response", http.StatusInternalServerError)
		return
	}
	log.InfoContext(ctx, "login success")

}
