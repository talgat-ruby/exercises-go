package redis_board

import (
	"github.com/my/repo/internal/config"
	"github.com/redis/go-redis/v9"
)

type LeaderBoardRedis struct {
	conf *config.Config
	rdb  *redis.Client
}

func NewLeaderBoardRedis(conf *config.Config, rdb *redis.Client) *LeaderBoardRedis {
	return &LeaderBoardRedis{
		conf: conf,
		rdb:  rdb,
	}
}
