package initialize

import (
	"loopy-manager/initialize/config"
	"loopy-manager/initialize/logs"
	"loopy-manager/initialize/mysql"
	"loopy-manager/initialize/redis"
	"loopy-manager/initialize/snowFlake"
	"loopy-manager/initialize/socketServer"
)

func InitConfig() {
	//数据库
	//database.MongodbInit(*system.Config.Mongodb)
	//database.MysqlInit(*system.Config.Mysql)
	redis.RedisInit(*config.Config.Redis)
	////系统
	snowFlake.SnowFlakeInit()
	logs.LogInit()
	//mysql主从结构体的初始化
	mysql.MysqlPoolInit()
	//Cache
	//go system.MysqlBinlogInit()
	//go messageQueue.RabbitmqInit()
	//接收设备
	socketServer.SocketServerStart()
}
