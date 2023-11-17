package initialize

import (
	"fmt"
	"loopy-manager/global"
)

func Init() {
	go SocketServer()
	RedisInit()
	MysqlInit()
	SnowFlakeInit()
	RabbitMqSendInit()

}

func DataBaseClose() {
	err := global.RedisClient.Close()
	if err != nil {
		fmt.Println("Error on closing redisService client.")
	}
	sql, err := global.MysqlClient.DB()
	if err != nil {
		fmt.Println("Error on closing dbService client.")
	}
	sql.Close()

}
