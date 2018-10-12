package background

import (
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"logic"
	"model"
	"net/http"
	"strconv"
)

//region Remark: 导航节点列表 Author:Qing
func GetNavigationList(c *gin.Context) {
	data := logic.DefaultAdminNavigation.DataList(c)
	for _, value := range data {
		fmt.Println(value)
	}
	fmt.Println(data)
	c.HTML(http.StatusOK, "navigation/list", gin.H{
		"title":    "BackGround Center",
		"datalist": logic.DefaultAdminNavigation.DataList(c),
	})
}

//endregion

//region Remark: 添加导航节点 Author:Qing
func GetNavigationAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "navigation/add", gin.H{
		"title": "BackGround Center",
		"info":  "",
	})
}

func PostNavigationAdd(c *gin.Context) {
	title := c.PostForm("title")
	url := c.PostForm("url")
	parent_id, _ := strconv.ParseInt(c.PostForm("parent_id"), 10, 64)
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	var is_show = true
	var is_sys = false

	if c.PostForm("is_show") != "true" {
		is_show = false
	}

	if c.PostForm("is_sys") != "false" {
		is_sys = false
	}

	//节点
	//node_title := c.PostFormArray("node_title")
	//route_action := c.PostFormArray("route_action")
	//node_sort := c.PostFormArray("node_sort")
	//

	has, err := databases.Orm.Insert(&model.AdminNavigation{
		Title:    title,
		Url:      url,
		ParentId: parent_id,
		Sort:     sort,
		IsShow:   is_show,
		IsSys:    is_sys,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	if has < 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"msg":    "添加失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"msg":    "添加成功",
		})
		return
	}
}

//endregion
