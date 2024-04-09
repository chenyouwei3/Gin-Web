package initialize

import (
	"loopy-manager/initialize/config/database"
	"loopy-manager/initialize/config/system"
)

func InitConfig() {
	//数据库
	//database.MongodbInit(*system.Config.Mongodb)
	database.MysqlInit(*system.Config.Mysql)
	database.RedisInit(*system.Config.Redis)
	////系统
	system.SnowFlakeInit()
	system.LogInit()
	////mysql
	system.MysqlPoolInit() //mysql主从结构体的初始化
	//go messageQueue.RabbitmqInit()
	////mysql
	//go system.MysqlBinlogInit()
}
