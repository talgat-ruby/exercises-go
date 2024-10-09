package game

import (
	"fmt"
	"sort"
	"strings"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
)

const (
	PointsWin  = 2
	PointsDraw = 1
	PointsLose = 0
)

type PlayerWithPoints struct {
	*player.Player
	points int
}

func (g *Game) leaderboard() string {
	pls := g.getAllPlayersWithPoints()

	strs := make([]string, 0, len(pls))
	for i, p := range pls {
		strs = append(strs, fmt.Sprintf("%d. %s %s %d", i+1, p.Name, p.URL, p.points))
	}

	return strings.Join(strs, "\n")
}

func (g *Game) getAllPlayersWithPoints() []*PlayerWithPoints {
	plsM := make(map[string]*PlayerWithPoints, len(g.Players))

	for _, p := range g.Players {
		pl := &PlayerWithPoints{p, 0}
		plsM[pl.URL] = pl
	}

	for _, m := range g.Matches {
		for _, r := range m.Rounds {
			if r.Winner == nil {
				plsM[r.Players[0].URL].points += PointsDraw
				plsM[r.Players[1].URL].points += PointsDraw
			}

			if r.Winner.URL == r.Players[0].URL {
				plsM[r.Players[0].URL].points += PointsWin
				plsM[r.Players[1].URL].points += PointsLose
			}

			plsM[r.Players[0].URL].points += PointsLose
			plsM[r.Players[1].URL].points += PointsWin
		}
	}

	pls := make([]*PlayerWithPoints, 0, len(plsM))
	for _, p := range plsM {
		pls = append(pls, p)
	}

	sort.Slice(
		pls, func(i, j int) bool {
			return pls[i].points > pls[j].points
		},
	)

	return pls
}
