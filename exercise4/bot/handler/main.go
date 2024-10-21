package handler

import (
	"github.com/talgat-ruby/exercises-go/exercise4/bot/ticTacToe"
)

type Handler struct {
	Player *ticTacToe.Player
}

func New(name, url string) *Handler {
	player := ticTacToe.New(name, url)
	return &Handler{
		Player: player,
	}
}
