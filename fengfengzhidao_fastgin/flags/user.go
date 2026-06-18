// 命令行里用户相关的操作

package flags

import (
	"fengfengzhidao_fastgin/models"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

type User struct {
}

func (User) Create() {
	// 先选角色
	var user models.UserModel
	fmt.Println("请选择角色：1-管理员 2-普通用户")

	_, err := fmt.Scanln((&user.RoleID))
	if err != nil {
		fmt.Println("输入错误", err)
		return
	}

	if user.RoleID != 1 && user.RoleID != 2 {
		fmt.Println("用户角色输入错误")
		return
	}

	fmt.Println("请输入用户名")
	fmt.Scanln(&user.Username)

	fmt.Println("请输入密码") // terminal上不能看到密码，且要输入两次对比
	// fmt.Scanln(&user.Password)
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("读取密码时发生错误：", err)
		return
	}
	fmt.Println("请再次输入密码")
	rePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("读取密码时发生错误：", err)
		return
	}
	if string(password) != string(rePassword) {
		fmt.Println("两次密码不一致")
		return
	}

	fmt.Println(user.RoleID)
	fmt.Println(user.Username)
	fmt.Println(string(password))
}

func (User) List() {

}

func (User) Remove() {

}
