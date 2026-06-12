package main

import (
	"fengfengzhidao_fastgin/core"
	"fengfengzhidao_fastgin/flags"
	"fengfengzhidao_fastgin/global"
)

func main() {
	core.InitLogger()
	flags.Run()
	// cfg := core.ReadConfig() // 赋值给全局变量，这样就可以动态更新配置
	global.Config = core.ReadConfig()
	// global.Config.DB.User = "root" //代码中修改settings配置参数
	// fmt.Println(global.Config.DB)
	// core.DumpConfig() //将修改的配置参数写入settings

	// logrus.Debugf("hello")
	// logrus.Infof("hello")
	// logrus.Warnf("hello")
	// logrus.Errorf("hello")

}
