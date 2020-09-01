package config

import (
	"fmt"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	configInfo := &ConfigInfo{
		Name: "common",
		Path: "./",
		Type: "yaml",
	}
	conf, err := LoadConfig(configInfo)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("%v", conf.RedisOptions.Host)
}
