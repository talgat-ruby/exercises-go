package authhandler

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"tracker/internal/auth"
	authdb "tracker/internal/db/auth_db"
	"tracker/internal/models"
	"tracker/utils/request"
	"tracker/utils/respone"
)

type AuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type RegisterRequest struct {
	Data *models.AuthUser `json:"data"`
}
type RegisterResponse struct {
	Data *models.Tokens `json:"data"`
}
type AuthentificatorHandler struct {
	database authdb.AuthDB
	logger   *slog.Logger
}

func NewAuthHandler(authDB *authdb.AuthentificatorDB, slogger *slog.Logger) *AuthentificatorHandler {
	return &AuthentificatorHandler{
		logger:   slogger,
		database: authDB,
	}
}

func (h *AuthentificatorHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if h.logger == nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	log := h.logger.With("method", "Register")
	requestBody := &RegisterRequest{}
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
	hashPassWOrd, salt, err := auth.HashPassword(requestBody.Data.Password, os.Getenv("PEPPER"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed hash password",
			"error", err,
		)
		http.Error(w, "failed hash password", http.StatusInternalServerError)
		return
	}
	user := &models.DBModelUser{
		Email:        requestBody.Data.Email,
		PasswordHash: hashPassWOrd,
		Salt:         salt,
	}

	dbResp, err := h.database.DBRegister(ctx, user)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to register user",
			"error", err,
		)
		http.Error(w, "failed to register user", http.StatusInternalServerError)
		return
	}
	fmt.Printf("id:(%d)\n", dbResp.ID)
	fmt.Printf("email:(%s)\n", dbResp.Email)
	fmt.Printf("id:(%s)\n", os.Getenv("SECRET_TOKEN"))

	tokenpair, err := auth.GenerateToken(&models.UserData{
		ID:    dbResp.ID,
		Email: dbResp.Email,
	}, os.Getenv("SECRET_TOKEN"))
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to generate token",
			"error", err,
		)
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}
	respbody := &RegisterResponse{
		Data: &models.Tokens{
			AccessToken:  tokenpair.AccessToken,
			RefreshToken: tokenpair.RefreshToken,
		},
	}
	err = respone.ResponseJSON(w, respbody, http.StatusOK)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to response",
			"error", err,
		)
		http.Error(w, "failed to response", http.StatusInternalServerError)
		return
	}
	log.InfoContext(ctx, "user registered", "email", dbResp.Email)
}
