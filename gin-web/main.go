package main

import (
	"gin-web/initialize/cacheRedis"
	"gin-web/initialize/config"
	mysqlDB "gin-web/initialize/mysql"
	"gin-web/initialize/runLog"
	"gin-web/models"
	"gin-web/models/authcCenter"
	"gin-web/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	//初始化配置文件
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	//设置运行模式
	if config.Conf.APP.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	//设置运行日志
	err = runLog.InitRunLog()
	if err != nil {
		panic(err)
	}
	//初始化mysql数据库
	err = mysqlDB.InitDB()
	if err != nil {
		panic(err)
	}
	//初始化缓存redis
	err = cacheRedis.InitRedis()
	if err != nil {
		panic(err)
	}
	//数据库迁移
	err = mysqlDB.DB.AutoMigrate(
		&authcCenter.User{},
		&authcCenter.Role{},
		&authcCenter.Api{},
		&models.OperationLog{},
	)
	if err != nil {
		panic(err)
	}
	go func() {
		log.Println(http.ListenAndServe("127.0.0.1:6066", nil))
	}()
}

func main() {
	//pprof检测程序性能
	routers.RouterServerRun() //http服务

}
