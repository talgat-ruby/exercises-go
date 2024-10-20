package game

import (
	"context"
	"fmt"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
)

func (g *Game) AddPlayer(ctx context.Context, newPlayer *player.Player) error {
	if g.State != StatePending {
		return fmt.Errorf("game has already started")
	}

	if g.isPlayerUrlExists(ctx, newPlayer) {
		return fmt.Errorf("player with url already exists")
	}

	g.Players = append(g.Players, newPlayer)
	return nil
}

func (g *Game) isPlayerUrlExists(_ context.Context, player *player.Player) bool {
	for _, p := range g.Players {
		if p.URL == player.URL {
			return true
		}
	}

	return false
}
