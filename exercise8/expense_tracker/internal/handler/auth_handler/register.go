package authhandler

import (
	"log/slog"
	"net/http"

	authdb "tracker/internal/db/auth_db"
	"tracker/internal/models"
	"tracker/utils/request"
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
type authHandler struct {
	database authdb.AuthDB
	logger   *slog.Logger
}

func NewAuthHandler(authDB authdb.AuthDB) AuthHandler {
	return &authHandler{
		database: authDB,
	}
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
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
	
}
