package main

import (
	"databases"
	"github.com/gin-gonic/gin"
	"github.com/polaris1119/config"
	"github.com/polaris1119/logger"
	"routers"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := routers.InitBackGroundRouter()
	defer databases.Orm.Close()
	//err := databases.Orm.Sync(new(model.Admin), new(model.AdminRole), new(model.AdminNavigation), new(model.AdminNavigationNode))
	//databases.Orm.Sync(new(model.Admin), new(model.AdminRole), new(model.AdminNavigation), new(model.AdminNavigationNode))
	//background.InitData()
	//err := databases.Orm.Sync2(new(model.AdminNavigation), new(model.AdminNavigationNode))
	//if err != nil {
	//	_, file, line, _ := runtime.Caller(1)
	//	fmt.Println(file, line) //***写入日志
	//	fmt.Println(err.Error())
	//}

	logger.Init(config.ROOT+"/log", config.ConfigFile.MustValue("global", "log_level", "DEBUG"))
	router.Run(":2009")

}
