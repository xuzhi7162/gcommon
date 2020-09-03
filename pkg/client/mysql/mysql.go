package mysql

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func NewMySQLClient(options *MySQLOptions) {
	mysqlURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=1&multiStatements=1&charset=utf8&collation=utf8_general_ci",
		options.Username,
		options.Password,
		options.Host,
		options.DbName)
	db, err := gorm.Open(mysql.Open(mysqlURL), &gorm.Config{})
	if err != nil {
		panic("数据库客户端创建失败")
	}
	// 配置数据库连接池
	configSQLPool(db, options)
	_db = db
}

func configSQLPool(db *gorm.DB, options *MySQLOptions) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("提取DB失败")
	}
	// 配置sql连接池
	sqlDB.SetMaxIdleConns(options.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(options.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(options.MaxConnectionLifeTime)
}

func getDB() *gorm.DB {
	return _db
}
