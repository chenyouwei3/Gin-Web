package main

import (
	"gin-web/initialize/cacheRedis"
	"gin-web/initialize/config"
	mysqlDB "gin-web/initialize/mysql"
	"gin-web/initialize/runLog"
	"gin-web/models/authcCenter"
	"gin-web/routers"
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
)

func init() {
	err := config.InitConfig() //初始化配置文件
	if err != nil {
		panic(err)
	}
	if config.Conf.APP.Mode == "debug" { //设置运行模式
		gin.SetMode(gin.DebugMode)
	}
	err = runLog.InitRunLog()
	if err != nil {
		panic(err)
	}
	err = mysqlDB.InitDB() //初始化mysql数据库
	if err != nil {
		panic(err)
	}
	err = cacheRedis.InitRedis() //初始化缓存redis
	if err != nil {
		panic(err)
	}
	err = mysqlDB.DB.AutoMigrate( //数据库迁移
		&authcCenter.User{},
		&authcCenter.Role{},
		&authcCenter.Api{},
	)
	if err != nil {
		panic(err)
	}
}

func main() {
	//pprof检测程序性能
	//go func() {
	//	log.Println(http.ListenAndServe("127.0.0.1:6066", nil))
	//}()
	routers.RouterServerRun()

}
