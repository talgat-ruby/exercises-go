package newredis

import (
	"github.com/my/repo/internal/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	newClient *redis.Client
}

func NewRedis(conf *config.Config) *Redis {
	rdb := RedisClient(conf)

	redis.NewClient(&redis.Options{
		Addr: conf.Redis.Host,
	})
	return &Redis{
		newClient: rdb,
	}
}

func RedisClient(conf *config.Config) *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr: conf.Redis.Host,
	})
}
