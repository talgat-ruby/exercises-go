package game

import (
	"fmt"
	"strings"
)

func (g *Game) Status() {
	fmt.Printf(
		`
Game status: %s

Players:
%s

`,
		g.State,
		g.playersList(),
	)
}

func (g *Game) playersList() string {
	strs := make([]string, 0, len(g.Players))

	for i, player := range g.Players {
		str := fmt.Sprintf("%d %s %s", i+1, player.Name, player.URL)
		strs = append(strs, str)
	}

	return strings.Join(strs, "\n")
}
