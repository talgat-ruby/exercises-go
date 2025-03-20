package redis_board

import (
	"context"
	"fmt"

	"github.com/my/repo/internal/types/pubsub"

	"github.com/redis/go-redis/v9"

	"log/slog"
)

func (m *LeaderBoardRedis) CreatePlayer(req pubsub.CreatePlayerReq) (pubsub.CreatePlayerResp, error) {
	if req == nil {
		slog.Error("req is nil")
	}
	userId := req.GetPlayerID()
	PlayerName := req.GetPlayerName()
	member := fmt.Sprintf("%s:%s", userId, PlayerName)
	//mdl := LeaderBoard{
	//	UserID: userId,
	//	Name:   req.GetPlayerName(),
	//}
	ctx := context.Background()

	err := m.rdb.ZAdd(ctx, m.conf.Redis.RedisName, redis.Z{
		Score: float64(req.GetPlayerScore()), Member: member,
	}).Err()
	if err != nil {
		return nil, err
	}
	return true, nil
}
