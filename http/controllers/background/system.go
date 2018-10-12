package background

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//region Remark: 系统基础设置 Author:Qing
func GetSystemBase(c *gin.Context) {
	c.HTML(http.StatusOK, "/system/base", gin.H{
		"title": "BackGround Login",
		"info":  "",
	})
}

//endregion

//region Remark: 系统屏蔽词 Author:Qing
func GetSystemShield(c *gin.Context) {
	c.HTML(http.StatusOK, "/system/shield", gin.H{
		"title": "BackGround Center",
		"info":  "",
	})
}

//endregion

//region Remark: 系统日志 Author:Qing
func GetSystemLog(c *gin.Context) {
	c.HTML(http.StatusOK, "/system/log", gin.H{
		"title": "BackGround Center",
		"info":  "",
	})
}

//endregion
