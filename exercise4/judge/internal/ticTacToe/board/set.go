package board

import (
	"fmt"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal"
)

func (b *Board) Set(i int, t internal.Token) error {
	if b[i] != internal.TokenEmpty {
		return fmt.Errorf("selected cell(#%d) is already set to %v ", i, b[i])
	}

	b[i] = t
	return nil
}
