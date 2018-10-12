package background

import (
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/polaris1119/logger"
	"logic"
	"model"
	"net/http"
	"strconv"
)

//region Remark: 管理员角色列表 Author:Qing
func GetAdminRoleList(c *gin.Context) {
	c.HTML(http.StatusOK, "/role/list", gin.H{
		"title":    "BackGround Center",
		"datalist": logic.DefaultRole.RoleList(c),
	})
}

//endregion

//region Remark: 添加新角色 Author:Qing
func GetAdminRoleAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "/role/add", gin.H{
		"title": "BackGround Center",
		"info":  "",
	})
}

func PostAdminRoleAdd(c *gin.Context) {
	role_name := c.PostForm("role_name")
	desccribe := c.PostForm("desccribe")
	if logic.DefaultRole.RoleIsExist(c, role_name) {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "该角色已存在",
		})
		return
	}
	has, err := databases.Orm.Insert(&model.AdminRole{RoleName: role_name, Describe: desccribe})
	if err != nil {
		fmt.Println(err.Error())
	}
	if has > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "新增角色成功",
			"url":    "/admin/admin_role/list",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"msg":    "新增角色失败",
	})
	return
}

//endregion

//region Remark: 编辑角色 Author:Qing
func GetAdminRoleEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	role, _ := logic.DefaultRole.GetOneRole(c, "id", id)

	c.HTML(http.StatusOK, "/role/edit", gin.H{
		"data": role,
	})
}

func PostAdminRoleEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	role_name := c.PostForm("role_name")
	desccribe := c.PostForm("desccribe")

	role := &model.AdminRole{}
	role.RoleName = role_name
	role.Describe = desccribe
	fmt.Println(role)
	has, err := databases.Orm.Id(id).Cols("role_name", "desc").Update(role)
	if err != nil {
		logger.Errorln("Role controller find errof:", err)
		fmt.Println(err.Error())
	}
	if has > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "更新成功",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpError,
		"msg":    "更新失败",
	})
	return
}

//endregion

//region Remark: 删除 角色 Author:Qing
func PostAdminRoleDel(c *gin.Context) {
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 64)
	ids := c.PostFormArray("id[]")
	if logic.DefaultRole.RoleHasAdmin(c, id, ids) {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "角色下有管理员,不能删除",
		})
		return
	}
	var has int64
	var err error
	if id == 0 {
		has, err = databases.Orm.In("id", ids).Delete(&model.AdminRole{})
	} else {
		has, err = databases.Orm.Id(id).Delete(&model.AdminRole{})
	}

	if err != nil {
		logger.Errorln("role controller del a role error")
	}

	if has < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":    "删除成功",
		"url":    "/admin/admin_role/list",
	})
	return
}

//endregion
