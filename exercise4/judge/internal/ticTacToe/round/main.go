package round

import (
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/board"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/move"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
)

type Round struct {
	Players [2]*player.Player `json:"players"`
	Board   *board.Board      `json:"board"`
	Moves   []*move.Move      `json:"moves"`
	Winner  *player.Player    `json:"winner"`
}

func New(players [2]*player.Player) *Round {
	return &Round{
		Players: players,
		Board:   board.New(),
		Moves:   []*move.Move{},
		Winner:  nil,
	}
}
