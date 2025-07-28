package main

import (
	conf "gin-web/init/config"
	mysqlDB "gin-web/init/mysql"
	"gin-web/init/runLog"
	"gin-web/internal/manage/middleware"
	"gin-web/internal/manage/routers"
	"log"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func init() {
	//配置文件加载
	err := conf.InitConfig("../configs", "config-app1", "yml")
	if err != nil {
		log.Println("配置文件加载失败", err)
		panic(err)
	}
	//设置gin框架运行模式
	if conf.SystemConfig.APP.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	//设置运行日志
	if err = runLog.InitRunLog(conf.SystemConfig); err != nil {
		log.Println("运行日志初始化失败", err)
		panic(err)
	}
	//设置mysql数据库
	if err = mysqlDB.InitDB(conf.SystemConfig.Mysql); err != nil {
		log.Println("MySql初始化失败", err)
		panic(err)
	}

	log.Println("系统初始化加载完毕")
}

func main() {
	//pprof设置
	//go func() {
	//	http.ListenAndServe("0.0.0.0:6060", nil)
	//}()
	//运行消费操作日志协程
	middleware.InitOperationLogWorker(runLog.ZapLog, mysqlDB.DB)
	//运行日志退出
	defer func() {
		if err := runLog.ZapLog.Sync(); err != nil {
			log.Println("运行日志刷出失败:", err)
		}
	}()
	if err := routers.NewRouter().Run(":" + conf.SystemConfig.APP.Port); err != nil {
		panic(err)
	}
}
