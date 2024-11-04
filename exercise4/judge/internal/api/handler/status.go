package handler

import (
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/game"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/pkg/httputils/response"
)

type ResponseStatus struct {
	Data *game.Game `json:"data"`
}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := slog.With(
		"handler", "Status",
		"path", r.URL.Path,
	)

	h.game.Status()

	if err := response.JSON(
		w,
		http.StatusOK,
		ResponseStatus{
			Data: h.game,
		},
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}
}
