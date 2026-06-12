package flags

import (
	"flag"
)

// 创建一个结构体保存从命令行里传入的值
type FlagOptions struct {
	File string // 配置文件
}

// 用public参数，这样main和config都可以获取
var Options FlagOptions

func Run() {
	// param1: 传入的参数变量映射到此处
	// param2：自定义的标识 -f
	// param3：如果不传入的默认值
	flag.StringVar(&Options.File, "f", "settings_dev.yaml", "config file path")
	flag.Parse()
	// fmt.Println(Options.File)
}
