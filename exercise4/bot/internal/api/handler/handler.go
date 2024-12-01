package handler

import (
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/model"
)

type Handler struct {
	game *model.Game
}

func New() *Handler {
	return &Handler{
		game: model.NewGame(),
	}
}
