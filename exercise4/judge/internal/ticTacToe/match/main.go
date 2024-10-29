package match

import (
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/round"
)

type Match struct {
	Players [2]*player.Player `json:"players"`
	Rounds  []*round.Round    `json:"rounds"`
}

func New(players [2]*player.Player) *Match {
	return &Match{
		Players: players,
		Rounds:  []*round.Round{},
	}
}
