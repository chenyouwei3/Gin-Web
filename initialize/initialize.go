package initialize

import (
	"loopy-manager/initialize/config/database"
	"loopy-manager/initialize/config/socketServer"
	"loopy-manager/initialize/config/system"
)

func InitConfig() {
	//数据库
	database.MongodbInit(*system.Config.Mongodb)
	database.MysqlInit(*system.Config.Mysql)
	database.RedisInit(*system.Config.Redis)
	////系统
	system.SnowFlakeInit()
	system.LogInit()
	//mysql主从结构体的初始化
	system.MysqlPoolInit()
	//Cache
	//go system.MysqlBinlogInit()
	//go messageQueue.RabbitmqInit()
	//接收设备
	socketServer.SocketServerStart()
}
