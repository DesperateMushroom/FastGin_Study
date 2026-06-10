package core

import (
	"fengfengzhidao_fastgin/config"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// 赋值给某个Config struct 的实例
func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile("settings_dev.yaml")
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
