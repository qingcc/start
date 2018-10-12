package background

import (
	"databases"
	"fmt"
	"model"
	"time"
)

//region Remark: 初始化数据 Author:Qing
func InitData() bool {
	has, err := databases.Orm.Insert(&model.AdminRole{
		RoleName: "超级管理员",
		IsSuper:  true,
		IsSys:    true,
	}, &model.AdminRole{
		RoleName: "普通管理员",
		IsSuper:  false,
		IsSys:    true,
	})
	if err != nil {
		fmt.Println("管理员角色初始化失败" + err.Error())
		return false
	}
	if has < 1 {
		fmt.Println("管理员角色初始化失败")
		return false
	}

	has, err = databases.Orm.Insert(&model.Admin{
		Name:          "admin",
		Email:         "admin@admin.com",
		Password:      "21232f297a57a5a743894a0e4a801fc3", //admin
		AdminRoleId:   1,
		LastLoginIp:   "127.0.0.1",
		LastLoginTime: time.Now().Format("2016-01-02 15:04:05"),
	}, &model.Admin{
		Name:          "qing",
		Email:         "qing@admin.com",
		Password:      "21232f297a57a5a743894a0e4a801fc3", //admin
		AdminRoleId:   2,
		LastLoginIp:   "127.0.0.1",
		LastLoginTime: time.Now().Format("2016-01-02 15:04:05"),
	})
	if err != nil {
		fmt.Println("管理员初始化失败" + err.Error())
		return false
	}
	if has < 1 {
		fmt.Println("管理员初始化失败")
		return false
	}

	return true
}

//endregion
