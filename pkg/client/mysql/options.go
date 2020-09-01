package mysql

import "time"

type MySQLOptions struct {
	Host                  string        `json:"host,omitempty" yaml:"host" description:"MySQL service host address"`
	Username              string        `json:"username,omitempty" yaml:"username"`
	Password              string        `json:"password" yaml:"password"`
	MaxIdleConnections    int           `json:"maxIdleConnections,omitempty" yaml:"maxIdleConnections"`
	MaxOpenConnections    int           `json:"maxOpenConnections,omitempty" yaml:"maxOpenConnections"`
	MaxConnectionLifeTime time.Duration `json:"maxConnectionLifeTime,omitempty" yaml:"maxConnectionLifeTime"`
}
