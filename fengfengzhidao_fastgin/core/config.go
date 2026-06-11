package core

import (
	"fengfengzhidao_fastgin/config"
	"fengfengzhidao_fastgin/flags"
	"fengfengzhidao_fastgin/global"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// 赋值给某个Config struct 的实例
func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile(flags.Options.File)
	if err != nil {
		fmt.Printf("配置文件读取错误 %s", err)
		return
	}
	// 可以用map接，或者别的结构体
	err = yaml.Unmarshal(byteData, cfg)
	if err != nil {
		fmt.Printf("配置文件格式错误 %s", err)
		return
	}

	return
}

// 把内存里的值变成配置文件
func DumpConfig() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		fmt.Printf("配置文件转换错误 %s", err)
		return
	}

	// param1: 写出的文件地址
	// param2：写出的文件
	// param3：权限
	err = os.WriteFile(flags.Options.File, byteData, 0666)
	if err != nil {
		fmt.Printf("配置文件格式错误 %s", err)
		return
	}
	fmt.Println("配置文件写入成功")
}
