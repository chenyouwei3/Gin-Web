package initialize

import (
	"loopy-manager/initialize/config/database"
	"loopy-manager/initialize/config/messageQueue"
	"loopy-manager/initialize/config/system"
)

func InitConfig() {
	//数据库
	//database.MongodbInit(*system.Config.Mongodb)
	database.MysqlInit(*system.Config.Mysql)
	database.RedisInit(*system.Config.Redis)
	//消息队列
	messageQueue.RabbitmqInit()
	//系统
	system.SnowFlakeInit()
	system.LogInit()
	//mysql
	system.MysqlPoolInit()
	go system.MysqlBinlogInit()
}
