package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/pkg/httputils/request"
)

type RequestJoin struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ResponseJoin struct {
	Message string `json:"name"`
}

func (h *Handler) Join(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := slog.With(
		"handler", "Join",
		"path", r.URL.Path,
	)

	var reqBody RequestJoin
	if err := request.JSON(w, r, &reqBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse json body",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create a player
	p := player.New(reqBody.Name, reqBody.URL)

	// check if a player responses to a ping
	if err := p.Ping(ctx); err != nil {
		log.ErrorContext(
			ctx,
			"failed to ping player",
			"player.name", p.Name,
			"player.remote", p.URL,
			"error", err,
		)
		http.Error(
			w,
			fmt.Errorf("ping to url failed %w, check if player is running", err).Error(),
			http.StatusBadRequest,
		)
		return
	}

	// add a player to the game
	if err := h.game.AddPlayer(ctx, p); err != nil {
		log.ErrorContext(
			ctx,
			"player was not added",
			"player.name", p.Name,
			"player.url", p.URL,
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.InfoContext(
		ctx,
		"new player joined the game",
		"player.name", p.Name,
		"player.url", p.URL,
	)
	w.WriteHeader(http.StatusNoContent)
}
