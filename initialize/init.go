package initialize

import (
	"LoopyTicker/global"
	"fmt"
)

func Init() {
	RedisInit()
	MysqlInit()
	SnowFlakeInit()
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
