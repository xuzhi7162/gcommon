package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xuzhi7162/gcommon/client/mysql"
	"github.com/xuzhi7162/gcommon/client/redis"
)

var ShareConfig *Config

// LoadConfig 加载配置文件到viper
func LoadConfig(configInfo *ConfigInfo) (*Config, error) {
	// 设置配置文件名、路径、文件路径
	viper.SetConfigName(configInfo.Name)
	viper.AddConfigPath(configInfo.Path)
	viper.SetConfigType(configInfo.Type)

	// 从配置文件中加载配置文件内容到viper上下文环境中
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Errorf("%s/%s.%s配置文件不存在", configInfo.Path, configInfo.Name, configInfo.Type)
			return nil, nil
		} else {
			panic(fmt.Errorf("配置文件解析失败 %v", err))
		}
	}

	// 初始化配置内容
	conf := newConfig()
	// 读取viper中的配置,并加载到conf中
	if err := viper.Unmarshal(conf); err != nil {
		log.Errorf("配置文件加载转换失败, %v", err)
	} else {
		ShareConfig = conf
	}
	return ShareConfig, nil
}

// newConfig 初始化配置文件结构体
func newConfig() *Config {
	return &Config{
		MySQLOptions: nil,
		RedisOptions: nil,
	}
}

// Config 配置文件结构体
type Config struct {
	MySQLOptions *mysql.MySQLOptions `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	RedisOptions *redis.RedisOptions `json:"redis" yaml:"redis" mapstructure:"redis"`
}

// ConfigInfo 配置文件名、路径、文件类型
type ConfigInfo struct {
	Name string
	Path string
	Type string
}
