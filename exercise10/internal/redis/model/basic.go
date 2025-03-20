package modelredis

import (
	"github.com/my/repo/internal/config"
	"github.com/redis/go-redis/v9"
)

type LeaderBoardRedisModel struct {
	conf *config.Config
	rdb  *redis.Client
}

func NewLeaderBoardRedisModel(conf *config.Config, rdb *redis.Client) *LeaderBoardRedisModel {
	return &LeaderBoardRedisModel{
		conf: conf,
		rdb:  rdb,
	}
}
