package player

import (
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal"
)

func (p *Player) SetToken(token internal.Token) *Player {
	cp := *p
	cp.Token = &token
	return &cp
}
