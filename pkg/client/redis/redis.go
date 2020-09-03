package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

var _client *redis.Client

func NewRedisClient(options *RedisOptions) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", options.Host, options.Port),
		Username: options.Username,
		Password: options.Password,
		DB:       0,
	})
	_client = client
}

func getClient() *redis.Client {
	return _client
}
