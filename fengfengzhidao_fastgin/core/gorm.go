package core

import (
	"fengfengzhidao_fastgin/global"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 获取gorm对象
func InitGorm() (db *gorm.DB) {
	// 通过不同的mode调用gorm

	cfg := global.Config.DB
	var dialector = cfg.DSN()
	if dialector == nil {
		return
	}
	/* 以下部分可以提取到config_db.go里，直接获取dialector
	switch cfg.Mode {
	// 在每个case里配置dsn
	case config.DBMysqlMode:
		// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		// 	cfg.User,
		// 	cfg.Password,
		// 	cfg.Host,
		// 	cfg.Port,
		// 	cfg.DBName,
		// )
		// dialector = mysql.Open(dsn) // 要下载mysql驱动
	case config.DBPgsqlMode:
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.Password,
			cfg.DBName,
		)
		dialector = postgres.Open(dsn) // 这是gorm.open的第一个参数
	case config.DBSqliteMode:
	default:
		logrus.Warnf("未配置sql连接")
		return
	}
	*/

	// 在外层使用gorm.open获得db对象
	// param2: 配置项
	db, err := gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生产实体外键
	})

	if err != nil {
		// 数据库都配置错了，肯定直接断开
		logrus.Fatalf("数据库连接失败 %s", err)
	}
	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalf("获取数据库连接失败 %s", err)
		return
	}
	err = sqlDB.Ping() // 执行pind才会真正尝试连接数据库，不然可能是延迟链接lazy connect
	if err != nil {
		logrus.Fatalf("数据库连接失败 %s", err)
		return
	}
	// 设置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	logrus.Infof("数据库连接成功！")

	return

}
