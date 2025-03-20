package board

import (
	"slices"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal"
)

const (
	Cols = 3
	Rows = 3
)

type Board [Cols * Rows]internal.Token

func New() *Board {
	board := &Board{}
	slc := slices.Repeat([]internal.Token{internal.TokenEmpty}, Cols*Rows)
	copy(board[:], slc)
	return board
}
