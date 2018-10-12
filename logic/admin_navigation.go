package logic

import (
	"databases"
	"github.com/gin-gonic/gin"
	"model"
)

type AdminNavigationLogic struct{}

var DefaultAdminNavigation = AdminNavigationLogic{}

//region Remark: 权限节点列表 Author:Qing
func (self AdminNavigationLogic) DataList(c *gin.Context) []*model.AdminNavigation {
	objLog := GetLogger(c)
	navigation := make([]*model.AdminNavigation, 0)
	err := databases.Orm.Find(&navigation)
	if err != nil {
		objLog.Errorf("RoleLogic find errof:", err)
		return nil
	}
	return navigation
}

//endregion
