package move

import (
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/board"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
)

type Move struct {
	Player *player.Player `json:"players"`
	Board  *board.Board   `json:"board"`
}

func New(p *player.Player, b *board.Board) *Move {
	return &Move{
		Player: p,
		Board:  b,
	}
}
