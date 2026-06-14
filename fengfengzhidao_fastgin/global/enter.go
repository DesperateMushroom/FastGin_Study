package global

import (
	"fengfengzhidao_fastgin/config"

	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB //设置全局变量的数据库gorm对象
)
