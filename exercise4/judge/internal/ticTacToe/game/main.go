package game

import (
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/match"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
)

type State string

const (
	StatePending  State = "PENDING"
	StateRunning  State = "RUNNING"
	StateFinished State = "FINISHED"
)

type Game struct {
	State   State            `json:"state"`
	Players []*player.Player `json:"players"`
	Matches []*match.Match   `json:"matches"`
}

func New() *Game {
	return &Game{
		State:   StatePending,
		Players: []*player.Player{},
		Matches: []*match.Match{},
	}
}
