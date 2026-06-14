package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBMode string

const (
	DBMysqlMode  = "mysql"
	DBPgsqlMode  = "pgsql"
	DBSqliteMode = "sqlite"
)

type DB struct {
	Mode     string `yaml:"mode"` //模式：mysql，pgsql，sqlite
	DBName   string `yaml:"db_name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password`
}

func (db DB) DSN() gorm.Dialector {
	switch db.Mode {
	// 在每个case里配置dsn
	case DBMysqlMode:
		// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		// 	cfg.User,
		// 	cfg.Password,
		// 	cfg.Host,
		// 	cfg.Port,
		// 	cfg.DBName,
		// )
		// dialector = mysql.Open(dsn) // 要下载mysql驱动
	case DBPgsqlMode:
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			db.Host,
			db.Port,
			db.User,
			db.Password,
			db.DBName,
		)
		return postgres.Open(dsn) // 这是gorm.open的第一个参数
	case DBSqliteMode:
	default:
		logrus.Warnf("未配置sql连接")
		return nil
	}
	return nil
}
