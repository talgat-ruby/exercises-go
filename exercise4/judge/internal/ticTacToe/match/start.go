package match

import (
	"context"
	"fmt"
	"time"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/round"
)

var tokens = [2]internal.Token{internal.TokenX, internal.TokenO}

func (m *Match) Start(ctx context.Context, num int) {
	m.startPrint(num)

	for i := range m.Players {
		for j := range tokens {
			pl1 := m.Players[i%len(m.Players)].SetToken(tokens[j%len(tokens)])
			pl2 := m.Players[(i+1)%len(m.Players)].SetToken(tokens[(j+1)%len(tokens)])
			pls := [2]*player.Player{pl1, pl2}
			newRound := round.New(pls)
			m.Rounds = append(m.Rounds, newRound)
			newRound.Start(ctx, len(m.Rounds))
		}
	}

	m.endPrint(num)
}

func (m *Match) startPrint(num int) {
	fmt.Printf(
		`
Match #%d: %s %s vs %s %s
`,
		num,
		m.Players[0].Name,
		m.Players[0].URL,
		m.Players[1].Name,
		m.Players[1].URL,
	)
	<-time.After(time.Second)
}

func (m *Match) endPrint(num int) {
	fmt.Printf(
		`
End match #%d: %s %s %d:%d %s %s
`,
		num,
		m.Players[0].Name,
		m.Players[0].URL,
		m.totalRoundsWonBy(m.Players[0]),
		m.totalRoundsWonBy(m.Players[1]),
		m.Players[1].Name,
		m.Players[1].URL,
	)
	<-time.After(time.Second)
}

func (m *Match) totalRoundsWonBy(p *player.Player) int {
	total := 0

	for _, r := range m.Rounds {
		if r.Winner.URL == p.URL {
			total += 1
		}
	}

	return total
}
