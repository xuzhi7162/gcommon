package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(options *RedisOptions) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", options.Host, options.Port),
		Username: options.Username,
		Password: options.Password,
		DB:       0,
	})
	return client
}
