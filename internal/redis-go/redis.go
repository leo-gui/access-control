package redis_go

import (
	"access-control/internal"
	"github.com/go-redis/redis"
)



func InitRedis() (Redis *redis.Client) {
	var err error
	redisConfig := internal.InitRedisByViper()
	Redis = redis.NewClient(&redis.Options{
		Addr:    redisConfig.Addr,
		Password: redisConfig.PassWord,
		DB:       redisConfig.DB,
	})
	_, err = Redis.Ping().Result()
	if err != nil {
		return nil
	}
	return Redis
}
