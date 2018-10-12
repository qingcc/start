package logic

import (
	"databases"
	"fmt"
	"model"
)

type AdminLogic struct{}

var DefaultAdmin = AdminLogic{}

//region Remark: 通过字段查询 管理员表 获取一个实例 Author:Qing
func (self AdminLogic) GetAdminByField(field string, val interface{}) *model.Admin {
	admin := new(model.Admin)
	_, err := databases.Orm.Where(field+" = ?", val).Get(admin)
	if err != nil {
		fmt.Println("查询管理员错误" + err.Error())
	}
	return admin
}

//endregion

//region Remark: 管理员列表 Author:Qing
func (self AdminLogic) GetAdminList(limit, page int, keyword string) ([]*model.AdminList, int, int, int) {
	admin := make([]*model.AdminList, 0)
	err := databases.Orm.Table("admin").Join("INNER", "admin_role", "admin.admin_role_id = admin_role.id")
	if keyword != "" {
		err = err.Where("name = ?", keyword)
	}

	err1 := *err
	num, _ := err1.Count()

	err2 := err.Limit(limit, limit*page).Find(&admin)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return admin, limit, page + 1, int(num)
}

//endregion
