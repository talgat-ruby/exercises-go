package board

import (
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal"
)

// Evaluate returns winner token, if draw empty token, nill if game still continues
func (b *Board) Evaluate(currentInd int, currentT internal.Token) *internal.Token {
	isWin := false
	firstInRowInd := (currentInd / Cols) * Cols
	firstInColInd := currentInd % Cols

	isWin =
		//row
		b.assess(currentT, firstInRowInd, Cols, 1) ||
			// col
			b.assess(currentT, firstInColInd, Rows, Cols) ||
			// diagonal 1
			b.assess(currentT, 0, Cols, Rows+1) ||
			// diagonal 2
			b.assess(currentT, Rows-1, Cols, Rows-1)

	if isWin {
		return &currentT
	}

	if b.hasNoEmpty() {
		draw := internal.TokenEmpty
		return &draw
	}

	return nil
}

func (b *Board) assess(t internal.Token, firstInd, total, step int) bool {
	i := firstInd

	for range total {
		if b[i] != t {
			return false
		}
		i += step
	}

	return true
}

func (b *Board) hasNoEmpty() bool {
	for _, t := range b {
		if t == internal.TokenEmpty {
			return false
		}
	}

	return true
}
