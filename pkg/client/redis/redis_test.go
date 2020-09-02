package redis

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	options := &RedisOptions{
		Host:     "127.0.0.1",
		Port:     "6379",
		Username: "",
		Password: "",
	}
	client := NewRedisClient(options)
	fmt.Printf("%v", client)
	err := client.Set(context.Background(), "gcommon::test::username", "admin", 0).Err()
	if err != nil {
		log.Errorf(err.Error())
	}
	client.Set(context.Background(), "gcommon::test::password", "root", 0)
}
