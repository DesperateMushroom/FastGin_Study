package flags

import (
	"fengfengzhidao_fastgin/global"
	"flag"
	"fmt"
	"os"
)

// 创建一个结构体保存从命令行里传入的值
type FlagOptions struct {
	File    string // 配置文件
	Version bool   // 版本号
	DB      bool   // 表结构
	Menu    string
	Type    string // 类型:create/list/remove
}

// 用public参数，这样main和config都可以获取
var Options FlagOptions

func Parse() {
	// param1: 传入的参数变量映射到此处
	// param2：自定义的标识 -f
	// param3：如果不传入的默认值
	flag.StringVar(&Options.File, "f", "settings_dev.yaml", "config file path")
	flag.StringVar(&Options.Menu, "m", "", "菜单：user")
	flag.StringVar(&Options.Type, "t", "", "类型：create; list; remove")

	flag.BoolVar(&Options.Version, "v", false, "打印当前版本号")
	flag.BoolVar(&Options.DB, "db", false, "迁移表结构")

	flag.Parse()
	// fmt.Println(Options.File)
}

func Run() {
	if Options.DB {
		MigrateDB()
		// return true 没必要返回了，创建表后可以直接exit了
		os.Exit(0)
	}

	if Options.Version {
		fmt.Printf("版本号为%s", global.Version)
		os.Exit(0)

	}

	if Options.Menu == "user" {
		var user User
		switch Options.Type {
		case "create":
			user.Create()

		case "list":
			user.List()
		case "remove":
			user.Remove()
		default:
			os.Exit(0)
		}
		os.Exit(0) // 只要命令行里参数有user，不必开启别的web服务，直接退掉
	}

}
