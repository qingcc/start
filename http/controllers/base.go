package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/polaris1119/logger"
	"logic"
)

//region Remark: 获取日志实例 Author:Qing
func getLogger(c *gin.Context) *logger.Logger {
	return logic.GetLogger(c)
}

//endregion
