package flags

import (
	"fengfengzhidao_fastgin/global"
	"flag"
	"fmt"
)

// 创建一个结构体保存从命令行里传入的值
type FlagOptions struct {
	File    string // 配置文件
	Version bool   // 版本号
	DB      bool   // 表结构
}

// 用public参数，这样main和config都可以获取
var Options FlagOptions

func Parse() {
	// param1: 传入的参数变量映射到此处
	// param2：自定义的标识 -f
	// param3：如果不传入的默认值
	flag.StringVar(&Options.File, "f", "settings_dev.yaml", "config file path")
	flag.BoolVar(&Options.Version, "v", false, "打印当前版本号")
	flag.BoolVar(&Options.DB, "db", false, "迁移表结构")

	flag.Parse()
	// fmt.Println(Options.File)
}

func Run() (ok bool) {
	if Options.DB {
		fmt.Println("表结构迁移")
		return true
	}

	if Options.Version {
		fmt.Printf("版本号为%s", global.Version)
		return true
	}

	return false

}
