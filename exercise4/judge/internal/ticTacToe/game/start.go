package game

import (
	"context"
	"fmt"
	"time"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/match"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
)

func (g *Game) Start(ctx context.Context) error {
	if g.State != StatePending {
		return fmt.Errorf("game state is %s", g.State)
	}

	if len(g.Players) < 2 {
		return fmt.Errorf("game has no enough players, %d player(s)", len(g.Players))
	}

	g.startPrint()

	g.State = StateRunning

	go func(ctx context.Context, players []*player.Player) {
		for i := 0; i < len(players); i++ {
			for j := i + 1; j < len(players); j++ {
				pls := [2]*player.Player{players[i], players[j]}
				newMatch := match.New(pls)
				g.Matches = append(g.Matches, newMatch)
				newMatch.Start(ctx, len(g.Matches))
			}
		}
		g.endPrint()
	}(ctx, g.Players)

	return nil
}

func (g *Game) startPrint() {
	fmt.Printf(
		`
Starting Game!

Our players are:
%s
`,
		g.playersList(),
	)
	<-time.After(time.Second)
}

func (g *Game) endPrint() {
	fmt.Printf(
		`
Game Completed!

Leaderboard:
%s
`,
		g.leaderboard(),
	)
	<-time.After(time.Second)
}
