package background

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"runtime"
)

//region Remark: 主页 Author:Qing
func GetCenter(c *gin.Context) {
	//logger.Errorln("连接mysql出错")
	//logger.Infof("this is %s", "2333") //this is 2333
	//logger.Infoln("主页2", "2333")				 //主页2 2333
	c.HTML(http.StatusOK, "/index", gin.H{
		"title": "BackGround Center",
		"info":  "",
	})
}

//endregion

//region Remark: 主页 Author:Qing
func GetWelcome(c *gin.Context) {
	var list map[string]interface{}
	list["goos"] = runtime.GOOS         //系统版本
	list["hostName"], _ = os.Hostname() //主机名
	list["goarch"] = runtime.GOARCH     //系统构架 如amd64
	list["numCPU"] = runtime.NumCPU()   //cpu数量

	c.HTML(http.StatusOK, "layouts/welcome", gin.H{
		"Title": "BackGround Center",
		"info":  "",
	})
}

//endregion
