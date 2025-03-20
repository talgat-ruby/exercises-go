package redis_board

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/my/repo/internal/types/pubsub"
)

func (l *LeaderBoardRedis) ListPlayers() (pubsub.ListPlayerResp, error) {
	fmt.Println("method", "ListPlayers:redis")
	ctx := context.Background()
	topPlayers, err := l.rdb.ZRevRangeWithScores(ctx, l.conf.Redis.RedisName, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	var players []pubsub.ItemPlayerResp
	for _, player := range topPlayers {
		member := player.Member.(string)
		parts := strings.SplitN(member, ":", 2)
		if len(parts) < 2 {
			slog.Error("Invalid player format")
			continue
		}
		playerID := parts[0]
		playerName := parts[1]

		players = append(players, &itemPlayersResp{
			ID:    playerID,
			Name:  playerName,
			Score: int(player.Score),
		})

	}
	if players == nil || len(players) == 0 {
		slog.Error("No players found")
		return nil, errors.New("no players found")
	}
	return &listPlayerResp{Players: players}, nil
}

type listPlayerResp struct {
	Players []pubsub.ItemPlayerResp
}

type itemPlayersResp struct {
	ID    string
	Name  string
	Score int
}

func (p *itemPlayersResp) GetPlayerID() string {
	return p.ID
}
func (p *itemPlayersResp) GetPlayerName() string {
	return p.Name
}
func (p *itemPlayersResp) GetPlayerScore() int {
	return p.Score
}

func (p *listPlayerResp) GetList() []pubsub.ItemPlayerResp {
	return p.Players
}
