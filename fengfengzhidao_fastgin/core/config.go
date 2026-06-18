package core

import (
	"fengfengzhidao_fastgin/config"
	"fengfengzhidao_fastgin/flags"
	"fengfengzhidao_fastgin/global"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// 赋值给某个Config struct 的实例
func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile(flags.Options.File)
	if err != nil {
		logrus.Fatalf("配置文件读取错误 %s", err) // fatal: 这里出错了也没必要继续走下去
		return
	}
	// 可以用map接，或者别的结构体
	err = yaml.Unmarshal(byteData, cfg) //将配置文件里的值读取进cfg里，这样程序就可以使用配置文件了
	if err != nil {
		logrus.Fatalf("配置文件格式错误 %s", err)
		return
	}
	logrus.Infof("%s 配置文件读取成功", flags.Options.File)

	return
}

// 把内存里的值变成配置文件
func DumpConfig() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		logrus.Errorf("配置文件转换错误 %s", err) // 此时写出数据，程序还可能在运行，最好不要fatal
		return
	}

	// param1: 写出的文件地址
	// param2：写出的文件
	// param3：权限
	err = os.WriteFile(flags.Options.File, byteData, 0666)
	if err != nil {
		logrus.Errorf("配置文件格式错误 %s", err)
		return
	}
	logrus.Infof("配置文件写入成功")
}
