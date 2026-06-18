// 命令行里数据库表单相关的操作

package flags

import (
	"fengfengzhidao_fastgin/global"
	"fengfengzhidao_fastgin/models"

	"github.com/sirupsen/logrus"
)

// 迁移表
func MigrateDB() {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		logrus.Errorf("表结构迁移失败 %s", err)
		return
	}

	logrus.Infof("表结构迁移成功")
}
