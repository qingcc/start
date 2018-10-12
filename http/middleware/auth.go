package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vendors/session"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//region Remark: 判断session里是否存在 adminid Author:Qing
		if session.HasSession(c, "adminid") == false {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		}
		//endregion

		//region Remark: 判断权限 Author:Qing

		//endregion

		c.Next()
	}
}
