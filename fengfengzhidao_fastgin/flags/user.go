// 命令行里用户相关的操作

package flags

import (
	"fengfengzhidao_fastgin/global"
	"fengfengzhidao_fastgin/models"
	"fengfengzhidao_fastgin/utils/pwd"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
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

	if user.Username == "" {
		fmt.Println("用户名不可为空！")
		return
	}

	var u models.UserModel // 查重名
	err = global.DB.Take(&u, "username=?", user.Username).Error
	if err == nil {
		// 找着了，数据库里有个重名的
		fmt.Println("用户名已存在")
		return
	}

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

	hashPwd := pwd.GenerateFromPassword(string(password))
	err = global.DB.Create(&models.UserModel{
		Username: user.Username,
		Password: hashPwd,
		RoleID:   user.RoleID,
	}).Error
	if err != nil {
		logrus.Errorf("用户创建失败 %s", err)
		return
	}
}

func (User) List() {
	var userList []models.UserModel

	// 查看创建最新的前十个
	global.DB.Order("created_at desc").Limit(10).Find(&userList)
	for _, model := range userList {
		fmt.Printf("用户id: %d - 用户名: %s - 用户昵称：%s - 用户角色： %d - 创建时间: %s\n",
			model.ID,
			model.Username,
			model.Nickname,
			model.RoleID,
			model.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}
}

func (User) Remove() {

}
