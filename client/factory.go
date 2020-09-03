package client

import (
	"github.com/xuzhi7162/gcommon/client/mysql"
	"github.com/xuzhi7162/gcommon/client/redis"
	"github.com/xuzhi7162/gcommon/config"
)

// InitClient 根据配置文件配置初始化所需要客户端
func InitClient() {
	conf := config.ShareConfig
	newMySQLClient(conf)
	newRedisClient(conf)
}

// newMySQLClient 初始化 mysql 客户端
func newMySQLClient(conf *config.Config) {
	if conf.MySQLOptions != nil {
		mysql.NewMySQLClient(conf.MySQLOptions)
	}
}

// newRedisClient 初始化 redis 客户端
func newRedisClient(conf *config.Config) {
	if conf.RedisOptions != nil {
		redis.NewRedisClient(conf.RedisOptions)
	}
}
