package mysql

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestNewMySQLClient(t *testing.T) {
	options := &MySQLOptions{
		Host:                  "127.0.0.1:3306",
		DbName:                "goadmin",
		Username:              "root",
		Password:              "root",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: 100000,
	}
	client := NewMySQLClient(options)
	fmt.Printf("client ====> %v \n", client)
	dbSQL, err := client.DB()
	if err != nil {
		log.Errorf(err.Error())
	}
	fmt.Printf("status ====> %v \n", dbSQL.Stats())
	fmt.Printf("ping ==== > %v \n", dbSQL.Ping())
	pool := client.ConnPool
	fmt.Printf("ping ==== > %v \n", pool)
}
