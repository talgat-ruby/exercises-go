package auth

import (
	"fmt"
	"net/http"
	"os"

	"expense_tracker/internal/auth"
	dbAuth "expense_tracker/internal/db/auth"
	"expense_tracker/pkg/httputils/request"
	"expense_tracker/pkg/httputils/response"
)

type AccessTokenRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenRequest struct {
	Data *AccessTokenRefreshTokenRequest `json:"data"`
}

type AccessTokenResponse struct {
	Data *AuthTokenPair `json:"data"`
}

func (h *Auth) AccessToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "AccessToken")

	// request parse
	requestBody := &AccessTokenRequest{}

	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	userData, err := auth.ParseToken(requestBody.Data.RefreshToken, os.Getenv("TOKEN_SECRET"))
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

	// db request
	dbResp, err := h.db.AccessToken(ctx, &dbAuth.AccessTokenInput{UserID: userData.ID})
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, "failed to query from db", http.StatusInternalServerError)
		return
	}
	if dbResp == nil {
		log.ErrorContext(
			ctx,
			"row is empty",
		)
		http.Error(w, "row is empty", http.StatusInternalServerError)
		return
	}

	tokenPair, err := auth.GenerateTokenPair(
		&auth.UserData{
			ID:    fmt.Sprint(dbResp.ID),
			Email: dbResp.Email,
		},
		os.Getenv("TOKEN_SECRET"),
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid request %w", err), http.StatusBadRequest)
		return
	}

	// response
	respBody := &AccessTokenResponse{
		Data: &AuthTokenPair{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		},
	}
	if err := response.JSON(
		w,
		http.StatusOK,
		respBody,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}

	log.InfoContext(
		ctx,
		"success generate access token",
	)
	return
}
