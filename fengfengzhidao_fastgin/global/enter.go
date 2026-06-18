package global

import (
	"fengfengzhidao_fastgin/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const Version = "1.0.0"

var (
	Config *config.Config
	DB     *gorm.DB //设置全局变量的数据库gorm对象
	Redis  *redis.Client
)
