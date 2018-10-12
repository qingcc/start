package routers

import (
	"databases"
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"http/controllers/background"
	"http/middleware"
)

func InitBackGroundRouter() *gin.Engine {
	router := gin.Default()
	router.Use(sessions.Sessions("go", databases.SessionStore()))
	databases.Init()
	//region Remark: 定义公共的中间件 Author; chijian

	//endregion
	router.Static("/public", "./static") //渲染html页面//router.LoadHTMLGlob("resources/**/*")
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "resources/background",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	})
	router.Use(middleware.Web("background"))
	router.GET("/login", background.GetLogin)
	router.POST("/login", background.PostLogin)

	v1 := router.Group("/admin", middleware.Auth())
	v1.GET("/center", background.GetCenter)
	v1.GET("/welcome", background.GetWelcome)

	//管理员
	v1.GET("/admin_user/list", background.GetAdminUserList)
	v1.GET("/admin_user/add", background.GetAdminUserAdd)
	v1.POST("/admin_user/add", background.PostAdminUserAdd)
	v1.GET("/admin_user/edit", background.GetAdminUserEdit)
	v1.POST("/admin_user/edit", background.PostAdminUserEdit)
	v1.POST("/admin_user/del", background.PostAdminUserDel)
	v1.POST("/admin_user/enable", background.PostAdminUserEnable)

	//角色
	v1.GET("/admin_role/list", background.GetAdminRoleList)
	v1.GET("/admin_role/add", background.GetAdminRoleAdd)
	v1.POST("/admin_role/add", background.PostAdminRoleAdd)
	v1.GET("/admin_role/edit", background.GetAdminRoleEdit)
	v1.POST("/admin_role/edit", background.PostAdminRoleEdit)
	v1.POST("/admin_role/del", background.PostAdminRoleDel)

	//权限节点
	v1.GET("/navigation/list", background.GetNavigationList)
	v1.GET("/navigation/add", background.GetNavigationAdd)
	v1.POST("/navigation/add", background.PostNavigationAdd)

	//系统
	v1.GET("/system/base", background.GetSystemBase)
	v1.GET("/system/shield", background.GetSystemShield)
	v1.GET("/system/log", background.GetSystemLog)

	return router
}
